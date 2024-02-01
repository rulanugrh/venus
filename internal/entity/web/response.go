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