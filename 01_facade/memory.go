package facade

import "fmt"

// MemoryFacade ...
type MemoryFacade interface {
	StartUp()
	ShutDown()
	MemoryOther()
}

type Memory8 struct{}

func (*Memory8) StartUp() {
	fmt.Println("内存[8G版]已启动")
}
func (*Memory8) ShutDown() {
	fmt.Println("内存[8G版]已关闭")
}
func (*Memory8) MemoryOther() {
	fmt.Println("内存[8G版]其他方法")
}

type Memory16 struct{}

func (*Memory16) StartUp() {
	fmt.Println("内存[16G版]已启动")
}
func (*Memory16) ShutDown() {
	fmt.Println("内存[16G版]已关闭")
}
func (*Memory16) MemoryOther() {
	fmt.Println("内存[16G版]其他方法")
}
