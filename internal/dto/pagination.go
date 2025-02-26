package dto

import "math"

type Pagination struct {
	TotalData       int `json:"totalData"`
	TotalDataInPage int `json:"totalDataInPage"`
	TotalPage       int `json:"totalPage"`
	ActivePage      int `json:"activePage"`
}

func NewPagination(totalData int, activePage int) Pagination {
	return Pagination{
		TotalData:       totalData,
		TotalDataInPage: 10,
		ActivePage:      activePage,
		TotalPage:       int(math.Max(float64(totalData/10), 1)),
	}
}
