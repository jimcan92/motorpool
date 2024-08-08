package main

import (
	"context"
	"log"

	"go.etcd.io/bbolt"
)

// App struct
type App struct {
	ctx context.Context
	db  *bbolt.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	db, err := bbolt.Open("./data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a bucket for storing models
	db.Update(func(tx *bbolt.Tx) error {
		_, modelErr := tx.CreateBucketIfNotExists([]byte("Models"))
		if modelErr != nil {
			return modelErr
		}
		_, makerErr := tx.CreateBucketIfNotExists([]byte("Makers"))
		if makerErr != nil {
			return makerErr
		}
		_, vehicleError := tx.CreateBucketIfNotExists([]byte("Vehicles"))
		return vehicleError
	})

	return &App{db: db}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
