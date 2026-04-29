package dealpusher

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	cockroach "github.com/cockroachdb/errors"
	synapse "github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/constants"
	"github.com/data-preservation-programs/go-synapse/pdp"
	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/go-synapse/spregistry"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

// OnChainPDPConfig configures the FWSS-pull adapter.
type OnChainPDPConfig struct {
	// DB is required.
	DB *gorm.DB
	// RPCURL is the FEVM JSON-RPC endpoint (read-only; we never submit txs).
	RPCURL string
	// SourceURLBase is the HTTPS base singularity serves pieces from. The
	// per-piece URL is constructed as <base>/piece/<pieceCidV2>. Required.
	SourceURLBase string
	// RecordKeeper is the FWSS contract address (hex). Empty defaults to
	// the network FWSS from go-synapse constants.
	RecordKeeper string
}

// OnChainPDP drives the FWSS-mediated pull flow. We never submit any
// PDPVerifier tx ourselves: the SP downloads pieces from our content
// provider via /pdp/piece/pull, then commits on-chain from its own wallet
// via /pdp/data-sets/create-and-add (new sets) or /pdp/data-sets/{id}/pieces
// (existing). The only EVM RPC use is the ServiceProviderRegistry view
// call that resolves an SP's PDP service URL.
type OnChainPDP struct {
	db            *gorm.DB
	ethClient     *ethclient.Client
	network       constants.Network
	chainID       *big.Int
	recordKeeper  common.Address
	sourceURLBase string
	spRegistry    *spregistry.Service

	spURLCacheMu sync.Mutex
	spURLCache   map[string]spInfo // delegated f4 provider string -> info
}

type spInfo struct {
	serviceURL string
	payee      common.Address
}

func NewOnChainPDP(ctx context.Context, cfg OnChainPDPConfig) (*OnChainPDP, error) {
	if cfg.DB == nil {
		return nil, errors.New("db is required")
	}
	if cfg.RPCURL == "" {
		return nil, errors.New("eth rpc URL is required")
	}
	if cfg.SourceURLBase == "" {
		return nil, errors.New("source URL base is required (--pdp-source-url-base / PDP_SOURCE_URL_BASE)")
	}

	ethClient, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, cockroach.Wrap(err, "failed to connect to FEVM RPC")
	}

	network, chainIDInt64, err := synapse.DetectNetwork(ctx, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, cockroach.Wrap(err, "failed to detect FEVM network")
	}
	chainID := big.NewInt(chainIDInt64)

	recordKeeper := common.HexToAddress(cfg.RecordKeeper)
	if recordKeeper == (common.Address{}) {
		recordKeeper = constants.WarmStorageAddresses[network]
		if recordKeeper == (common.Address{}) {
			ethClient.Close()
			return nil, fmt.Errorf("no FWSS recordKeeper address for network %s; set --pdp-record-keeper", network)
		}
	}

	spRegAddr := constants.SPRegistryAddresses[network]
	if spRegAddr == (common.Address{}) {
		ethClient.Close()
		return nil, fmt.Errorf("no ServiceProviderRegistry address for network %s", network)
	}
	spReg, err := spregistry.NewService(ethClient, spRegAddr, nil, chainID)
	if err != nil {
		ethClient.Close()
		return nil, cockroach.Wrap(err, "failed to bind ServiceProviderRegistry")
	}

	Logger.Infow("initialized FWSS-pull adapter",
		"network", network,
		"chainId", chainIDInt64,
		"recordKeeper", recordKeeper.Hex(),
		"spRegistry", spRegAddr.Hex(),
		"sourceURLBase", cfg.SourceURLBase,
	)

	return &OnChainPDP{
		db:            cfg.DB,
		ethClient:     ethClient,
		network:       network,
		chainID:       chainID,
		recordKeeper:  recordKeeper,
		sourceURLBase: strings.TrimSuffix(cfg.SourceURLBase, "/"),
		spRegistry:    spReg,
		spURLCache:    map[string]spInfo{},
	}, nil
}

func (o *OnChainPDP) Close() error {
	if o.ethClient != nil {
		o.ethClient.Close()
	}
	return nil
}

