package data

type ServerAPIrequest struct {
	Function string                 `json:"function"`
	Data     map[string]interface{} `json:"data"`
}
