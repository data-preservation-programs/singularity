package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Tables = []any{
	&Global{},
	&Worker{},
	&Dataset{},
	&Source{},
	&Chunk{},
	&Item{},
	&Directory{},
	&Car{},
	&CarBlock{},
	&Deal{},
	&Schedule{},
	&Wallet{},
	&WalletAssignment{},
	&ItemPart{},
}

var logger = logging.Logger("model")

func AutoMigrate(db *gorm.DB) error {
	logger.Info("Auto migrating tables")
	err := db.AutoMigrate(Tables...)
	if err != nil {
		return errors.Wrap(err, "failed to auto migrate")
	}

	logger.Debug("Creating instance id")
	err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&Global{Key: "instance_id", Value: uuid.NewString()}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create instance id")
	}

	salt := make([]byte, 8)
	_, err = rand.Read(salt)
	if err != nil {
		return errors.Wrap(err, "failed to generate salt")
	}
	encoded := base64.StdEncoding.EncodeToString(salt)
	row := Global{
		Key:   "salt",
		Value: encoded,
	}

	logger.Debug("Creating encryption salt")
	err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(row).Error
	if err != nil {
		return errors.Wrap(err, "failed to create salt")
	}

	return nil
}

func DropAll(db *gorm.DB) error {
	logger.Info("Dropping all tables")
	for _, table := range Tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}
