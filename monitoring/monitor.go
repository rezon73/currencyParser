package monitoring

type Monitor interface {
	Check() []error
}