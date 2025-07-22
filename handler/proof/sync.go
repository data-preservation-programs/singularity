package proof

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type SyncProofRequest struct {
	DealID   *uint64 `json:"dealId"`   // specific deal ID to sync proofs for
	Provider *string `json:"provider"` // specific provider to sync proofs for
}

// Message represents a Filecoin message from the chain
type Message struct {
	To     string `json:"To"`
	From   string `json:"From"`
	Method uint64 `json:"Method"`
	Params string `json:"Params"`
}

// MessageReceipt represents a message receipt
type MessageReceipt struct {
	ExitCode int    `json:"ExitCode"`
	Return   string `json:"Return"`
	GasUsed  int64  `json:"GasUsed"`
}

// InvocResult represents the result of StateReplay
type InvocResult struct {
	MsgCid struct {
		CID string `json:"/"`
	} `json:"MsgCid"`
	Msg      Message        `json:"Msg"`
	MsgRct   MessageReceipt `json:"MsgRct"`
	Error    string         `json:"Error"`
	Duration int64          `json:"Duration"`
}

// TipSet represents a tipset with block CIDs
type TipSet []struct {
	CID string `json:"/"`
}

// ChainHead represents the chain head response
type ChainHead struct {
	Height int64  `json:"Height"`
	Cids   TipSet `json:"Cids"`
}

// Miner actor method numbers for proof-related operations
const (
	MethodSubmitWindowedPoSt = 5 // WindowPoSt proofs (spacetime)
	MethodProveCommitSector  = 7 // Sector commit proofs (replication)
	MethodPreCommitSector    = 6 // PreCommit sector
)

// SyncHandler synchronizes proofs from the Filecoin chain into the database.
//
// This handler can sync proofs for all deals, a specific deal, or a specific provider
// based on the request parameters. It fetches messages from the Filecoin chain,
// identifies proof-related messages, and stores them in the database.
//
// Parameters:
//   - ctx:         The context for the operation which provides facilities for timeouts and cancellations.
//   - db:          The database connection for performing CRUD operations related to proofs.
//   - lotusClient: The Lotus client for interacting with the Filecoin chain.
//   - request:     The request object which contains the sync criteria.
//
// Returns:
//   - An error indicating any issues that occurred during the sync operation.
func (DefaultHandler) SyncHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request SyncProofRequest) error {
	db = db.WithContext(ctx)

	if request.DealID != nil {
		// Sync proofs for a specific deal
		return syncProofsForDeal(ctx, db, lotusClient, *request.DealID)
	}

	if request.Provider != nil {
		// Sync proofs for a specific provider
		return syncProofsForProvider(ctx, db, lotusClient, *request.Provider, nil)
	}

	// Sync proofs for all active deals
	return syncAllProofs(ctx, db, lotusClient)
}

func syncProofsForDeal(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, dealID uint64) error {
	// Get the deal from database
	var deal model.Deal
	err := db.Where("deal_id = ?", dealID).First(&deal).Error
	if err != nil {
		return errors.Wrapf(err, "failed to find deal with ID %d", dealID)
	}

	// Search for messages related to this deal's provider
	return syncProofsForProvider(ctx, db, lotusClient, deal.Provider, &dealID)
}

