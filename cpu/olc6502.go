package cpu

import "reflect"

type flags6502 int
const (
	C flags6502 = 1 << 0
	Z flags6502 = 1 << 1
	I flags6502 = 1 << 2
	D flags6502 = 1 << 3
	B flags6502 = 1 << 4
	U flags6502 = 1 << 5
	V flags6502 = 1 << 6
	N flags6502 = 1 << 7
)

const MAX_INSTR = 16 * 16
const BASE_STKP = 0x0100

type instruction struct {
	name     string
	operate  func() uint8
	addrmode func() uint8
	cycles   uint8
}

type olc6502 struct {
	mBus *bus

	a uint8
	x uint8
	y uint8
	stkp uint8
	pc uint16
	status uint8

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode uint8
	cycles uint8

	lookup [MAX_INSTR]instruction
}

type Olc6502 interface {
	ConnectBus(b *bus)

	Read(addr uint16)
	Write(addr uint16, data uint8)

	GetFlag(flag flags6502) uint8
	SetFlag(flag flags6502, val bool)

	// Adressing modes
	IMP() uint8
	ZP0() uint8
	ZPY() uint8
	ABS() uint8
	ABY() uint8
	IZX() uint8
	IMM() uint8
	ZPX() uint8
	REL() uint8
	ABX() uint8
	IND() uint8
	IZY() uint8

	// Opcodes
	ADC() uint8
	AND() uint8
	ASL() uint8
	BCC() uint8
	BCS() uint8
	BEQ() uint8
	BIT() uint8
	BMI() uint8
	BNE() uint8
	BPL() uint8
	BRK() uint8
	BVC() uint8
	BVS() uint8
	CLC() uint8
	CLD() uint8
	CLI() uint8
	CLV() uint8
	CMP() uint8
	CPX() uint8
	CPY() uint8
	DEC() uint8
	DEX() uint8
	DEY() uint8
	EOR() uint8
	INC() uint8
	INX() uint8
	INY() uint8
	JMP() uint8
	JSR() uint8
	LDA() uint8
	LDX() uint8
	LDY() uint8
	LSR() uint8
	NOP() uint8
	ORA() uint8
	PHA() uint8
	PHP() uint8
	PLA() uint8
	PLP() uint8
	ROL() uint8
	ROR() uint8
	RTI() uint8
	RTS() uint8
	SBC() uint8
	SEC() uint8
	SED() uint8
	SEI() uint8
	STA() uint8
	STX() uint8
	STY() uint8
	TAX() uint8
	TAY() uint8
	TSX() uint8
	TXA() uint8
	TXS() uint8
	TYA() uint8

	XXX() uint8

	Clock()
	Reset()
	Irq()
	Nmi()

	fetch() uint8

}

func (cpu *olc6502) ConnectBus(b *bus) {
	cpu.mBus = b
}

func (cpu *olc6502) Read(addr uint16) uint8 {
	//fmt.Printf("cpu Read: %v\n", addr)
	return cpu.mBus.Read(addr, false)
}

func (cpu *olc6502) Write(addr uint16, data uint8) {
	cpu.mBus.Write(addr, data)
}

func (cpu *olc6502) GetFlag(flag flags6502) uint8 {
	if cpu.status & uint8(flag) > 0 {
		return 1
	} else {
		return 0
	}
}

func (cpu *olc6502) SetFlag(flag flags6502, val bool) {
	if val {
		cpu.status |= uint8(flag)
	} else {
		cpu.status &= ^uint8(flag)
	}
}

// Adressing modes

func (cpu *olc6502) IMP() uint8 {
	cpu.fetched = cpu.a
	return 0
}

