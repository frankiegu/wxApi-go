package wxApi

import (
	"encoding/json"
	"fmt"
)

type wxErr struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type mchErr struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

func (m mchErr) IsRequestSuccess() bool {
	return m.ReturnCode == "SUCCESS" && m.ResultCode == "SUCCESS" && (m.ErrCode == "SUCCESS" || m.ErrCode == "")
}

func (m mchErr) IsResponseUnCertain() bool {
	return m.ErrCode == "SYSTEMERROR"
}

func (m mchErr) Error() string {
	if m.ErrCodeDes != "" {
		return m.ErrCodeDes
	} else if m.ReturnMsg != "" {
		return m.ReturnMsg
	} else {
		return ""
	}
}

func parseJsonErr(raw []byte) (e wxErr, err error) {
	err = json.Unmarshal(raw, &e)
	if err != nil {
		return
	}
	if e.ErrCode > 0 {
		err = fmt.Errorf("微信提示: %v", e.ErrMsg)
	}
	return
}
