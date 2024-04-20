package databases

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConn *sql.DB
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error executing migration")
		panic(err)
	}
	DbConn = dbParam
	fmt.Printf("Applied %d migrations!\n", n)
}
