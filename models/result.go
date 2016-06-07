package models

//"errors"
//"strconv"
//"time"

const (
	ResultFalse = iota
	ResultTrue
	ResultSessionErr
)

type Result struct {
	Sucess  string
	Info    interface{}
	Message string
}
