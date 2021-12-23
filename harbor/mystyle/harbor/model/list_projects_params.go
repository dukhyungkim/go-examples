package model

import (
	"fmt"
	"strings"
)

type ListProjectParams struct {
	Query      string
	Page       int64
	PageSize   int64
	Sort       string
	Name       string
	Public     bool
	Owner      string
	WithDetail bool
}

func NewListProjectsParams() *ListProjectParams {
	return &ListProjectParams{
		Page:       1,
		PageSize:   10,
		WithDetail: true,
	}
}

func (p *ListProjectParams) ToParamString() string {
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

	if p.Name != "" {
		params = append(params, fmt.Sprintf("name=%s", p.Name))
	}

	if !p.Public {
		params = append(params, fmt.Sprintf("public=%t", p.Public))
	}

	if p.Owner != "" {
		params = append(params, fmt.Sprintf("owner=%s", p.Owner))
	}

	if !p.WithDetail {
		params = append(params, fmt.Sprintf("with_detail=%t", p.WithDetail))
	}

	return strings.Join(params, "&")
}
