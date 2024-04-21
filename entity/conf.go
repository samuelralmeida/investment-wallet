package entity

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	SQLITE Sqlite
}

type Sqlite struct {
	FILENAME string
}

func LoadConf() *Conf {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file:", err)
	}

	return &Conf{
		SQLITE: Sqlite{
			FILENAME: os.Getenv("SQLITE_FILENAME"),
		},
	}
}
