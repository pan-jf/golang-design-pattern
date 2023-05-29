package facade

import "fmt"

// CPUFacade ...
type CPUFacade interface {
	StartUp()
	ShutDown()
	CPUOther()
}

type CPU857 struct{}

func (*CPU857) StartUp() {
	fmt.Println("cpu[857]已启动")
}
func (*CPU857) ShutDown() {
	fmt.Println("cpu[857]已关闭")
}
func (*CPU857) CPUOther() {
	fmt.Println("cpu[857]别的方法")
}

type CPU858 struct{}

func (*CPU858) StartUp() {
	fmt.Println("cpu[858]已启动")
}
func (*CPU858) ShutDown() {
	fmt.Println("cpu[858]已关闭")
}
func (*CPU858) CPUOther() {
	fmt.Println("cpu[858]别的方法")
}
