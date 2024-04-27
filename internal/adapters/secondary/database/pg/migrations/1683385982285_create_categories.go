package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateCategoriesTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.Category{}) {
		if err := tx.AutoMigrate(&pgmodels.Category{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982285_create_categories")
	}

	return nil
}

func RollbackCreateCategoriesTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.Category{})
	if err != nil {
		return err
	}

	return nil
}
