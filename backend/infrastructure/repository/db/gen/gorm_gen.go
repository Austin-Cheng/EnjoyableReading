package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

var ExcludeTables = []string{
	"schema_migrations",
}

type Generator struct {
	dsn     string
	outPath string
	db      *gorm.DB
	g       *gen.Generator
	opts    []gen.ModelOpt
}

func NewGenerator(opts ...gen.ModelOpt) *Generator {
	host := os.Getenv("host")
	port := os.Getenv("port")
	username := os.Getenv("username")
	password := os.Getenv("password")
	database := os.Getenv("database")
	outpath := os.Getenv("outpath")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	g := &Generator{
		dsn:     dsn,
		outPath: outpath,
		opts:    opts,
	}
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	g.db = db
	g.g = gen.NewGenerator(gen.Config{
		OutPath:        g.outPath,
		FieldNullable:  true,
		FieldCoverable: true,
		FieldSignable:  true,
	})
	g.g.UseDB(db)
	return g
}

func (g *Generator) genModels(ts ...string) []any {
	excludeTablesSet := make(map[string]struct{})
	for _, tableName := range ExcludeTables {
		excludeTablesSet[tableName] = struct{}{}
	}
	tableModels := make([]interface{}, 0, len(ts))
	for _, tableName := range ts {
		if _, ok := excludeTablesSet[tableName]; ok {
			continue
		}
		tableModels = append(tableModels, g.g.GenerateModel(tableName, g.opts...))
	}
	return tableModels
}

func (g *Generator) GenTables(ts ...string) {
	models := g.genModels(ts...)
	if len(models) > 0 {
		fmt.Println("OK")
	}
}

func main() {
	ts := []string{"tags"}
	g := NewGenerator(gen.WithMethod(GenIDMethod{}))
	g.GenTables(ts...)
}
