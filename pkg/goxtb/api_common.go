package goxtb

import "encoding/json"

type ApiRequest struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"arguments"`
}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	AppId    string `json:"appId"`
	AppName  string `json:"appName"`
}

type LoginResponse struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
}
