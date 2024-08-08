package main

import (
	"encoding/json"

	"go.etcd.io/bbolt"
)

func (app *App) CreateModel(model Model) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Models"))
		encoded, err := json.Marshal(model)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(model.ID), encoded)
	})
}

func (app *App) GetModel(id string) (Model, error) {
	var model Model
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Models"))
		data := bucket.Get([]byte(id))
		if data == nil {
			return nil
		}
		return json.Unmarshal(data, &model)
	})
	return model, err
}

func (app *App) GetAllModels() ([]Model, error) {
	var models []Model
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Models"))
		return bucket.ForEach(func(k, v []byte) error {
			var model Model
			if err := json.Unmarshal(v, &model); err != nil {
				return err
			}
			models = append(models, model)
			return nil
		})
	})
	return models, err
}

func (app *App) UpdateModel(model Model) error {
	return app.CreateModel(model)
}

func (app *App) DeleteModel(id string) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Models"))
		return bucket.Delete([]byte(id))
	})
}
