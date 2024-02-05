package http

import "fmt"

type ErrorResp struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

func NewError(err error, code int) ErrorResp {
	return ErrorResp{
		Code:    code,
		Error:   err.Error(),
		Success: false,
	}
}

func NewSuccess(code int) SuccessResp {
	return SuccessResp{Code: code}
}

type SuccessResp struct {
	Code int `json:"code"`
}

// SuccessResp can never marshal to success:false
func (s *SuccessResp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"code":%d,"success":true}`, s.Code)), nil
}
