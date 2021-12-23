package model

import (
	"fmt"
	"strings"
)

type ListRepositoriesParams struct {
	Query    string
	Sort     string
	Page     int64
	PageSize int64
}

func NewListRepositoriesParams() *ListRepositoriesParams {
	return &ListRepositoriesParams{
		Page:     1,
		PageSize: 10,
	}
}

func (p *ListRepositoriesParams) ToParamString() string {
	params := []string{"?"}

	if p.Query != "" {
		params = append(params, fmt.Sprintf("q=%s", p.Query))
	}

	if p.Page != 0 {
		params = append(params, fmt.Sprintf("page=%d", p.Page))
	}

	if p.PageSize != 0 {
		params = append(params, fmt.Sprintf("page_size=%d", p.PageSize))
	}

	if p.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%s", p.Sort))
	}

	return strings.Join(params, "&")
}
