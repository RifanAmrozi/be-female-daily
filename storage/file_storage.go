package storage

import (
	"be-female-daily/model"
	"encoding/json"
	"os"
)

type FileStorage struct {
	FilePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{FilePath: filePath}
}

func (fs *FileStorage) GetBrands() ([]model.Brand, error) {
	data, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return nil, err
	}
	var dataStorage model.DataStorage
	if len(data) == 0 {
		return dataStorage.Brands, nil // empty file, return empty slice
	}
	err = json.Unmarshal(data, &dataStorage)

	return dataStorage.Brands, err
}

func (fs *FileStorage) SaveBrand(brand model.Brand) (model.Brand, error) {
	data, err := fs.GetAllData()
	if err != nil {
		return model.Brand{}, err
	}
	brands := data.Brands
	flag := true
	// Check if brand already exists (optional)
	for i, b := range brands {
		if b.ID == brand.ID {
			brands[i] = brand
			flag = false // Brand exists, update it
			break
		}
	}

	if flag {
		brands = append(brands, brand)
	}
	dataStorage := model.DataStorage{
		Brands:   brands,
		Missions: data.Missions,
		Tickets:  data.Tickets,
		Users:    data.Users,
	}
	dataNew, err := json.MarshalIndent(dataStorage, "", "  ")
	if err != nil {
		return model.Brand{}, err
	}
	return brand, os.WriteFile(fs.FilePath, dataNew, 0644)
}

func (fs *FileStorage) GetMission() ([]model.Mission, error) {
	data, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return nil, err
	}

	var dataStorage model.DataStorage
	if len(data) == 0 {
		return dataStorage.Missions, nil // empty file, return empty slice
	}
	err = json.Unmarshal(data, &dataStorage)

	return dataStorage.Missions, err
}
func (fs *FileStorage) SaveMission(mission model.Mission) (model.Mission, error) {
	data, err := fs.GetAllData()
	if err != nil {
		return model.Mission{}, err
	}
	missions := data.Missions
	flag := true
	// Check if mission already exists (optional)
	for i, m := range missions {
		if m.ID == mission.ID {
			missions[i] = mission
			flag = false
			break
		}
	}
	if flag {
		missions = append(missions, mission)
	}
	dataStorage := model.DataStorage{
		Brands:   data.Brands,  // Assuming you want to keep brands intact
		Missions: missions,     // Updated missions
		Tickets:  data.Tickets, // Assuming you want to keep tickets intact
		Users:    data.Users,   // Assuming you want to keep users intact
	}
	dataNew, err := json.MarshalIndent(dataStorage, "", "  ")
	if err != nil {
		return model.Mission{}, err
	}
	return mission, os.WriteFile(fs.FilePath, dataNew, 0644)
}
func (fs *FileStorage) GetAllData() (model.DataStorage, error) {
	var dataStorage model.DataStorage
	data, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return dataStorage, err
	}
	if len(data) == 0 {
		return dataStorage, nil // empty file, return empty slice
	}
	err = json.Unmarshal(data, &dataStorage)

	return dataStorage, err
}
func (fs *FileStorage) SaveAllData(data model.DataStorage) error {
	dataNew, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, dataNew, 0644)
}
