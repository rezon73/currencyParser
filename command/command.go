package command


type Command interface {
	Exec() error
}