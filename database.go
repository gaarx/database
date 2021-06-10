package database

import (
	"github.com/gaarx/gaarx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

func GetConnString(user, pass, host, port, dbName string) string {
	return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4,utf8&parseTime=true&sql_mode=ansi"
}

// WithDatabase is option for gaarx and include Database
func WithDatabase(conn string, entities ...interface{}) gaarx.Option {
	// TODO: get rid of mySQL
	return func(app *gaarx.App) error {
		db, err := gorm.Open(
			"mysql",
			conn,
		)
		if err != nil {
			log.Fatal().Err(err)
			panic(err)
		}
		db.Set("gorm:table_options", "CHARSET=utf8")
		for _, e := range entities {
			db.AutoMigrate(e)
		}
		app.SetDatabase(db)
		return nil
	}
}

// WithDatabaseNoMigrate is option for gaarx and include Database without auto migrate entities
func WithDatabaseNoMigrate(conn string) gaarx.Option {
	return func(app *gaarx.App) error {
		db, err := gorm.Open(
			"mysql",
			conn,
		)
		if err != nil {
			log.Fatal().Err(err)
			panic(err)
		}
		db.Set("gorm:table_options", "CHARSET=utf8")
		app.SetDatabase(db)
		return nil
	}
}
