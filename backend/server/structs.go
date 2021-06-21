package server

//Map strukturasi
type Map struct {
	Error       bool        `json:"error"`
	Message     string      `json:"message,omitempty"`
	CurrentPage uint        `json:"currentPage,omitempty"`
	TotalPage   uint        `json:"totalPage,omitempty"`
	Total       uint        `json:"total,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
