package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateTransactionTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.Transaction{}) {
		if err := tx.AutoMigrate(&pgmodels.Transaction{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982286_create_expenses")
	}

	return nil
}

func RollbackCreateTransactionTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.Transaction{})
	if err != nil {
		return err
	}

	return nil
}
