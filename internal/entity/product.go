package entity

type Product struct {
	ID                string `json:"id" db:"id"`
	Name              string `json:"name" db:"name"`
	Slug              string `json:"slug" db:"slug"`
	Descriptions      string `json:"descriptions" db:"descriptions"`
	ShortDescriptions string `json:"short_descriptions" db:"short_descriptions"`
	Images            string `json:"images" db:"images"`
	Price             int    `json:"price" db:"price"`
	Quantity          int    `json:"quantity" db:"quantity"`
	Status            int    `json:"status" db:"status"`
	CreatedAt         string `json:"created_at" db:"created_at"`
	UpdatedAt         string `json:"updated_at" db:"updated_at"`
}
