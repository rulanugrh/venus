package web

type Failure struct {
	Code    int
	Message string
	Error   error
}

type Success struct {
	Code    int
	Message string
	Data    interface{}
}

type Error struct {
	Message string
	Code    int
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