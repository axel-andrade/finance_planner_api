package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateExpensesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.Expense{}) {
		if err := tx.AutoMigrate(&pgmodels.Expense{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982286_create_expenses")
	}

	return nil
}

func RollbackCreateExpensesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.Expense{})
	if err != nil {
		return err
	}

	return nil
}
