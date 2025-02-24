package dto

type Pagination struct {
	TotalData       int `json:"totalData"`
	TotalDataInPage int `json:"totalDataInPage"`
	TotalPage       int `json:"totalPage"`
	ActivePage      int `json:"activePage"`
}
