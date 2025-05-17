package policy

type Policy interface {
	Name() string
	Approve(Request) (ok bool, err error)
}
