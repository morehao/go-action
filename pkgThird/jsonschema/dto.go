package jsonschema

type DefaultRender struct {
	Code    int         `json:"code" doc:"错误码"`
	Message string      `json:"message" doc:"错误信息"`
	Data    interface{} `json:"data" doc:"数据"`
}
