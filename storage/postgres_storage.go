package storage

import (
	"be-female-daily/model"
	// "github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	// DB *sqlx.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	// db, err := sqlx.Connect("postgres", dsn)
	// if err != nil {
	// 	return nil, err
	// }

	return &PostgresStorage{
		// DB: db,
	}, nil
}

func (p *PostgresStorage) GetAllData() (model.DataStorage, error) {
	// Implement the logic to retrieve all data from PostgreSQL
	return model.DataStorage{}, nil
}
func (p *PostgresStorage) SaveAllData(data model.DataStorage) error {
	// Implement the logic to save all data to PostgreSQL
	return nil
}
func (p *PostgresStorage) GetBrands() ([]model.Brand, error) {
	// Implement the logic to retrieve brands from PostgreSQL
	return []model.Brand{}, nil
}
func (p *PostgresStorage) SaveBrand(brand model.Brand) (model.Brand, error) {
	// Implement the logic to save a brand to PostgreSQL
	return brand, nil
}
func (p *PostgresStorage) GetMission() ([]model.Mission, error) {
	// Implement the logic to retrieve missions from PostgreSQL
	return []model.Mission{}, nil
}
func (p *PostgresStorage) SaveMission(mission model.Mission) (model.Mission, error) {
	// Implement the logic to save a mission to PostgreSQL
	return mission, nil
}
