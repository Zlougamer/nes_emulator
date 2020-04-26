package olcCpu

type bus struct {
	cpu *olc6502
	ram [64 * 1024]uint8
}

func (bus* bus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		bus.ram[addr] = data
	}
}

func (bus* bus) Read(addr uint16, isReadOnly bool) uint8 {
	if addr >= 0x000 && addr <= 0xFFFF {
		return bus.ram[addr]
	}
	return 0
}

func CreateBus(cpu *olc6502) *bus {
	b := &bus{}
	b.cpu = cpu

	for i := range b.ram {
		b.ram[i] = 0x0000
	}

	return b
}
