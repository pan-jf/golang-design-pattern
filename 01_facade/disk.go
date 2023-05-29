package facade

import "fmt"

// DiskFacade ...
type DiskFacade interface {
	StartUp()
	ShutDown()
	DiskOther()
}

type DiskSSD struct{}

func (*DiskSSD) StartUp() {
	fmt.Println("固态硬盘已启动")
}
func (*DiskSSD) ShutDown() {
	fmt.Println("固态硬盘已关闭")
}
func (*DiskSSD) DiskOther() {
	fmt.Println("固态硬盘其他方法")
}

type DiskHDD struct{}

func (*DiskHDD) StartUp() {
	fmt.Println("机械硬盘已启动")
}
func (*DiskHDD) ShutDown() {
	fmt.Println("机械硬盘已关闭")
}
func (*DiskHDD) DiskOther() {
	fmt.Println("机械硬盘其他方法")
}
