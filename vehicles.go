package main

import (
	"encoding/json"

	"go.etcd.io/bbolt"
)

func (app *App) CreateVehicle(model Vehicle) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Vehicles"))
		encoded, err := json.Marshal(model)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(model.ID), encoded)
	})
}

func (app *App) GetVehicle(id string) (Vehicle, error) {
	var model Vehicle
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Vehicles"))
		data := bucket.Get([]byte(id))
		if data == nil {
			return nil
		}
		return json.Unmarshal(data, &model)
	})
	return model, err
}

func (app *App) GetAllVehicles() ([]Vehicle, error) {
	var models []Vehicle
	err := app.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Vehicles"))
		return bucket.ForEach(func(k, v []byte) error {
			var model Vehicle
			if err := json.Unmarshal(v, &model); err != nil {
				return err
			}
			models = append(models, model)
			return nil
		})
	})
	return models, err
}

func (app *App) UpdateVehicle(model Vehicle) error {
	return app.CreateVehicle(model)
}

func (app *App) DeleteVehicle(id string) error {
	return app.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Vehicles"))
		return bucket.Delete([]byte(id))
	})
}
