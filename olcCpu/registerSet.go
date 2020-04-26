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
	A uint8
	X uint8
	Y uint8
	Stkp uint8
	Pc uint16
	Status uint8
}

func (r *registerSet) getFlag(flag uint8) uint8 {
	if r.Status & flag > 0 {
		return 1
	} else {
		return 0
	}
}

func (r *registerSet) setFlag(flag uint8, val bool) {
	if val {
		r.Status |= flag
	} else {
		r.Status &= ^flag
	}
}

func CreateRegisterSet() *registerSet {
	r := &registerSet{}

	r.A = 0x00
	r.X = 0x00
	r.Y = 0x00
	r.Stkp = 0xFF
	r.Pc = 0x00
	r.Status = 0x00

	return r
}