// PullPiecesToFWSS implements PDPProofSetManager.
func (o *OnChainPDP) PullPiecesToFWSS(
	ctx context.Context,
	evmSigner signer.EVMSigner,
	provider string,
	pieces []PDPPieceInput,
	cfg PDPSchedulingConfig,
) (PDPPullResult, error) {
	if evmSigner == nil {
		return PDPPullResult{}, errors.New("evm signer is required")
	}
	if len(pieces) == 0 {
		return PDPPullResult{}, errors.New("no pieces provided")
	}
	if cfg.BatchSize > 0 && len(pieces) > cfg.BatchSize {
		return PDPPullResult{}, fmt.Errorf("piece count %d exceeds configured batch size %d", len(pieces), cfg.BatchSize)
	}

	payerAddr := evmSigner.EVMAddress()
	clientDelegated, err := commonToDelegatedAddress(payerAddr)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "derive delegated payer address")
	}
	clientAddrStr := clientDelegated.String()

	info, err := o.lookupSPInfo(ctx, provider)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "SP service-URL lookup")
	}

	// Convert v1 piece CIDs to CommPv2 (what FWSS / Curio expect).
	pieceCIDsV2 := make([]cid.Cid, len(pieces))
	pullInputs := make([]pdp.PullPieceInput, len(pieces))
	for i, p := range pieces {
		// FR32 padding: padded = raw * 128/127, raw = padded * 127/128.
		payloadSize := uint64(p.PieceSize) * 127 / 128
		v2, err := commcid.PieceCidV2FromV1(p.PieceCID, payloadSize)
		if err != nil {
			return PDPPullResult{}, cockroach.Wrapf(err, "convert piece %s to CommPv2", p.PieceCID)
		}
		pieceCIDsV2[i] = v2
		pullInputs[i] = pdp.PullPieceInput{
			PieceCID:  v2.String(),
			SourceURL: o.sourceURLBase + "/piece/" + v2.String(),
		}
	}

	existing, err := o.findProofSetWithRoom(ctx, clientAddrStr, provider, cfg.MaxPiecesPerProofSet)
	if err != nil {
		return PDPPullResult{}, err
	}

	authHelper := pdp.NewAuthHelper(evmSigner.SignDigest, payerAddr, o.recordKeeper, o.chainID)
	pdpServer := pdp.NewServer(info.serviceURL)

	if existing != nil {
		return o.addToExisting(ctx, pdpServer, authHelper, existing, pullInputs, pieceCIDsV2, cfg)
	}

	return o.createNewSet(ctx, pdpServer, authHelper, payerAddr, info.payee, provider, clientAddrStr, info.serviceURL, pullInputs, pieceCIDsV2, cfg)
}

// addToExisting adds pieces to an assembling proof set we've already
// created. extraData is the AddPieces blob alone.
func (o *OnChainPDP) addToExisting(
	ctx context.Context,
	pdpServer *pdp.Server,
	authHelper *pdp.AuthHelper,
	existing *model.PDPProofSet,
	pullInputs []pdp.PullPieceInput,
	pieceCIDsV2 []cid.Cid,
	cfg PDPSchedulingConfig,
) (PDPPullResult, error) {
	clientDataSetID, ok := new(big.Int).SetString(existing.ClientDataSetID, 10)
	if !ok || clientDataSetID == nil {
		return PDPPullResult{}, fmt.Errorf("proof set %d has invalid clientDataSetID %q", existing.SetID, existing.ClientDataSetID)
	}

	addExtra, err := signAddPiecesExtra(authHelper, clientDataSetID, pieceCIDsV2)
	if err != nil {
		return PDPPullResult{}, err
	}

	if err := waitForPullComplete(ctx, pdpServer, pdp.PullPiecesOptions{
		RecordKeeper: o.recordKeeper.Hex(),
		Pieces:       pullInputs,
		ExtraData:    addExtra,
		DataSetID:    existing.SetID,
	}, cfg.PullTimeout); err != nil {
		return PDPPullResult{}, err
	}

	addResp, err := pdpServer.AddPieces(ctx, int(existing.SetID), pieceCIDsV2, addExtra)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "POST /pdp/data-sets/{id}/pieces")
	}
	status, err := pdpServer.WaitForPieceAddition(ctx, int(existing.SetID), addResp.TxHash, cfg.PullTimeout)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "wait for add-pieces tx")
	}
	if status.AddMessageOK == nil || !*status.AddMessageOK {
		return PDPPullResult{}, fmt.Errorf("add-pieces tx %s did not confirm successfully", addResp.TxHash)
	}

	if err := o.db.WithContext(ctx).
		Model(&model.PDPProofSet{}).
		Where("set_id = ?", existing.SetID).
		UpdateColumn("piece_count", gorm.Expr("piece_count + ?", len(pieceCIDsV2))).Error; err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "increment piece_count")
	}

	return PDPPullResult{DataSetID: existing.SetID}, nil
}

