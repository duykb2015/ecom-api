package repo

import (
	"database/sql"

	"github.com/duykb2015/login-api/internal/entity"
)

type ProductRepo struct {
	*sql.DB
}

func New(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (p *ProductRepo) GetAllProduct() ([]entity.Product, error) {
	rows, err := p.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.Product, 0)

	for rows.Next() {
		e := entity.Product{}

		err = rows.Scan(&e.ID, &e.Name, &e.Slug, &e.ShortDescriptions, &e.Descriptions, &e.Images, &e.Price, &e.Quantity, &e.Status, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entities = append(entities, e)
	}

	return entities, nil
}
