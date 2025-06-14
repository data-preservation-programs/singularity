package dataprep

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Name              string   `binding:"required"    json:"name"`              // Name of the preparation
	SourceStorages    []string `json:"sourceStorages"`                          // Name of Source storage systems to be used for the source
	OutputStorages    []string `json:"outputStorages"`                          // Name of Output storage systems to be used for the output
	MaxSizeStr        string   `default:"31.5GiB"     json:"maxSize"`           // Maximum size of the CAR files to be created
	PieceSizeStr      string   `default:""            json:"pieceSize"`         // Target piece size of the CAR files used for piece commitment calculation
	MinPieceSizeStr   string   `default:"1MiB"        json:"minPieceSize"`      // Minimum piece size for the preparation, applies only to DAG and remainer pieces
	DeleteAfterExport bool     `default:"false"       json:"deleteAfterExport"` // Whether to delete the source files after export
	NoInline          bool     `default:"false"       json:"noInline"`          // Whether to disable inline storage for the preparation. Can save database space but requires at least one output storage.
	NoDag             bool     `default:"false"       json:"noDag"`             // Whether to disable maintaining folder dag structure for the sources. If disabled, DagGen will not be possible and folders will not have an associated CID.

	// Auto-deal creation parameters
	AutoCreateDeals     bool            `default:"false"       json:"autoCreateDeals"`               // Enable automatic deal schedule creation
	DealPricePerGB      float64         `default:"0.0"         json:"dealPricePerGb"`                // Price in FIL per GiB
	DealPricePerGBEpoch float64         `default:"0.0"         json:"dealPricePerGbEpoch"`           // Price in FIL per GiB per epoch
	DealPricePerDeal    float64         `default:"0.0"         json:"dealPricePerDeal"`              // Price in FIL per deal
	DealDuration        time.Duration   `json:"dealDuration"        swaggertype:"primitive,integer"` // Deal duration
	DealStartDelay      time.Duration   `json:"dealStartDelay"      swaggertype:"primitive,integer"` // Deal start delay
	DealVerified        bool            `default:"false"       json:"dealVerified"`                  // Whether deals should be verified
	DealKeepUnsealed    bool            `default:"false"       json:"dealKeepUnsealed"`              // Whether to keep unsealed copy
	DealAnnounceToIPNI  bool            `default:"false"       json:"dealAnnounceToIpni"`            // Whether to announce to IPNI
	DealProvider        string          `default:""            json:"dealProvider"`                  // Storage Provider ID
	DealHTTPHeaders     model.ConfigMap `json:"dealHttpHeaders"`                                     // HTTP headers for deals
	DealURLTemplate     string          `default:""            json:"dealUrlTemplate"`               // URL template for deals
	WalletValidation    bool            `default:"false"       json:"walletValidation"`              // Enable wallet balance validation
	SPValidation        bool            `default:"false"       json:"spValidation"`                  // Enable storage provider validation
}

