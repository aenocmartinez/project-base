package formrequest

type UpdateSequenceFormRequest struct {
	Id     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Prefix string `json:"prefix"`
	Value  int64  `json:"value"`
}
