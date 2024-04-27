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

	if err := MigrateCreateExpenseTypesTable(db); err != nil {
		RollbackCreateExpenseTypesTable(db)
		log.Fatal(err)
	}

	if err := MigrateCreateExpensesTable(db); err != nil {
		RollbackCreateExpensesTable(db)
		log.Fatal(err)
	}

	if err := MigrateCreateIncomeTypesTable(db); err != nil {
		RollbackCreateIncomeTypesTable(db)
		log.Fatal(err)
	}

	if err := MigrateCreateIncomesTable(db); err != nil {
		RollbackCreateIncomesTable(db)
		log.Fatal(err)
	}
}
