package shared_models

type ResponseModel struct {
	Status  int    `json: "status"`
	Message string `json: "message"`
	Error   any    `json: "error"`
	Data    any    `json: "data"`
}
