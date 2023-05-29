package facade

// Computer facade implement
type Computer struct {
	CPU    CPUFacade
	Memory MemoryFacade
	Disk   DiskFacade
}

// ComputerUnit ...
type ComputerUnit interface {
	StartUp()
	ShutDown()
}

func (c *Computer) Start() {
	c.CPU.StartUp()
	c.Memory.StartUp()
	c.Disk.StartUp()
}

func (c *Computer) ShutDown() {
	c.CPU.ShutDown()
	c.Memory.ShutDown()
	c.Disk.ShutDown()
}
