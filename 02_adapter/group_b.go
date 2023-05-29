package adapter

// BRequest 是被适配的目标接口
type BRequest interface {
	SpecificRequest() string
}

// bRequestImpl b实现的具体类
type bRequestImpl struct{}

// SpecificRequest 是目标类的一个方法
func (*bRequestImpl) SpecificRequest() string {
	return "group b request impl"
}
