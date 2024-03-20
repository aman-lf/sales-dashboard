package utils

import (
	"strconv"

	"github.com/aman-lf/sales-server/model"
)

func GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText string) *model.PipelineParams {
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 20
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	sortBy := sortByStr
	if sortBy == "" {
		sortBy = defaultSort
	}

	var sortOrder model.SortOrder
	sortOrder = 1
	if sortOrderStr == "-1" {
		sortOrder = -1
	}

	return &model.PipelineParams{
		Limit:      limit,
		Page:       page,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
		SearchText: searchText,
	}
}
