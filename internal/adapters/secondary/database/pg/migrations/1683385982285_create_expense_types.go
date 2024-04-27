package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateExpenseTypesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.ExpenseType{}) {
		if err := tx.AutoMigrate(&pgmodels.ExpenseType{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982285_expense_types")
	}

	return nil
}

func RollbackCreateExpenseTypesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.ExpenseType{})
	if err != nil {
		return err
	}

	return nil
}
