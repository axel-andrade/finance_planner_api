package migrations

import (
	"fmt"

	pgmodels "github.com/axel-andrade/finance_planner_api/internal/adapters/secondary/database/pg/models"
	"gorm.io/gorm"
)

func MigrateCreateUsersTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&pgmodels.User{}) {
		if err := tx.AutoMigrate(&pgmodels.User{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982284_create_users")
	}

	return nil
}

func RollbackCreateUsersTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&pgmodels.User{})
	if err != nil {
		return err
	}

	return nil
}
