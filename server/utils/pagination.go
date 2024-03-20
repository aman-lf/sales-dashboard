package utils

type PaginationData struct {
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
}

func GetOffsetValue(perPage, page int64) int64 {
	return int64(perPage * page)
}
