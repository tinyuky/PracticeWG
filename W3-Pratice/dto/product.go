package dto

type Product struct {
	ID        int64       `json:"id,omitempty"`
	Category  interface{} `json:category`
	Name      string      `json:"name,omitempty"`
	Image     string      `json"image,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
}
