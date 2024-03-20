package model

import "math"

type SortOrder int

const (
	Ascending  SortOrder = 1
	Descending SortOrder = -1
)

type PipelineParams struct {
	Limit      int64     `json:"limit"`
	Offset     int64     `json:"offset"`
	Page       int64     `json:"page"`
	TotalPage  int64     `json:"totalPage"`
	SortBy     string    `json:"sortBy"`
	SortOrder  SortOrder `json:"sortOrder"`
	SearchText string    `json:"searchText"`
}

func (p *PipelineParams) CalculateTotalPageCount(count int) {
	totalPage := math.Ceil(float64(count) / float64(p.Limit))
	p.TotalPage = int64(math.Max(totalPage, 1))
}

func (p *PipelineParams) VerifyPage() {
	if p.Page > p.TotalPage {
		p.Page = p.TotalPage
	}
}

func (p *PipelineParams) CalculateOffset() {
	p.Offset = p.Limit * (p.Page - 1)
}
