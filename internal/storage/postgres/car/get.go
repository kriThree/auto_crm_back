package storage_car

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s CarStorage) GetByClientId(ctx context.Context, userId int64) ([]storage_models.Car, error) {

	const op = "postgres.car.Get"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, number,description,client_id
		FROM cars
		WHERE client_id = $1;
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var cars []storage_models.Car
	for rows.Next() {
		var car storage_models.Car
		err = rows.Scan(&car.Id, &car.Number,&car.Description,&car.ClientId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func (s CarStorage) GetById(ctx context.Context, id int64) (car storage_models.Car, err error) {

	const op = "postgres.car.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, number,description,client_id
		FROM cars
		WHERE id = $1;
	`)
	if err != nil {
		return storage_models.Car{}, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&car.Id, &car.ClientId,&car.Number,&car.Description,&car.ClientId)

	if err != nil {
		return storage_models.Car{}, fmt.Errorf("%s : %w", op, err)
	}

	return car, nil
}
