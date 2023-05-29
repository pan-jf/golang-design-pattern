package adapter

// groupA实现的方法名叫RequestA
// groupB实现的方法名叫SpecificRequest
// 通过适配器进行转换

type requestAdapter struct {
}

func (*requestAdapter) Request(requestType string) string {
	switch requestType {
	case "a":
		impl := aRequestImpl{}
		return impl.RequestA()
	case "b":
		impl := bRequestImpl{}
		return impl.SpecificRequest()
	}
	return "未找到对应实例"
}
