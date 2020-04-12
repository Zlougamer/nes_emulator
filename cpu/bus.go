package cpu

type bus struct {
	cpu olc6502
	ram [64 * 1024]uint8
}

func (bus* bus) Write(addr uint16, data uint8) {
	//fmt.Printf("bus Write: %v %v\n", addr, data)
	if addr >= 0x0000 && addr <= 0xFFFF {
		bus.ram[addr] = data
	}
}

func (bus* bus) Read(addr uint16, isReadOnly bool) uint8 {
	//fmt.Printf("bus Read: %v %v\n", addr, isReadOnly)
	if addr >= 0x000 && addr <= 0xFFFF {
		return bus.ram[addr]
	}
	return 0
}

func CreateBus() *bus {
	b := &bus{}

	for i := range b.ram {
		b.ram[i] = 0x0000
	}

	b.cpu.ConnectBus(b)
	return b
}
