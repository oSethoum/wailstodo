package db

import (
	"app/backend/config"
	"app/backend/models"

	"fmt"
	"log"
	"os"

	"github.com/oSethoum/namegorm"
	sqlite "github.com/oSethoum/sqlite3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Client *gorm.DB
)

func Connect() {
	dbname := fmt.Sprintf("db.sqlite?_fk=1&_pragma_key=%s&_pragma_cipher_page_size=4096", config.Secret)
	client, _ := gorm.Open(sqlite.Open(dbname), &gorm.Config{

		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
			}),

		NamingStrategy: namegorm.NamingStrategy{
			ColumnNameCase: namegorm.CamelCase,
		},
	})

	client.AutoMigrate(
		&models.Todo{},
	)

	Client = client
}