func (cpu *olc6502) ZP0() uint8 {
	cpu.addrAbs = uint16(cpu.Read(uint16(cpu.pc)))
	cpu.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (cpu *olc6502) ZPY() uint8 {
	cpu.addrAbs = uint16(cpu.Read(uint16(cpu.pc)) + cpu.y)
	cpu.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (cpu *olc6502) ABS() uint8 {
	// fundamentally we are interestion in that addr mode
	// because it helps to address the entire ram
	var lo uint8
	var hi uint8
	lo = cpu.Read(uint16(cpu.pc))
	cpu.pc++
	hi = cpu.Read(uint16(cpu.pc))
	cpu.pc++

	cpu.addrAbs = uint16((hi << 8) | lo)
	return 0
}

func (cpu *olc6502) ABY() uint8 {
	var lo uint8
	var hi uint8
	lo = cpu.Read(uint16(cpu.pc))
	cpu.pc++
	hi = cpu.Read(uint16(cpu.pc))
	cpu.pc++

	cpu.addrAbs = uint16((hi << 8) | lo)
	cpu.addrAbs += uint16(cpu.y)

	if (cpu.addrAbs & 0xFF00) != uint16(hi << 8) {
		return 1
	}
	return 0
}

func (cpu *olc6502) IZX() uint8 {
	var ptrZero uint16
	ptrZero = uint16(cpu.Read(cpu.pc))
	cpu.pc++

	lo := cpu.Read(uint16(ptrZero + uint16(cpu.x)) & 0x00FF)
	hi := cpu.Read(uint16(ptrZero + uint16(cpu.x) + 1) & 0x00FF)

	cpu.addrAbs = uint16((hi << 8) | lo)
	return 0
}

func (cpu *olc6502) IMM() uint8 {
	cpu.addrAbs = cpu.pc
	cpu.pc++
	return 0
}

func (cpu *olc6502) ZPX() uint8 {
	cpu.addrAbs = uint16(cpu.Read(cpu.pc) + cpu.x)
	cpu.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (cpu *olc6502) REL() uint8 {
	// Only applies to branches instructions
	cpu.addrRel = uint16(cpu.Read(cpu.pc))
	cpu.pc++

	if cpu.addrRel & 0x80 > 0 {  // is it negative jump
		cpu.addrRel |= 0xFF00  // then i fill hi bytes to negative
	}
	return 0
}

func (cpu *olc6502) ABX() uint8 {
	var lo uint8
	var hi uint8
	lo = cpu.Read(uint16(cpu.pc))
	cpu.pc++
	hi = cpu.Read(uint16(cpu.pc))
	cpu.pc++

	cpu.addrAbs = uint16((hi << 8) | lo)
	cpu.addrAbs += uint16(cpu.x)

	if (cpu.addrAbs & 0xFF00) != uint16(hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (cpu *olc6502) IND() uint8 {
	var ptrLo uint8
	var ptrHi uint8
	ptrLo = cpu.Read(uint16(cpu.pc))
	cpu.pc++
	ptrHi = cpu.Read(uint16(cpu.pc))
	cpu.pc++

	ptr := uint16((ptrHi << 8) | ptrLo)

	if ptrLo == 0x00FF {  // Simulate page boundary hardware bug
		cpu.addrAbs = uint16((cpu.Read(ptr & 0xFF00) << 8) | (cpu.Read(ptr + 0)))
	} else {  // Behave normally
		cpu.addrAbs = uint16((cpu.Read(ptr + 1) << 8) | (cpu.Read(ptr + 0)))
	}
	return 0
}

func (cpu *olc6502) IZY() uint8 {
	var ptrZero uint16
	ptrZero = uint16(cpu.Read(uint16(cpu.pc)))
	cpu.pc++

	lo := cpu.Read(uint16(ptrZero) & 0x00FF)
	hi := cpu.Read(uint16(ptrZero + 1) & 0x00FF)

	cpu.addrAbs = uint16((hi << 8) | lo)
	cpu.addrAbs += uint16(cpu.y)

	if (cpu.addrAbs & 0xFF00) != (uint16(hi) << 8) {
		return 1
	}
	return 0
}

// Opcodes (Instructions)

func (cpu *olc6502) fetch() uint8 {
	addrmode := reflect.ValueOf(cpu.lookup[cpu.opcode].addrmode)
	impOpcode := reflect.ValueOf(cpu.IMP)

	if addrmode.Pointer() == impOpcode.Pointer() {
		cpu.fetched = cpu.Read(cpu.addrAbs)
	}
	return cpu.fetched
}

func (cpu *olc6502) ADC() uint8 {
	cpu.fetch()
	val := uint16(cpu.a) + uint16(cpu.fetched) + uint16(cpu.GetFlag(C))
	cpu.SetFlag(C, val > 255)
	cpu.SetFlag(Z, (val & 0x00FF) == 0)
	cpu.SetFlag(N, (val & 0x0080) > 0)
	vFlagVal := ^(uint16(cpu.a) ^ uint16(cpu.fetched)) & (uint16(cpu.a) ^ uint16(val))
	cpu.SetFlag(V, vFlagVal > 0)
	cpu.a = uint8(val & 0x00FF)
	return 1
}

func (cpu *olc6502) AND() uint8 {
	cpu.fetch()
	cpu.a = cpu.a & cpu.fetched
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a == 0x80)
	return 1
}

func (cpu *olc6502) ASL() uint8 {
	cpu.fetch()
	temp := uint16(cpu.fetched) << 1

	cpu.SetFlag(C, temp & 0xFF00 > 0)
	cpu.SetFlag(Z, temp & 0x00FF == 0x00)
	cpu.SetFlag(N, temp & 0x80 > 0)

	addrmode := reflect.ValueOf(cpu.lookup[cpu.opcode].addrmode)
	impOpcode := reflect.ValueOf(cpu.IMP)

	if addrmode.Pointer() == impOpcode.Pointer() {
		cpu.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}
	return 0
}

func (cpu *olc6502) BCC() uint8 {
	if cpu.GetFlag(C) == 0 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BCS() uint8 {
	if cpu.GetFlag(C) == 1 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}
		
		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BEQ() uint8 {
	if cpu.GetFlag(Z) == 1 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BIT() uint8 {
	cpu.fetch()
	temp := cpu.a & cpu.fetched
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x00)
	cpu.SetFlag(N, cpu.fetched & (1 << 7) > 0)
	cpu.SetFlag(V, cpu.fetched & (1 << 6) > 0)
	return 0
}

func (cpu *olc6502) BMI() uint8 {
	if cpu.GetFlag(N) == 1 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BNE() uint8 {
	if cpu.GetFlag(Z) == 0 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BPL() uint8 {
	if cpu.GetFlag(N) == 0 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BRK() uint8 {
	cpu.pc++

	cpu.SetFlag(I, true)
	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8((cpu.pc >> 8) & 0x00FF))
	cpu.stkp--
	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8((cpu.pc >> 8) & 0x00FF))
	cpu.stkp--

	cpu.SetFlag(B, true)
	cpu.Write(BASE_STKP + uint16(cpu.stkp), cpu.status)
	cpu.stkp--
	cpu.SetFlag(B, false)

	cpu.pc = uint16(cpu.Read(0xFFFE)) | (uint16(cpu.Read(0xFFFF)) << 8)

	return 0
}

func (cpu *olc6502) BVC() uint8 {
	if cpu.GetFlag(V) == 0 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) BVS() uint8 {
	if cpu.GetFlag(V) == 1 {
		cpu.cycles++
		cpu.addrAbs = cpu.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}
	return 0
}

func (cpu *olc6502) CLC() uint8 {
	cpu.SetFlag(C, false)
	return 0
}

func (cpu *olc6502) CLD() uint8 {
	cpu.SetFlag(D, false)
	return 0
}

func (cpu *olc6502) CLI() uint8 {
	cpu.SetFlag(I, false)
	return 0
}

func (cpu *olc6502) CLV() uint8 {
	cpu.SetFlag(V, false)
	return 0
}

func (cpu *olc6502) CMP() uint8 {
	cpu.fetch()
	temp := uint16(cpu.a) - uint16(cpu.fetched)
	cpu.SetFlag(C, cpu.a >= cpu.fetched)
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)
	return 1
}

func (cpu *olc6502) CPX() uint8 {
	cpu.fetch()
	temp := uint16(cpu.x) - uint16(cpu.fetched)
	cpu.SetFlag(C, cpu.x >= cpu.fetched)
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (cpu *olc6502) CPY() uint8 {
	cpu.fetch()
	temp := uint16(cpu.y) - uint16(cpu.fetched)
	cpu.SetFlag(C, cpu.y >= cpu.fetched)
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (cpu *olc6502) DEC() uint8 {
	cpu.fetch()
	temp := uint16(cpu.fetched) - 1
	cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (cpu *olc6502) DEX() uint8 {
	cpu.x--
	cpu.SetFlag(Z, cpu.x == 0x00)
	cpu.SetFlag(N, (cpu.x & 0x80) > 0)
	return 0
}

func (cpu *olc6502) DEY() uint8 {
	cpu.y--
	cpu.SetFlag(Z, cpu.y == 0x00)
	cpu.SetFlag(N, (cpu.y & 0x80) > 0)
	return 0
}

func (cpu *olc6502) EOR() uint8 {
	cpu.fetch()
	cpu.a = cpu.a ^ cpu.fetched
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a == 0x80)
	return 1
}

func (cpu *olc6502) INC() uint8 {
	cpu.fetch()
	temp := uint16(cpu.fetched) + 1
	cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (cpu *olc6502) INX() uint8 {
	cpu.x++
	cpu.SetFlag(Z, cpu.x == 0x00)
	cpu.SetFlag(N, (cpu.x & 0x80) > 0)
	return 0
}

func (cpu *olc6502) INY() uint8 {
	cpu.y++
	cpu.SetFlag(Z, cpu.y == 0x00)
	cpu.SetFlag(N, (cpu.y & 0x80) > 0)
	return 0
}

func (cpu *olc6502) JMP() uint8 {
	cpu.pc = cpu.addrAbs
	return 0
}

func (cpu *olc6502) JSR() uint8 {
	cpu.pc--

	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8((cpu.pc >> 8) & 0x00FF))
	cpu.stkp--
	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8(cpu.pc & 0x00FF))
	cpu.stkp--

	cpu.pc = cpu.addrAbs
	return 0
}

func (cpu *olc6502) LDA() uint8 {
	cpu.fetch()
	cpu.a = cpu.fetched
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, (cpu.a & 0x80) > 0)
	return 1
}

func (cpu *olc6502) LDX() uint8 {
	cpu.fetch()
	cpu.x = cpu.fetched
	cpu.SetFlag(Z, cpu.x == 0x00)
	cpu.SetFlag(N, (cpu.x & 0x80) > 0)
	return 1
}

func (cpu *olc6502) LDY() uint8 {
	cpu.fetch()
	cpu.y = cpu.fetched
	cpu.SetFlag(Z, cpu.y == 0x00)
	cpu.SetFlag(N, (cpu.y & 0x80) > 0)
	return 1
}

func (cpu *olc6502) LSR() uint8 {
	cpu.fetch()
	cpu.SetFlag(C, cpu.fetched & 0x0001 > 0)
	temp := cpu.fetched >> 1
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)

	addrmode := reflect.ValueOf(cpu.lookup[cpu.opcode].addrmode)
	impOpcode := reflect.ValueOf(cpu.IMP)

	if addrmode.Pointer() == impOpcode.Pointer() {
		cpu.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (cpu *olc6502) NOP() uint8 {
	switch cpu.opcode {
	case 0x1C:
		return 1
	case 0x3C:
		return 1
	case 0x5C:
		return 1
	case 0x7C:
		return 1
	case 0xDC:
		return 1
	case 0xFC:
		return 1
	default:
		return 0
	}
}

func (cpu *olc6502) ORA() uint8 {
	cpu.fetch()
	cpu.a = cpu.a | cpu.fetched
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a == 0x80)
	return 1
}

func (cpu *olc6502) PHA() uint8 {
	cpu.Write(BASE_STKP + uint16(cpu.stkp), cpu.a)
	cpu.stkp--
	return 0
}

func (cpu *olc6502) PHP() uint8 {
	cpu.Write(BASE_STKP + uint16(cpu.stkp), cpu.status | uint8(B) | uint8(U))
	cpu.SetFlag(B, false)
	cpu.SetFlag(U, false)
	cpu.stkp--
	return 0
}

func (cpu *olc6502) PLA() uint8 {
	cpu.stkp++
	cpu.a = cpu.Read(BASE_STKP + uint16(cpu.stkp))
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a & 0x80 > 0)
	return 0
}

func (cpu *olc6502) PLP() uint8 {
	cpu.stkp++
	cpu.status = cpu.Read(BASE_STKP + uint16(cpu.stkp))
	cpu.SetFlag(U, true)
	return 0
}

func (cpu *olc6502) ROL() uint8 {
	cpu.fetch()
	temp := (uint16(cpu.fetched) << 1) | uint16(cpu.GetFlag(C))


	cpu.SetFlag(C, temp & 0xFF00 > 0)
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)

	addrmode := reflect.ValueOf(cpu.lookup[cpu.opcode].addrmode)
	impOpcode := reflect.ValueOf(cpu.IMP)

	if addrmode.Pointer() == impOpcode.Pointer() {
		cpu.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (cpu *olc6502) ROR() uint8 {
	cpu.fetch()
	temp := (uint16(cpu.GetFlag(C)) << 7) | (uint16(cpu.fetched) >> 1)

	cpu.SetFlag(C, cpu.fetched & 0x01 > 0)
	cpu.SetFlag(Z, (temp & 0x00FF) == 0x0000)
	cpu.SetFlag(N, (temp & 0x0080) > 0)

	addrmode := reflect.ValueOf(cpu.lookup[cpu.opcode].addrmode)
	impOpcode := reflect.ValueOf(cpu.IMP)

	if addrmode.Pointer() == impOpcode.Pointer() {
		cpu.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (cpu *olc6502) RTI() uint8 {
	cpu.stkp++
	cpu.status = cpu.Read(BASE_STKP + uint16(cpu.stkp))
	cpu.status &= ^uint8(B)
	cpu.status &= ^uint8(U)

	cpu.stkp++
	cpu.pc = uint16(cpu.Read(BASE_STKP + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.Read(BASE_STKP + uint16(cpu.stkp))) << 8
	return 0
}

func (cpu *olc6502) RTS() uint8 {
	cpu.stkp++
	cpu.pc = uint16(cpu.Read(BASE_STKP + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.Read(BASE_STKP + uint16(cpu.stkp))) << 8

	cpu.pc++

	return 0
}

func (cpu *olc6502) SBC() uint8 {
	cpu.fetch()
	invFetch := uint16(cpu.fetched & 0x00FF)
	val := uint16(cpu.a) + invFetch + uint16(cpu.GetFlag(C))
	cpu.SetFlag(C, val & 0xFF00 > 0)
	cpu.SetFlag(Z, (val & 0x00FF) == 0)
	cpu.SetFlag(N, (val & 0x0080) > 0)
	vFlagVal := (val ^ uint16(cpu.a)) & (val ^ invFetch) & 0x0080
	cpu.SetFlag(V, vFlagVal > 0)
	cpu.a = uint8(val & 0x00FF)
	return 1
}

func (cpu *olc6502) SEC() uint8 {
	cpu.SetFlag(C, true)
	return 0
}

func (cpu *olc6502) SED() uint8 {
	cpu.SetFlag(D, true)
	return 0
}

func (cpu *olc6502) SEI() uint8 {
	cpu.SetFlag(I, true)
	return 0
}

func (cpu *olc6502) STA() uint8 {
	cpu.Write(cpu.addrAbs, cpu.a)
	return 0
}

func (cpu *olc6502) STX() uint8 {
	cpu.Write(cpu.addrAbs, cpu.x)
	return 0
}

func (cpu *olc6502) STY() uint8 {
	cpu.Write(cpu.addrAbs, cpu.y)
	return 0
}

func (cpu *olc6502) TAX() uint8 {
	cpu.x = cpu.a
	cpu.SetFlag(Z, cpu.x == 0x00)
	cpu.SetFlag(N, cpu.x & 0x80 > 0)
	return 0
}

func (cpu *olc6502) TAY() uint8 {
	cpu.y = cpu.a
	cpu.SetFlag(Z, cpu.y == 0x00)
	cpu.SetFlag(N, cpu.y & 0x80 > 0)
	return 0
}

func (cpu *olc6502) TSX() uint8 {
	cpu.x = cpu.stkp
	cpu.SetFlag(Z, cpu.x == 0x00)
	cpu.SetFlag(N, cpu.x & 0x80 > 0)
	return 0
}

func (cpu *olc6502) TXA() uint8 {
	cpu.a = cpu.x
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a & 0x80 > 0)
	return 0
}

func (cpu *olc6502) TXS() uint8 {
	cpu.stkp = cpu.x
	return 0
}

func (cpu *olc6502) TYA() uint8 {
	cpu.a = cpu.y
	cpu.SetFlag(Z, cpu.a == 0x00)
	cpu.SetFlag(N, cpu.a & 0x80 > 0)
	return 0
}

func (cpu *olc6502) XXX() uint8 {
	return 0
}


// End of Opcodes

func (cpu *olc6502) Clock() {
	if cpu.cycles == 0 {
		opcode := cpu.Read(uint16(cpu.pc))  // Recheck it
		cpu.pc++

		cpu.cycles = cpu.lookup[opcode].cycles

		additionalCycleAddr := cpu.lookup[opcode].addrmode()
		additionalCycleOp := cpu.lookup[opcode].operate()

		cpu.cycles += additionalCycleAddr & additionalCycleOp
	}

	cpu.cycles--
}

func (cpu *olc6502) Reset() {
	cpu.a = 0x00
	cpu.x = 0x00
	cpu.y = 0x00
	cpu.stkp = 0xFD
	cpu.status = 0x00 | uint8(U)

	cpu.addrAbs = 0xFFFC
	hi := uint16(cpu.Read(cpu.addrAbs + 0))
	lo := uint16(cpu.Read(cpu.addrAbs + 1))

	cpu.pc = (hi << 8) | lo

	cpu.addrAbs = 0x0000
	cpu.addrRel = 0x0000
	cpu.fetched = 0x00

	cpu.cycles = 8
}


func (cpu *olc6502) Irq() {
	if cpu.GetFlag(I) > 0 {
		cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8((cpu.pc >> 8) & 0x00FF))
		cpu.stkp--
		cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8(cpu.pc & 0x00FF))
		cpu.stkp--

		cpu.SetFlag(B, false)
		cpu.SetFlag(U, true)
		cpu.SetFlag(I, true)
		cpu.Write(BASE_STKP + uint16(cpu.stkp), cpu.status)
		cpu.stkp--

		cpu.addrAbs = 0xFFFE
		hi := uint16(cpu.Read(cpu.addrAbs + 0))
		lo := uint16(cpu.Read(cpu.addrAbs + 1))
		cpu.pc = (hi << 8) | lo

		cpu.cycles = 7
	}
}


func (cpu *olc6502) Nmi() {
	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8((cpu.pc >> 8) & 0x00FF))
	cpu.stkp--
	cpu.Write(BASE_STKP + uint16(cpu.stkp), uint8(cpu.pc & 0x00FF))
	cpu.stkp--

	cpu.SetFlag(B, false)
	cpu.SetFlag(U, true)
	cpu.SetFlag(I, true)
	cpu.Write(BASE_STKP + uint16(cpu.stkp), cpu.status)
	cpu.stkp--

	cpu.addrAbs = 0xFFFA
	hi := uint16(cpu.Read(cpu.addrAbs + 0))
	lo := uint16(cpu.Read(cpu.addrAbs + 1))
	cpu.pc = (hi << 8) | lo

	cpu.cycles = 8
}


func CreateOlc6502() *olc6502 {
	olc := &olc6502{}

	olc.a = 0x00
	olc.x = 0x00
	olc.y = 0x00
	olc.stkp = 0x00
	olc.pc = 0x00
	olc.status = 0x00

	olc.fetched = 0x00

	olc.addrAbs = 0x0000
	olc.addrRel = 0x0000

	olc.opcode = 0x00
	olc.cycles = 0x00

	olc.mBus = CreateBus()

	olc.lookup = [MAX_INSTR]instruction{
		{"BRK", olc.BRK, olc.IMM, 7}, {"ORA", olc.ORA, olc.IZX, 6}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 3}, {"ORA", olc.ORA, olc.ZP0, 3}, {"ASL", olc.ASL, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"PHP", olc.PHP, olc.IMP, 3}, {"BPL", olc.BPL, olc.REL, 2}, {"ORA", olc.ORA, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"ORA", olc.ORA, olc.ZPX, 4}, {"ASL", olc.ASL, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"CLC", olc.CLC, olc.IMP, 2},
		{"JSR", olc.JSR, olc.ABS, 6}, {"AND", olc.AND, olc.IZX, 6}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"BIT", olc.BIT, olc.ZP0, 3}, {"AND", olc.AND, olc.ZP0, 3}, {"ROL", olc.ROL, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"PLP", olc.PLP, olc.IMP, 4}, {"BMI", olc.BMI, olc.REL, 2}, {"AND", olc.AND, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"AND", olc.AND, olc.ZPX, 4}, {"ROL", olc.ROL, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"SEC", olc.SEC, olc.IMP, 2},
		{"RTI", olc.RTI, olc.IMP, 6}, {"EOR", olc.EOR, olc.IZX, 6}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 3}, {"EOR", olc.EOR, olc.ZP0, 3}, {"LSR", olc.LSR, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"PHA", olc.PHA, olc.IMP, 3}, {"BVC", olc.BVC, olc.REL, 2}, {"EOR", olc.EOR, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"EOR", olc.EOR, olc.ZPX, 4}, {"LSR", olc.LSR, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"CLI", olc.CLI, olc.IMP, 2},
		{"RTS", olc.RTS, olc.IMP, 6}, {"ADC", olc.ADC, olc.IZX, 6}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 3}, {"ADC", olc.ADC, olc.ZP0, 3}, {"ROR", olc.ROR, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"PLA", olc.PLA, olc.IMP, 4}, {"BVS", olc.BVS, olc.REL, 2}, {"ADC", olc.ADC, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"ADC", olc.ADC, olc.ZPX, 4}, {"ROR", olc.ROR, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"SEI", olc.SEI, olc.IMP, 2},
		{"???", olc.NOP, olc.IMP, 2}, {"STA", olc.STA, olc.IZX, 6}, {"???", olc.NOP, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 6}, {"STY", olc.STY, olc.ZP0, 3}, {"STA", olc.STA, olc.ZP0, 3}, {"STX", olc.STX, olc.ZP0, 3}, {"???", olc.XXX, olc.IMP, 3}, {"DEY", olc.DEY, olc.IMP, 2}, {"BCC", olc.BCC, olc.REL, 2}, {"STA", olc.STA, olc.IZY, 6}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 6}, {"STY", olc.STY, olc.ZPX, 4}, {"STA", olc.STA, olc.ZPX, 4}, {"STX", olc.STX, olc.ZPY, 4}, {"???", olc.XXX, olc.IMP, 4}, {"TYA", olc.TYA, olc.IMP, 2},
		{"LDY", olc.LDY, olc.IMM, 2}, {"LDA", olc.LDA, olc.IZX, 6}, {"LDX", olc.LDX, olc.IMM, 2}, {"???", olc.XXX, olc.IMP, 6}, {"LDY", olc.LDY, olc.ZP0, 3}, {"LDA", olc.LDA, olc.ZP0, 3}, {"LDX", olc.LDX, olc.ZP0, 3}, {"???", olc.XXX, olc.IMP, 3}, {"TAY", olc.TAY, olc.IMP, 2}, {"BCS", olc.BCS, olc.REL, 2}, {"LDA", olc.LDA, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 5}, {"LDY", olc.LDY, olc.ZPX, 4}, {"LDA", olc.LDA, olc.ZPX, 4}, {"LDX", olc.LDX, olc.ZPY, 4}, {"???", olc.XXX, olc.IMP, 4}, {"CLV", olc.CLV, olc.IMP, 2},
		{"CPY", olc.CPY, olc.IMM, 2}, {"CMP", olc.CMP, olc.IZX, 6}, {"???", olc.NOP, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"CPY", olc.CPY, olc.ZP0, 3}, {"CMP", olc.CMP, olc.ZP0, 3}, {"DEC", olc.DEC, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"INY", olc.INY, olc.IMP, 2}, {"BNE", olc.BNE, olc.REL, 2}, {"CMP", olc.CMP, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"CMP", olc.CMP, olc.ZPX, 4}, {"DEC", olc.DEC, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"CLD", olc.CLD, olc.IMP, 2},
		{"CPX", olc.CPX, olc.IMM, 2}, {"SBC", olc.SBC, olc.IZX, 6}, {"???", olc.NOP, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"CPX", olc.CPX, olc.ZP0, 3}, {"SBC", olc.SBC, olc.ZP0, 3}, {"INC", olc.INC, olc.ZP0, 5}, {"???", olc.XXX, olc.IMP, 5}, {"INX", olc.INX, olc.IMP, 2}, {"BEQ", olc.BEQ, olc.REL, 2}, {"SBC", olc.SBC, olc.IZY, 5}, {"???", olc.XXX, olc.IMP, 2}, {"???", olc.XXX, olc.IMP, 8}, {"???", olc.NOP, olc.IMP, 4}, {"SBC", olc.SBC, olc.ZPX, 4}, {"INC", olc.INC, olc.ZPX, 6}, {"???", olc.XXX, olc.IMP, 6}, {"SED", olc.SED, olc.IMP, 2},
	}

	return olc
}
