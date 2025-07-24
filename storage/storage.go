package storage

import "be-female-daily/model"

type Storage interface {
	GetBrands() ([]model.Brand, error)
	SaveBrand(brand model.Brand) (model.Brand, error)
	GetMission() ([]model.Mission, error)
	SaveMission(mission model.Mission) (model.Mission, error)
	GetAllData() (model.DataStorage, error)
	SaveAllData(data model.DataStorage) error
}
