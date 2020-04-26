package olcCpu

import (
	log "github.com/sirupsen/logrus"
)

type addressingModes struct {
	mBus *bus
	regSet *registerSet
	manEl *managingElement
	
	modeByName map[string]func() uint8
}

func (a *addressingModes) fillModeByNameHelper() {
	a.modeByName = make(map[string]func() uint8)

	a.modeByName["imp"] = a.imp
	a.modeByName["zp0"] = a.zp0
	a.modeByName["zpy"] = a.zpy
	a.modeByName["abs"] = a.abs
	a.modeByName["aby"] = a.aby
	a.modeByName["izx"] = a.izx
	a.modeByName["imm"] = a.imm
	a.modeByName["zpx"] = a.zpx
	a.modeByName["rel"] = a.rel
	a.modeByName["abx"] = a.abx
	a.modeByName["ind"] = a.ind
	a.modeByName["izy"] = a.izy
}

func (a *addressingModes) imp() uint8 {
	regSet := a.regSet
	manEl := a.manEl

	manEl.fetched = regSet.A
	return 0
}

func (a *addressingModes) zp0() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	manEl.addrAbs = uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++
	manEl.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) zpy() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	manEl.addrAbs = uint16(mBus.Read(uint16(regSet.Pc), false) + regSet.Y)
	regSet.Pc++
	manEl.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) abs() uint8 {
	// fundamentally we are interesting in that addr mode
	// because it helps to address the entire ram
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	lo := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++
	hi := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++

	manEl.addrAbs = (hi << 8) | lo
	log.Debugf(
		"abs() addrAbs: %x, hi %x. lo: %x, pc: %x",
		manEl.addrAbs,
		hi,
		lo,
		regSet.Pc,
	)
	return 0
}

func (a *addressingModes) aby() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	lo := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++
	hi := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++

	manEl.addrAbs = (hi << 8) | lo
	manEl.addrAbs += uint16(regSet.Y)

	if (manEl.addrAbs & 0xFF00) != hi << 8 {
		return 1
	}
	return 0
}

func (a *addressingModes) izx() uint8 {
	regSet := a.regSet
	manEl := a.manEl
mBus := a.mBus	

	ptrZero := uint16(mBus.Read(regSet.Pc, false))
	regSet.Pc++

	lo := uint16(mBus.Read(uint16(ptrZero + uint16(regSet.X)) & 0x00FF, false))
	hi := uint16(mBus.Read(uint16(ptrZero + uint16(regSet.X) + 1) & 0x00FF, false))

	manEl.addrAbs = (hi << 8) | lo
	return 0
}

func (a *addressingModes) imm() uint8 {
	regSet := a.regSet
	manEl := a.manEl

	manEl.addrAbs = regSet.Pc
	regSet.Pc++
	return 0
}

func (a *addressingModes) zpx() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	manEl.addrAbs = uint16(mBus.Read(regSet.Pc, false) + regSet.X)
	regSet.Pc++
	manEl.addrAbs &= 0x00FF
	return 0
}

func (a *addressingModes) rel() uint8 {
	// Only applies to branches instructions
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	manEl.addrRel = uint16(mBus.Read(regSet.Pc, false))
	regSet.Pc++

	if manEl.addrRel & 0x80 > 0 {  // is it negative jump
		manEl.addrRel |= 0xFF00  // then i fill hi bytes to negative
	}
	return 0
}

func (a *addressingModes) abx() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	lo := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++
	hi := uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++

	manEl.addrAbs = (hi << 8) | lo
	manEl.addrAbs += uint16(regSet.X)

	if (manEl.addrAbs & 0xFF00) != hi << 8 {
		return 1
	} else {
		return 0
	}
}

func (a *addressingModes) ind() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	var ptrLo uint16
	var ptrHi uint16
	ptrLo = uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++
	ptrHi = uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++

	ptr := (ptrHi << 8) | ptrLo

	if ptrLo == 0x00FF {  // Simulate page boundary hardware bug
		manEl.addrAbs = uint16((mBus.Read(ptr & 0xFF00, false) << 8) | (mBus.Read(ptr + 0, false)))
	} else {  // Behave normally
		manEl.addrAbs = uint16((mBus.Read(ptr + 1, false) << 8) | (mBus.Read(ptr + 0, false)))
	}
	return 0
}

func (a *addressingModes) izy() uint8 {
	regSet := a.regSet
	manEl := a.manEl
	mBus := a.mBus

	var ptrZero uint16
	ptrZero = uint16(mBus.Read(uint16(regSet.Pc), false))
	regSet.Pc++

	lo := uint16(mBus.Read(uint16(ptrZero) & 0x00FF, false))
	hi := uint16(mBus.Read(uint16(ptrZero + 1) & 0x00FF, false))

	manEl.addrAbs = (hi << 8) | lo
	manEl.addrAbs += uint16(regSet.Y)

	if (manEl.addrAbs & 0xFF00) != hi << 8 {
		return 1
	}
	return 0
}

func CreateAdressingModes(
		mBus *bus, regSet *registerSet, manEl *managingElement,
) *addressingModes {
	a := &addressingModes{}

	a.fillModeByNameHelper()

	a.mBus = mBus
	a.regSet = regSet
	a.manEl = manEl

	return a
}
