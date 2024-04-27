package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateIncomesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.Income{}) {
		if err := tx.AutoMigrate(&pgmodels.Income{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982288_create_incomes")
	}

	return nil
}

func RollbackCreateIncomesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.Income{})
	if err != nil {
		return err
	}

	return nil
}
