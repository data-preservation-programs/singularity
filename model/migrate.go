package model

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Tables = []interface{}{
	&Global{},
	&Worker{},
	&Dataset{},
	&Source{},
	&Chunk{},
	&Item{},
	&Directory{},
	&Car{},
	&RawBlock{},
	&ItemBlock{},
	&CarBlock{},
	&Deal{},
	&Schedule{},
	&Wallet{},
	&WalletAssignment{},
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(Tables...)
	if err != nil {
		return errors.Wrap(err, "failed to auto migrate")
	}

	err = db.FirstOrCreate(&Global{}, Global{Key: "instance_id", Value: uuid.NewString()}).Error
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

	err = db.FirstOrCreate(&Global{}, row).Error
	if err != nil {
		return errors.Wrap(err, "failed to create salt")
	}

	return nil
}

func DropAll(db *gorm.DB) error {
	for _, table := range Tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}
