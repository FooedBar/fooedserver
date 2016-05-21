package models

type GenericResponseWithCount struct {
	Count int64 `json:"count"`
}

type GenericLimitOffsetRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (req *GenericLimitOffsetRequest) Parameters() {
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
}
