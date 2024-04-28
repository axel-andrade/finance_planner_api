package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateMonthlyFinancesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.MonthlyFinance{}) {
		if err := tx.AutoMigrate(&pgmodels.MonthlyFinance{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982285_monthly_finances")
	}

	return nil
}

func RollbackCreateMonthlyFinancesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.MonthlyFinance{})
	if err != nil {
		return err
	}

	return nil
}
