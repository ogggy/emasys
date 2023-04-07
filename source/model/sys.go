package model

type BaseResp struct {
	Ok  bool  `json:"ok"`
	Err error `json:"error"`
}
