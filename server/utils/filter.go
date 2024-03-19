package utils

import (
	"strconv"

	"github.com/aman-lf/sales-server/model"
)

func GetPipelineFilter(limitStr, offsetStr, sortByStr, defaultSort, sortOrderStr, searchText string) *model.PipelineParams {
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
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
		Offset:     offset,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
		SearchText: searchText,
	}
}
