package olcCpu

import (
	log "github.com/sirupsen/logrus"
)

type addressingModes struct {
	cpu *olc6502
	modeByName map[string]func() uint8
}

type AddressingModes interface {
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
}

func (a *addressingModes) fillModeByName() {
	a.modeByName = make(map[string]func() uint8)

	a.modeByName["IMP"] = a.IMP
	a.modeByName["ZP0"] = a.ZP0
	a.modeByName["ZPY"] = a.ZPY
	a.modeByName["ABS"] = a.ABS
	a.modeByName["ABY"] = a.ABY
	a.modeByName["IZX"] = a.IZX
	a.modeByName["IMM"] = a.IMM
	a.modeByName["ZPX"] = a.ZPX
	a.modeByName["REL"] = a.REL
	a.modeByName["ABX"] = a.ABX
	a.modeByName["IND"] = a.IND
	a.modeByName["IZY"] = a.IZY
}

func (a *addressingModes) IMP() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu
	
	cpu.fetched = regSet.a
	return 0
}

func (a *addressingModes) ZP0() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	cpu.addrAbs = uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) ZPY() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	cpu.addrAbs = uint16(cpu.Read(uint16(regSet.pc)) + regSet.y)
	regSet.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) ABS() uint8 {
	// fundamentally we are interestion in that addr mode
	// because it helps to address the entire ram
	regSet := a.cpu.regSet
	cpu := a.cpu

	lo := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++
	hi := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++

	cpu.addrAbs = (hi << 8) | lo
	log.Debugf(
		"ABS() addrAbs: %x, hi %x. lo: %x, pc: %x",
		cpu.addrAbs,
		hi,
		lo,
		regSet.pc,
	)
	return 0
}

func (a *addressingModes) ABY() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	lo := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++
	hi := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(regSet.y)

	if (cpu.addrAbs & 0xFF00) != hi << 8 {
		return 1
	}
	return 0
}

func (a *addressingModes) IZX() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	ptrZero := uint16(cpu.Read(regSet.pc))
	regSet.pc++

	lo := uint16(cpu.Read(uint16(ptrZero + uint16(regSet.x)) & 0x00FF))
	hi := uint16(cpu.Read(uint16(ptrZero + uint16(regSet.x) + 1) & 0x00FF))

	cpu.addrAbs = (hi << 8) | lo
	return 0
}

func (a *addressingModes) IMM() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	cpu.addrAbs = regSet.pc
	regSet.pc++
	return 0
}

func (a *addressingModes) ZPX() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	cpu.addrAbs = uint16(cpu.Read(regSet.pc) + regSet.x)
	regSet.pc++
	cpu.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) REL() uint8 {
	// Only applies to branches instructions
	regSet := a.cpu.regSet
	cpu := a.cpu

	cpu.addrRel = uint16(cpu.Read(regSet.pc))
	regSet.pc++

	if cpu.addrRel & 0x80 > 0 {  // is it negative jump
		cpu.addrRel |= 0xFF00  // then i fill hi bytes to negative
	}
	return 0
}

func (a *addressingModes) ABX() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	lo := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++
	hi := uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(regSet.x)

	if (cpu.addrAbs & 0xFF00) != hi << 8 {
		return 1
	} else {
		return 0
	}
}

func (a *addressingModes) IND() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	var ptrLo uint16
	var ptrHi uint16
	ptrLo = uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++
	ptrHi = uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++

	ptr := (ptrHi << 8) | ptrLo

	if ptrLo == 0x00FF {  // Simulate page boundary hardware bug
		cpu.addrAbs = uint16((cpu.Read(ptr & 0xFF00) << 8) | (cpu.Read(ptr + 0)))
	} else {  // Behave normally
		cpu.addrAbs = uint16((cpu.Read(ptr + 1) << 8) | (cpu.Read(ptr + 0)))
	}
	return 0
}

func (a *addressingModes) IZY() uint8 {
	regSet := a.cpu.regSet
	cpu := a.cpu

	var ptrZero uint16
	ptrZero = uint16(cpu.Read(uint16(regSet.pc)))
	regSet.pc++

	lo := uint16(cpu.Read(uint16(ptrZero) & 0x00FF))
	hi := uint16(cpu.Read(uint16(ptrZero + 1) & 0x00FF))

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(regSet.y)

	if (cpu.addrAbs & 0xFF00) != hi << 8 {
		return 1
	}
	return 0
}

func CreateAdressingModes(cpu *olc6502) *addressingModes {
	a := &addressingModes{}

	a.fillModeByName()
	a.cpu = cpu

	a.cpu.connectAdressingModes(a)

	return a
}
