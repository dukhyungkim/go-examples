package model

func (p *Pagination) GetOffset() int {
	offset := 0
	if p == nil {
		return offset
	}

	if p.Offset != nil {
		offset = *p.Offset
	}
	return offset
}

func (p *Pagination) GetLimit() int {
	limit := 20

	if p == nil {
		return limit
	}

	if p.Limit != nil {
		limit = *p.Limit
	}
	return limit
}
