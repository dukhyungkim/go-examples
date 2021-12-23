package model

import (
	"fmt"
	"strings"
)

type ListArtifactsParams struct {
	Query               string
	Sort                string
	Page                int64
	PageSize            int64
	WithTag             bool
	WithLabel           bool
	WithScanOverview    bool
	WithSignature       bool
	WithImmutableStatus bool
}

func NewListArtifactsParams() *ListArtifactsParams {
	return &ListArtifactsParams{
		Page:     1,
		PageSize: 10,
		WithTag:  true,
	}
}

func (p *ListArtifactsParams) ToParamString() string {
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

	if p.WithTag {
		params = append(params, fmt.Sprintf("%t", p.WithTag))
	}

	if p.WithLabel {
		params = append(params, fmt.Sprintf("%t", p.WithLabel))
	}

	if p.WithScanOverview {
		params = append(params, fmt.Sprintf("%t", p.WithScanOverview))
	}

	if p.WithSignature {
		params = append(params, fmt.Sprintf("%t", p.WithSignature))
	}

	if p.WithImmutableStatus {
		params = append(params, fmt.Sprintf("%t", p.WithImmutableStatus))
	}

	return strings.Join(params, "&")
}
