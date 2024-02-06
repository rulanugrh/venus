package web

type Failure struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Error   interface{} `json:"error"`
}

type Success struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Message string `json:"msg"`
	Code    int `json:"code"`
}

func (err Error) Error() string {
	return err.Message
}

type ValidationList struct {
	Field string      `json:"field"`
	Error interface{} `json:"error"`
}

type ValidationError struct {
	Message string           `json:"message"`
	Errors  []ValidationList `json:"error"`
}

func (err ValidationError) Error() string {
	return err.Message
}

type WebValidationError struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
}