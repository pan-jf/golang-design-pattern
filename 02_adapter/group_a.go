package adapter

// ARequest 是适配的目标接口
type ARequest interface {
	RequestA() string
}

// aRequestImpl a实现的具体类
type aRequestImpl struct{}

// RequestA a实现的方法
func (*aRequestImpl) RequestA() string {
	return "group a request impl"
}
