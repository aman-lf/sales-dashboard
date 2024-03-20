package utils

import (
	"reflect"
	"testing"

	"github.com/aman-lf/sales-server/model"
)

func TestEmptyPipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := ""
	sortByStr := ""
	defaultSort := ""
	sortOrderStr := ""
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       1,
		SortBy:     "",
		SortOrder:  1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("EmptyPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestLimitPipelineFilter(t *testing.T) {
	limitStr := "5"
	pageStr := ""
	sortByStr := ""
	defaultSort := ""
	sortOrderStr := ""
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      5,
		Page:       1,
		SortBy:     "",
		SortOrder:  1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("LimitPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestPagePipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := "2"
	sortByStr := ""
	defaultSort := ""
	sortOrderStr := ""
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       2,
		SortBy:     "",
		SortOrder:  1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("EmptyPagelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestSortByPipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := ""
	sortByStr := "test"
	defaultSort := ""
	sortOrderStr := ""
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       1,
		SortBy:     "test",
		SortOrder:  1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortByPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestDefaultSortPipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := ""
	sortByStr := ""
	defaultSort := "test"
	sortOrderStr := ""
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       1,
		SortBy:     "test",
		SortOrder:  1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DefaultSortPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestSortOrderPipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := ""
	sortByStr := ""
	defaultSort := ""
	sortOrderStr := "-1"
	searchText := ""
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       1,
		SortBy:     "",
		SortOrder:  -1,
		SearchText: "",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortOrderPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestSearchPipelineFilter(t *testing.T) {
	limitStr := ""
	pageStr := ""
	sortByStr := ""
	defaultSort := ""
	sortOrderStr := ""
	searchText := "test"
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      20,
		Page:       1,
		SortBy:     "",
		SortOrder:  1,
		SearchText: "test",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SearchPipelineFilter returned %+v, expected %+v", result, expected)
	}
}

func TestAllPipelineFilter(t *testing.T) {
	limitStr := "10"
	pageStr := "2"
	sortByStr := "name"
	defaultSort := "id"
	sortOrderStr := "-1"
	searchText := "testName"
	result := GetPipelineFilter(limitStr, pageStr, sortByStr, defaultSort, sortOrderStr, searchText)
	expected := &model.PipelineParams{
		Limit:      10,
		Page:       2,
		SortBy:     "name",
		SortOrder:  -1,
		SearchText: "testName",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AllPipelineFilter returned %+v, expected %+v", result, expected)
	}
}