// createNewSet atomically creates a FWSS-listened proof set with the first
// batch of pieces. extraData is the abi.encode(bytes,bytes) of the
// CreateDataSet and AddPieces extras.
func (o *OnChainPDP) createNewSet(
	ctx context.Context,
	pdpServer *pdp.Server,
	authHelper *pdp.AuthHelper,
	payerAddr common.Address,
	payee common.Address,
	provider string,
	clientAddrStr string,
	serviceURL string,
	pullInputs []pdp.PullPieceInput,
	pieceCIDsV2 []cid.Cid,
	cfg PDPSchedulingConfig,
) (PDPPullResult, error) {
	clientDataSetID := randomClientDataSetID()

	createExtra, err := signCreateDataSetExtra(authHelper, payerAddr, payee, clientDataSetID)
	if err != nil {
		return PDPPullResult{}, err
	}
	addExtra, err := signAddPiecesExtra(authHelper, clientDataSetID, pieceCIDsV2)
	if err != nil {
		return PDPPullResult{}, err
	}
	combined, err := pdp.EncodeCreateDataSetAndAddPiecesExtraData(createExtra, addExtra)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "wrap create+add extra data")
	}

	if err := waitForPullComplete(ctx, pdpServer, pdp.PullPiecesOptions{
		RecordKeeper: o.recordKeeper.Hex(),
		Pieces:       pullInputs,
		ExtraData:    combined,
		DataSetID:    0,
	}, cfg.PullTimeout); err != nil {
		return PDPPullResult{}, err
	}

	createResp, err := pdpServer.CreateDataSetAndAddPieces(ctx, o.recordKeeper.Hex(), pieceCIDsV2, combined)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "POST /pdp/data-sets/create-and-add")
	}

	status, err := pdpServer.WaitForDataSetCreation(ctx, createResp.TxHash, cfg.PullTimeout)
	if err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "wait for create-and-add tx")
	}
	if status.DataSetID == nil {
		return PDPPullResult{}, fmt.Errorf("create-and-add tx %s confirmed but no dataSetId returned", createResp.TxHash)
	}
	dataSetID := uint64(*status.DataSetID)

	row := model.PDPProofSet{
		SetID:           dataSetID,
		ClientAddress:   clientAddrStr,
		Provider:        provider,
		HandoffState:    model.ProofSetAssembling,
		ClientDataSetID: clientDataSetID.String(),
		ServiceURL:      serviceURL,
		PieceCount:      len(pieceCIDsV2),
	}
	// Upsert so we win the race against pdptracker materializing the same
	// SetID from the DataSetCreated event with listener-as-client_address.
	if err := o.db.WithContext(ctx).
		Where("set_id = ?", dataSetID).
		Assign(row).
		FirstOrCreate(&model.PDPProofSet{}).Error; err != nil {
		return PDPPullResult{}, cockroach.Wrap(err, "persist new proof set")
	}

	return PDPPullResult{DataSetID: dataSetID}, nil
}

// waitForPullComplete blocks until the SP-side piece transfer reports
// complete, fails permanently, or times out.
func waitForPullComplete(ctx context.Context, pdpServer *pdp.Server, opts pdp.PullPiecesOptions, timeout time.Duration) error {
	resp, err := pdpServer.WaitForPullPieces(ctx, opts, timeout)
	if err != nil {
		return cockroach.Wrap(err, "wait for /pdp/piece/pull")
	}
	switch resp.Status {
	case pdp.PullStatusComplete:
		return nil
	case pdp.PullStatusFailed:
		return fmt.Errorf("/pdp/piece/pull reported failed for %d pieces", len(opts.Pieces))
	default:
		return fmt.Errorf("/pdp/piece/pull timed out at status %q", resp.Status)
	}
}

