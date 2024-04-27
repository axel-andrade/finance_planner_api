package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateIncomeTypesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.ExpenseType{}) {
		if err := tx.AutoMigrate(&pgmodels.ExpenseType{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982287_income_types")
	}

	return nil
}

func RollbackCreateIncomeTypesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.IncomeType{})
	if err != nil {
		return err
	}

	return nil
}
