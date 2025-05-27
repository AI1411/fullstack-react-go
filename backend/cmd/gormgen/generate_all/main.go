package main

import (
	"gorm.io/gen"

	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	applogger "github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/domain/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler(
		"postgres",
		"postgres",
		"localhost",
		"5432",
		"gen",
		applogger.New(applogger.DefaultConfig()),
	)
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn)

	// schema_migrationsを除外してテーブル生成
	tables, err := sqlHandler.Conn.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	// schema_migrationsを除外
	var filteredTables []interface{}
	for _, tableName := range tables {
		if tableName != "schema_migrations" {
			model := g.GenerateModel(tableName)
			filteredTables = append(filteredTables, model)
		}
	}

	g.ApplyBasic(filteredTables...)

	// Generate the code
	g.Execute()
}
