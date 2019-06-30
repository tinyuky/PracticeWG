package entities

type Products struct {
	ID        int64  `db:"id_product" json:"id,omitempty"`
	Category  int    `db:"fk_category" json:"category_id"`
	Name      string `db:"name" json:"name,omitempty"`
	Image     string `db:"image" json"image,omitempty"`
	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
}
