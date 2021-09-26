package app

import (
	"encoding/gob"
	"log"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/quasoft/memstore"
)

var (
	Store sessions.Store
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	gob.Register(map[string]interface{}{})
	Store = memstore.NewMemStore([]byte("secret123"))
	return nil
}
