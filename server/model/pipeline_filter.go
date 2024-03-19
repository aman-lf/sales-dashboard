package model

type SortOrder int

const (
	Ascending  SortOrder = 1
	Descending SortOrder = -1
)

type PipelineParams struct {
	Limit      int64     `json:"limit" form:"limit"`
	Offset     int64     `json:"offset" form:"offset"`
	SortBy     string    `json:"sortBy" form:"sortBy"`
	SortOrder  SortOrder `json:"sortOrder" form:"sortOrder"`
	SearchText string    `json:"searchText" form:"searchText"`
}