func signCreateDataSetExtra(authHelper *pdp.AuthHelper, payer, payee common.Address, clientDataSetID *big.Int) (string, error) {
	sig, err := authHelper.SignCreateDataSet(clientDataSetID, payee, nil)
	if err != nil {
		return "", cockroach.Wrap(err, "sign CreateDataSet")
	}
	return pdp.EncodeDataSetCreateData(payer, clientDataSetID, nil, sig.Signature)
}

func signAddPiecesExtra(authHelper *pdp.AuthHelper, clientDataSetID *big.Int, pieceCIDsV2 []cid.Cid) (string, error) {
	// FWSS uses (payer, clientDataSetId) as the cross-tx replay key, not
	// the addPieces in-extraData nonce; zero is fine here.
	nonce := big.NewInt(0)
	sig, err := authHelper.SignAddPieces(clientDataSetID, nonce, pieceCIDsV2, nil)
	if err != nil {
		return "", cockroach.Wrap(err, "sign AddPieces")
	}
	return pdp.EncodeAddPiecesExtraData(nonce, nil, sig.Signature)
}

func (o *OnChainPDP) findProofSetWithRoom(ctx context.Context, clientAddress, provider string, maxPieces int) (*model.PDPProofSet, error) {
	if maxPieces <= 0 {
		return nil, errors.New("max pieces per proof set must be > 0")
	}
	var ps model.PDPProofSet
	err := o.db.WithContext(ctx).
		Where("client_address = ? AND provider = ? AND deleted = FALSE AND handoff_state = ?",
			clientAddress, provider, model.ProofSetAssembling).
		Where("piece_count < ?", maxPieces).
		Order("set_id").
		First(&ps).Error
	if err == nil {
		return &ps, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, cockroach.Wrap(err, "query assembling proof sets")
}

func (o *OnChainPDP) lookupSPInfo(ctx context.Context, provider string) (spInfo, error) {
	o.spURLCacheMu.Lock()
	if cached, ok := o.spURLCache[provider]; ok {
		o.spURLCacheMu.Unlock()
		return cached, nil
	}
	o.spURLCacheMu.Unlock()

	evmAddr, err := delegatedToCommonAddress(provider)
	if err != nil {
		return spInfo{}, cockroach.Wrapf(err, "decode provider address %s", provider)
	}

	pi, err := o.spRegistry.GetProviderByAddress(ctx, evmAddr)
	if err != nil {
		return spInfo{}, cockroach.Wrapf(err, "ServiceProviderRegistry.getProviderByAddress(%s)", evmAddr.Hex())
	}
	if pi == nil {
		return spInfo{}, fmt.Errorf("provider %s not registered in ServiceProviderRegistry", provider)
	}
	pdpProduct, ok := pi.Products["PDP"]
	if !ok || pdpProduct == nil || pdpProduct.Data == nil {
		return spInfo{}, fmt.Errorf("provider %s has no PDP product registered", provider)
	}
	if !pdpProduct.IsActive {
		return spInfo{}, fmt.Errorf("provider %s PDP product is inactive", provider)
	}
	if pdpProduct.Data.ServiceURL == "" {
		return spInfo{}, fmt.Errorf("provider %s has no PDP serviceURL registered", provider)
	}

	info := spInfo{
		serviceURL: pdpProduct.Data.ServiceURL,
		payee:      pi.Payee,
	}
	o.spURLCacheMu.Lock()
	o.spURLCache[provider] = info
	o.spURLCacheMu.Unlock()
	return info, nil
}

func delegatedToCommonAddress(s string) (common.Address, error) {
	addr, err := address.NewFromString(s)
	if err != nil {
		return common.Address{}, err
	}
	if addr.Protocol() != address.Delegated {
		return common.Address{}, fmt.Errorf("address %s is not delegated (f410)", s)
	}
	payload := addr.Payload()
	if len(payload) < 21 {
		return common.Address{}, fmt.Errorf("delegated payload too short: %d", len(payload))
	}
	return common.BytesToAddress(payload[1:21]), nil
}

func commonToDelegatedAddress(subaddr common.Address) (address.Address, error) {
	addr, err := address.NewDelegatedAddress(10, subaddr.Bytes())
	if err != nil {
		return address.Undef, cockroach.Wrap(err, "encode delegated address")
	}
	return addr, nil
}

func randomClientDataSetID() *big.Int {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("crypto/rand failure: %v", err))
	}
	return new(big.Int).SetBytes(b)
}