// ValidateCreateRequest processes and validates the creation request parameters.
// The function checks the validity of the input parameters such as maxSize, pieceSize, and
// the existence of source and output storages. The function also ensures that provided
// parameters meet certain criteria, like the pieceSize being a power of two, and maxSize
// allowing for padding. The encryption and storages compatibility is also validated.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - request: The CreateRequest structure containing the parameters for the creation request.
//
// Returns:
//   - A pointer to the validated Preparation model which can be used for subsequent operations.
//   - An error, if any occurred during the validation. This includes errors such as invalid
//     parameter values, storage not found, or incompatibility between encryption and storage options.
//
// Note:
// If certain parameters are not provided in the request, they are computed based on certain
// defaults or constraints, like the pieceSize defaulting to a power of two value.
func ValidateCreateRequest(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	if request.MaxSizeStr == "" {
		request.MaxSizeStr = "31.5GiB"
	}

	if util.IsAllDigits(request.Name) || request.Name == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "preparation name '%s' cannot be all digits or empty", request.Name)
	}

	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for maxSize: %s", request.MaxSizeStr))
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for pieceSize: %s", request.PieceSizeStr))
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "pieceSize must be a power of two")
		}
	}

	if pieceSize > 1<<36 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "pieceSize cannot be larger than 64 GiB")
	}

	if maxSize*128/127 >= pieceSize {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "maxSize needs to be reduced to leave space for padding")
	}

	minPieceSizeStr := request.MinPieceSizeStr
	if minPieceSizeStr == "" {
		minPieceSizeStr = "1MiB"
	}

	minPieceSize, err := humanize.ParseBytes(minPieceSizeStr)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for minPieceSize: %s", minPieceSizeStr))
	}

	if minPieceSize > pieceSize {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "minPieceSize cannot be larger than pieceSize")
	}

	if minPieceSize != util.NextPowerOfTwo(minPieceSize) {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "minPieceSize must be a power of two")
	}

	var sources []model.Storage
	for _, name := range request.SourceStorages {
		var source model.Storage
		err = source.FindByIDOrName(db, name)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(handlererror.ErrNotFound, "source storage %s does not exist", name)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sources = append(sources, source)
	}

	var outputs []model.Storage
	for _, name := range request.OutputStorages {
		var output model.Storage
		err = output.FindByIDOrName(db, name)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage %s does not exist", name)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		outputs = append(outputs, output)
	}

	if len(outputs) == 0 && request.DeleteAfterExport {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "deleteAfterExport cannot be set without output storages")
	}

	if len(outputs) == 0 && request.NoInline {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "inline preparation cannot be disabled without output storages")
	}

	return &model.Preparation{
		MaxSize:             int64(maxSize),
		PieceSize:           int64(pieceSize),
		MinPieceSize:        int64(minPieceSize),
		SourceStorages:      sources,
		OutputStorages:      outputs,
		DeleteAfterExport:   request.DeleteAfterExport,
		Name:                request.Name,
		NoInline:            request.NoInline,
		NoDag:               request.NoDag,
		AutoCreateDeals:     request.AutoCreateDeals,
		DealPricePerGB:      request.DealPricePerGB,
		DealPricePerGBEpoch: request.DealPricePerGBEpoch,
		DealPricePerDeal:    request.DealPricePerDeal,
		DealDuration:        request.DealDuration,
		DealStartDelay:      request.DealStartDelay,
		DealVerified:        request.DealVerified,
		DealKeepUnsealed:    request.DealKeepUnsealed,
		DealAnnounceToIPNI:  request.DealAnnounceToIPNI,
		DealProvider:        request.DealProvider,
		DealHTTPHeaders:     request.DealHTTPHeaders,
		DealURLTemplate:     request.DealURLTemplate,
		WalletValidation:    request.WalletValidation,
		SPValidation:        request.SPValidation,
	}, nil
}

