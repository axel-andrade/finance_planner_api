package migrations

import (
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := MigrateCreateUsersTable(db); err != nil {
		RollbackCreateUsersTable(db)
		log.Fatal(err)
	}

	if err := MigrateCreateCategoriesTable(db); err != nil {
		RollbackCreateCategoriesTable(db)
		log.Fatal(err)
	}
}
