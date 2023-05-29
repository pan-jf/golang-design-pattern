package facade

import "testing"

// TestFacadeAPI1 ...
func TestFacadeAPI1(t *testing.T) {
	computer := &Computer{
		CPU:    &CPU857{},
		Memory: &Memory8{},
		Disk:   &DiskHDD{},
	}
	// 启动3个组件不需要每个都调用一次，只需要启动computer即可
	computer.Start()

}

// TestFacadeAPI2 ...
func TestFacadeAPI2(t *testing.T) {
	computer := &Computer{
		CPU:    &CPU858{},
		Memory: &Memory16{},
		Disk:   &DiskSSD{},
	}

	computer.Start()
}
