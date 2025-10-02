package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func _202509171710_wallet_assignments_cascade() *gormigrate.Migration {
	// set wallet_assignments fks to cascade on delete, otherwise preparations with wallets cannot be deleted
	return &gormigrate.Migration{
		ID: "202509171710",
		Migrate: func(tx *gorm.DB) error {
			if tx.Dialector.Name() == "postgres" {
				// switch to cascade for preparation and wallet fks in postgres
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP CONSTRAINT IF EXISTS fk_wallet_assignments_preparation").Error; err != nil {
					return errors.Wrap(err, "drop fk_wallet_assignments_preparation")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_preparation FOREIGN KEY (preparation_id) REFERENCES preparations(id) ON DELETE CASCADE").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_preparation cascade")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP CONSTRAINT IF EXISTS fk_wallet_assignments_wallet").Error; err != nil {
					return errors.Wrap(err, "drop fk_wallet_assignments_wallet")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_wallet FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_wallet cascade")
				}
				return nil
			}
			if tx.Dialector.Name() == "mysql" {
				// align mysql fks to cascade semantics
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP FOREIGN KEY fk_wallet_assignments_preparation").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_preparation FOREIGN KEY (preparation_id) REFERENCES preparations(id) ON DELETE CASCADE").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_preparation cascade (mysql)")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP FOREIGN KEY fk_wallet_assignments_wallet").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_wallet FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_wallet cascade (mysql)")
				}
				return nil
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if tx.Dialector.Name() == "postgres" {
				// revert to no action to match previous behavior
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP CONSTRAINT IF EXISTS fk_wallet_assignments_preparation").Error; err != nil {
					return errors.Wrap(err, "drop fk_wallet_assignments_preparation (rollback)")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_preparation FOREIGN KEY (preparation_id) REFERENCES preparations(id) ON DELETE NO ACTION").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_preparation no action (rollback)")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP CONSTRAINT IF EXISTS fk_wallet_assignments_wallet").Error; err != nil {
					return errors.Wrap(err, "drop fk_wallet_assignments_wallet (rollback)")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_wallet FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE NO ACTION").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_wallet no action (rollback)")
				}
				return nil
			}
			if tx.Dialector.Name() == "mysql" {
				// revert mysql fks back to no action equivalent
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP FOREIGN KEY fk_wallet_assignments_preparation").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_preparation FOREIGN KEY (preparation_id) REFERENCES preparations(id)").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_preparation no action (mysql rollback)")
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments DROP FOREIGN KEY fk_wallet_assignments_wallet").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE wallet_assignments ADD CONSTRAINT fk_wallet_assignments_wallet FOREIGN KEY (wallet_id) REFERENCES wallets(id)").Error; err != nil {
					return errors.Wrap(err, "add fk_wallet_assignments_wallet no action (mysql rollback)")
				}
				return nil
			}
			return nil
		},
	}
}
