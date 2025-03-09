package web

type PaginationQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type PaginationDTO struct {
	TotalCount int `json:"totalCount"`
	FirstRow   int `json:"firstRow,omitempty"`
	LastRow    int `json:"lastRow,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

func ToPaginationDTO(req interface{}, count, limit, offset int) *PaginationDTO {
	if req == nil || limit == 0 {
		return nil
	}

	firstRow := offset + 1
	lastRow := offset + int(count)
	if count == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(count) + limit - 1) / limit

	return &PaginationDTO{
		TotalCount: count,
		FirstRow:   firstRow,
		LastRow:    lastRow,
		TotalPages: totalPages,
	}
}
