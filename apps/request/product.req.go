package request

type CreateProductRequestPayload struct {
	Name         string `json:"name"`
	ID_Categorie int    `json:"id_categorie"`
	Stock        int16  `json:"stock"`
	Price        int    `json:"price"`
}

type ListProductRequestPayload struct {
	// cursor berfungsi untuk last id pada page tersebut atau bisa juga created_at
	Cursor       int `json:"cursor" query:"cursor"`
	Size         int `json:"size" query:"size"`
	CategoriesID int `json:"categories_id" query:"categories_id"`
}

func (l ListProductRequestPayload) GenerateDefaultValue() ListProductRequestPayload {
	if l.Cursor < 0 {
		l.Cursor = 0
	}
	if l.Size <= 0 {
		l.Size = 10
	}
	return l
}
