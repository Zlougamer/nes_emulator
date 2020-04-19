package olcCpu

const (
	C  = 1 << iota
	Z
	I
	D
	B
	U
	V
	N
)

type registerSet struct {
	cpu *olc6502

	a uint8
	x uint8
	y uint8
	stkp uint8
	pc uint16
	status uint8
}

func (r *registerSet) getFlag(flag uint8) uint8 {
	if r.status & flag > 0 {
		return 1
	} else {
		return 0
	}
}

func (r *registerSet) setFlag(flag uint8, val bool) {
	if val {
		r.status |= flag
	} else {
		r.status &= ^flag
	}
}

func CreateRegisterSet(cpu *olc6502) *registerSet {
	r := &registerSet{}
	r.cpu = cpu

	r.a = 0x00
	r.x = 0x00
	r.y = 0x00
	r.stkp = 0x00
	r.pc = 0x00
	r.status = 0x00

	r.cpu.connectRegisterSet(r)
	return r
}
