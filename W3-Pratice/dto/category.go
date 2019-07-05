package dto

type Category struct {
	ID    int64 `json:"id,omitempty"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