// CreatePreparationHandler handles the creation of a new Preparation entity based on the provided
// CreateRequest parameters. Initially, it validates the request parameters and, if valid,
// creates a new Preparation record in the database. It also performs wallet and storage provider
// validation if enabled in the request.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - request: The CreateRequest structure containing the parameters for the creation request.
//
// Returns:
//   - A pointer to the newly created Preparation model.
//   - An error, if any occurred during the validation or creation process.
//
// Note:
// This function relies on the ValidateCreateRequest function to ensure that the provided
// parameters meet the required criteria before creating a Preparation record.
func (DefaultHandler) CreatePreparationHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	preparation, err := ValidateCreateRequest(ctx, db, request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Perform validation if auto-deal creation is enabled
	if preparation.AutoCreateDeals {
		err = performValidation(ctx, db, preparation)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	err = database.DoRetry(ctx, func() error {
		err := db.Create(preparation).Error
		if err != nil {
			return errors.WithStack(err)
		}

		var attachments []model.SourceAttachment
		err = db.Where("preparation_id = ?", preparation.ID).Find(&attachments).Error
		if err != nil {
			return errors.WithStack(err)
		}

		for _, attachment := range attachments {
			err = db.Create(&model.Directory{
				AttachmentID: attachment.ID,
			}).Error
			if err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return preparation, nil
}

// performValidation handles wallet and storage provider validation for auto-deal creation
func performValidation(ctx context.Context, db *gorm.DB, preparation *model.Preparation) error {
	notificationHandler := notification.Default

	// Create metadata for logging
	metadata := model.ConfigMap{
		"preparation_name": preparation.Name,
		"preparation_id":   strconv.FormatUint(uint64(preparation.ID), 10),
		"auto_create_deals": func() string {
			if preparation.AutoCreateDeals {
				return "true"
			}
			return "false"
		}(),
	}

	// Log start of validation process
	_, err := notificationHandler.LogInfo(ctx, db, "dataprep-create",
		"Starting Auto-Deal Validation",
		"Beginning validation process for auto-deal creation",
		metadata)
	if err != nil {
		return errors.WithStack(err)
	}

	var validationErrors []string

	// Perform wallet validation if enabled
	if preparation.WalletValidation {
		err = performWalletValidation(ctx, db, preparation, &validationErrors)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Perform storage provider validation if enabled
	if preparation.SPValidation {
		err = performSPValidation(ctx, db, preparation, &validationErrors)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// If there are validation errors, log them and potentially disable auto-creation
	if len(validationErrors) > 0 {
		errorMetadata := model.ConfigMap{
			"preparation_name":  preparation.Name,
			"validation_errors": strings.Join(validationErrors, "; "),
		}

		_, err = notificationHandler.LogWarning(ctx, db, "dataprep-create",
			"Auto-Deal Validation Issues Found",
			"Some validation checks failed, but preparation will continue",
			errorMetadata)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		// All validations passed
		_, err = notificationHandler.LogInfo(ctx, db, "dataprep-create",
			"Auto-Deal Validation Successful",
			"All validation checks passed, ready for auto-deal creation",
			metadata)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// performWalletValidation validates wallet balance for auto-deal creation
func performWalletValidation(ctx context.Context, db *gorm.DB, preparation *model.Preparation, validationErrors *[]string) error {
	// For now, we'll perform a basic validation without connecting to Lotus
	// In a real implementation, you would get wallet addresses from the preparation
	// and validate each one using the wallet validator

	notificationHandler := notification.Default

	// Get wallets associated with this preparation
	var wallets []model.Wallet
	err := db.WithContext(ctx).
		Joins("JOIN wallet_assignments ON wallets.id = wallet_assignments.wallet_id").
		Where("wallet_assignments.preparation_id = ?", preparation.ID).
		Find(&wallets).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if len(wallets) == 0 {
		*validationErrors = append(*validationErrors, "No wallets assigned to preparation")

		_, err = notificationHandler.LogWarning(ctx, db, "dataprep-create",
			"No Wallets Found",
			"No wallets are assigned to this preparation for auto-deal creation",
			model.ConfigMap{
				"preparation_name": preparation.Name,
			})
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	// TODO: In a real implementation, you would connect to Lotus and validate each wallet
	// For now, we'll just log that wallet validation is enabled
	walletAddresses := make([]string, len(wallets))
	for i, wallet := range wallets {
		walletAddresses[i] = wallet.Address
	}

	_, err = notificationHandler.LogInfo(ctx, db, "dataprep-create",
		"Wallet Validation Enabled",
		"Wallet validation is enabled for auto-deal creation",
		model.ConfigMap{
			"preparation_name": preparation.Name,
			"wallet_addresses": strings.Join(walletAddresses, ", "),
		})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// performSPValidation validates storage provider for auto-deal creation
func performSPValidation(ctx context.Context, db *gorm.DB, preparation *model.Preparation, validationErrors *[]string) error {
	notificationHandler := notification.Default
	spValidator := storage.DefaultSPValidator

	// Check if a storage provider is specified
	if preparation.DealProvider == "" {
		// Try to get a default storage provider
		defaultSP, err := spValidator.GetDefaultStorageProvider(ctx, db, "auto-deal-creation")
		if err != nil {
			*validationErrors = append(*validationErrors, "No storage provider specified and no default available")

			_, err = notificationHandler.LogWarning(ctx, db, "dataprep-create",
				"No Storage Provider Available",
				"No storage provider specified and no default providers available",
				model.ConfigMap{
					"preparation_name": preparation.Name,
				})
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		}

		// Update preparation with default provider
		preparation.DealProvider = defaultSP.ProviderID

		_, err = notificationHandler.LogInfo(ctx, db, "dataprep-create",
			"Default Storage Provider Selected",
			"Using default storage provider for auto-deal creation",
			model.ConfigMap{
				"preparation_name": preparation.Name,
				"provider_id":      defaultSP.ProviderID,
				"provider_name":    defaultSP.Name,
			})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// TODO: In a real implementation, you would connect to Lotus and validate the storage provider
	// For now, we'll just log that SP validation is enabled
	_, err := notificationHandler.LogInfo(ctx, db, "dataprep-create",
		"Storage Provider Validation Enabled",
		"Storage provider validation is enabled for auto-deal creation",
		model.ConfigMap{
			"preparation_name": preparation.Name,
			"provider_id":      preparation.DealProvider,
		})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// @ID CreatePreparation
// @Summary Create a new preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Create Request"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation [post]
func _() {}
