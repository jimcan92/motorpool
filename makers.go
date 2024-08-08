package main

import (
	"encoding/json"

	"go.etcd.io/bbolt"
)

func (app *App) CreateMaker(maker Maker) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Makers"))
		encoded, err := json.Marshal(maker)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(maker.ID), encoded)
	})
}

func (app *App) GetMaker(id string) (Maker, error) {
	var maker Maker
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Makers"))
		data := bucket.Get([]byte(id))
		if data == nil {
			return nil
		}
		return json.Unmarshal(data, &maker)
	})
	return maker, err
}

func (app *App) GetAllMakers() ([]Maker, error) {
	var makers []Maker
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Makers"))
		return bucket.ForEach(func(k, v []byte) error {
			var maker Maker
			if err := json.Unmarshal(v, &maker); err != nil {
				return err
			}
			makers = append(makers, maker)
			return nil
		})
	})
	return makers, err
}

func (app *App) UpdateMaker(maker Maker) error {
	return app.CreateMaker(maker)
}

func (app *App) DeleteMaker(id string) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Makers"))
		return bucket.Delete([]byte(id))
	})
}
