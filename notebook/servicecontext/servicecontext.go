package servicecontext

type ServiceContext struct {
	ID   uint
	Data string
	Done bool
}

func NewSCXT() *ServiceContext {
	return &ServiceContext{}
}