func syncAllProofs(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error {
	// Get all active deals
	var deals []model.Deal
	err := db.Where("state IN ?", []string{
		string(model.DealActive),
		string(model.DealPublished),
	}).Find(&deals).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Group deals by provider to avoid duplicate work
	providerDeals := make(map[string][]uint64)
	for _, deal := range deals {
		if deal.DealID != nil {
			providerDeals[deal.Provider] = append(providerDeals[deal.Provider], *deal.DealID)
		}
	}

	// Sync proofs for each provider
	for provider := range providerDeals {
		if err := syncProofsForProvider(ctx, db, lotusClient, provider, nil); err != nil {
			// Log error but continue with other providers
			fmt.Printf("Error syncing proofs for provider %s: %v\n", provider, err)
		}
	}

	return nil
}

func syncProofsForProvider(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, provider string, dealID *uint64) error {
	// Get current chain head
	var chainHead ChainHead
	err := lotusClient.CallFor(ctx, &chainHead, "Filecoin.ChainHead")
	if err != nil {
		return errors.Wrap(err, "failed to get chain head")
	}

	// Look back 2000 epochs (about 16 hours) for proof messages
	lookbackEpochs := int64(2000)
	fromHeight := chainHead.Height - lookbackEpochs
	if fromHeight < 0 {
		fromHeight = 0
	}

	// Search for messages from this provider
	messageFilter := map[string]interface{}{
		"From": provider,
	}

	var messageCids []struct {
		CID string `json:"/"`
	}
	err = lotusClient.CallFor(ctx, &messageCids, "Filecoin.StateListMessages",
		messageFilter, chainHead.Cids, fromHeight)
	if err != nil {
		return errors.Wrap(err, "failed to list messages")
	}

	// Process each message
	for _, msgCid := range messageCids {
		if err := processMessage(ctx, db, lotusClient, msgCid.CID, provider, dealID); err != nil {
			// Log error but continue processing other messages
			fmt.Printf("Error processing message %s: %v\n", msgCid.CID, err)
		}
	}

	return nil
}

func processMessage(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, messageCid, provider string, dealID *uint64) error {
	// Get message details
	var msg Message
	err := lotusClient.CallFor(ctx, &msg, "Filecoin.ChainGetMessage", map[string]string{
		"/": messageCid,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get message")
	}

	// Check if this is a proof-related message
	var proofType model.ProofType
	var method string

	switch msg.Method {
	case MethodSubmitWindowedPoSt:
		proofType = model.ProofOfSpacetime
		method = "SubmitWindowedPoSt"
	case MethodProveCommitSector:
		proofType = model.ProofOfReplication
		method = "ProveCommitSector"
	case MethodPreCommitSector:
		proofType = model.ProofOfReplication
		method = "PreCommitSector"
	default:
		// Not a proof message, skip
		return nil
	}

	// Get message execution result
	var invocResult InvocResult
	err = lotusClient.CallFor(ctx, &invocResult, "Filecoin.StateReplay",
		nil, map[string]string{"/": messageCid})
	if err != nil {
		return errors.Wrap(err, "failed to replay message")
	}

	// Search for the message to get the tipset and height
	var msgLookup struct {
		Message struct {
			CID string `json:"/"`
		} `json:"Message"`
		Receipt MessageReceipt `json:"Receipt"`
		TipSet  TipSet         `json:"TipSet"`
		Height  int64          `json:"Height"`
	}

	err = lotusClient.CallFor(ctx, &msgLookup, "Filecoin.StateSearchMsg",
		nil, map[string]string{"/": messageCid}, 2000, true)
	if err != nil {
		return errors.Wrap(err, "failed to search message")
	}

	// Extract sector ID from message params if possible
	var sectorID *uint64
	if len(msg.Params) > 0 {
		// This is a simplified parsing - in practice you'd need to decode the CBOR params
		// For now, we'll leave it as nil
	}

	// Create block CID from tipset
	var blockCID string
	if len(msgLookup.TipSet) > 0 {
		blockCID = msgLookup.TipSet[0].CID
	}

	// Check if proof already exists
	var existingProof model.Proof
	err = db.Where("message_id = ?", messageCid).First(&existingProof).Error
	if err == nil {
		// Proof already exists, skip
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "failed to check existing proof")
	}

	// Create new proof record
	proof := model.Proof{
		DealID:    dealID,
		ProofType: proofType,
		MessageID: messageCid,
		BlockCID:  blockCID,
		Height:    msgLookup.Height,
		Method:    method,
		Verified:  invocResult.MsgRct.ExitCode == 0,
		SectorID:  sectorID,
		Provider:  provider,
		ErrorMsg:  invocResult.Error,
	}

	// Save proof to database
	err = db.Create(&proof).Error
	if err != nil {
		return errors.Wrap(err, "failed to create proof record")
	}

	return nil
}

// @ID SyncProofs
// @Summary Sync proofs from Filecoin chain
// @Description Synchronize proofs from the Filecoin chain into the database
// @Tags Proof
// @Accept json
// @Produce json
// @Param request body SyncProofRequest true "SyncProofRequest"
// @Success 200 {object} string "success"
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /proof/sync [post]
func _() {}
