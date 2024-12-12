package request

type PaginationRequestPayload struct {
	Page  int `json:"page" query:"page"`
	Limit int `json:"limit" query:"limit"`
}

func (p PaginationRequestPayload) GenerateDefaultValue() PaginationRequestPayload {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	return p
}
