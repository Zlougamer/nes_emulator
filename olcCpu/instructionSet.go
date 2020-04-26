package olcCpu

type instructionSet struct {
	mBus *bus
	regSet *registerSet
	manEl *managingElement

	instrByName map[string]func() uint8
}

func (i *instructionSet) fillInstrByName() {
	i.instrByName = make(map[string]func() uint8)

	i.instrByName["adc"] = i.adc
	i.instrByName["and"] = i.and
	i.instrByName["asl"] = i.asl
	i.instrByName["bcc"] = i.bcc
	i.instrByName["bcs"] = i.bcs
	i.instrByName["beq"] = i.beq
	i.instrByName["bit"] = i.bit
	i.instrByName["bmi"] = i.bmi
	i.instrByName["bne"] = i.bne
	i.instrByName["bpl"] = i.bpl
	i.instrByName["brk"] = i.brk
	i.instrByName["bvc"] = i.bvc
	i.instrByName["bvs"] = i.bvs
	i.instrByName["clc"] = i.clc
	i.instrByName["cld"] = i.cld
	i.instrByName["cli"] = i.cli
	i.instrByName["clv"] = i.clv
	i.instrByName["cmp"] = i.cmp
	i.instrByName["cpx"] = i.cpx
	i.instrByName["cpy"] = i.cpy
	i.instrByName["dec"] = i.dec
	i.instrByName["dex"] = i.dex
	i.instrByName["dey"] = i.dey
	i.instrByName["eor"] = i.eor
	i.instrByName["inc"] = i.inc
	i.instrByName["inx"] = i.inx
	i.instrByName["iny"] = i.iny
	i.instrByName["jmp"] = i.jmp
	i.instrByName["jsr"] = i.jsr
	i.instrByName["lda"] = i.lda
	i.instrByName["ldx"] = i.ldx
	i.instrByName["ldy"] = i.ldy
	i.instrByName["lsr"] = i.lsr
	i.instrByName["nop"] = i.nop
	i.instrByName["ora"] = i.ora
	i.instrByName["pha"] = i.pha
	i.instrByName["php"] = i.php
	i.instrByName["pla"] = i.pla
	i.instrByName["plp"] = i.plp
	i.instrByName["rol"] = i.rol
	i.instrByName["ror"] = i.ror
	i.instrByName["rti"] = i.rti
	i.instrByName["rts"] = i.rts
	i.instrByName["sbc"] = i.sbc
	i.instrByName["sec"] = i.sec
	i.instrByName["sed"] = i.sed
	i.instrByName["sei"] = i.sei
	i.instrByName["sta"] = i.sta
	i.instrByName["stx"] = i.stx
	i.instrByName["sty"] = i.sty
	i.instrByName["tax"] = i.tax
	i.instrByName["tay"] = i.tay
	i.instrByName["tsx"] = i.tsx
	i.instrByName["txa"] = i.txa
	i.instrByName["txs"] = i.txs
	i.instrByName["tya"] = i.tya
	i.instrByName["xxx"] = i.xxx

}

func (i *instructionSet) adc() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	
	manEl.fetch()
	temp := uint16(regSet.A) + uint16(manEl.fetched) + uint16(regSet.getFlag(C))
	regSet.setFlag(C, temp > 255)
	regSet.setFlag(Z, (temp & 0x00FF) == 0)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	vFlagVal := (
		^(uint16(regSet.A) ^ uint16(manEl.fetched)) &
		(uint16(regSet.A) ^ uint16(temp))) & 0x0080
	regSet.setFlag(V, vFlagVal > 0)

	regSet.A = uint8(temp & 0x00FF)
	return 1
}

func (i *instructionSet) and() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.A = regSet.A & manEl.fetched
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A & 0x80 != 0)
	return 1
}

func (i *instructionSet) asl() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	temp := uint16(manEl.fetched) << 1

	regSet.setFlag(C, temp & 0xFF00 > 0)
	regSet.setFlag(Z, temp & 0x00FF == 0x00)
	regSet.setFlag(N, temp & 0x80 > 0)

	addrmodeName := manEl.lookup[manEl.opcode].addrmodeName

	if addrmodeName == "imp" {
		regSet.A = uint8(temp & 0x00FF)
	} else {
		mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	}
	return 0
}

func (i *instructionSet) bcc() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(C) == 0 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) bcs() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(C) == 1 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) beq() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(Z) == 1 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) bit() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	temp := regSet.A & manEl.fetched
	regSet.setFlag(Z, (temp & 0x00FF) == 0x00)
	regSet.setFlag(N, manEl.fetched & (1 << 7) > 0)
	regSet.setFlag(V, manEl.fetched & (1 << 6) > 0)
	return 0
}

func (i *instructionSet) bmi() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(N) == 1 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) bne() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(Z) == 0 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) bpl() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(N) == 0 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) brk() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	//regSet.Pc++

	regSet.setFlag(I, true)
	mBus.Write(BASE_STKP + uint16(regSet.Stkp), uint8((regSet.Pc >> 8) & 0x00FF))
	regSet.Stkp--
	mBus.Write(BASE_STKP + uint16(regSet.Stkp), uint8((regSet.Pc) & 0x00FF))
	regSet.Stkp--

	regSet.setFlag(B, true)
	mBus.Write(BASE_STKP + uint16(regSet.Stkp), regSet.Status)
	regSet.Stkp--
	regSet.setFlag(B, false)

	regSet.Pc = uint16(mBus.Read(0xFFFE, false)) | (uint16(mBus.Read(0xFFFF, false)) << 8)

	return 0
}

func (i *instructionSet) bvc() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(V) == 0 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) bvs() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	if regSet.getFlag(V) == 1 {
		manEl.cycles++
		manEl.addrAbs = regSet.Pc + manEl.addrRel

		if (manEl.addrAbs & 0xFF00) != (regSet.Pc & 0xFF00) {
			manEl.cycles++
		}

		regSet.Pc = manEl.addrAbs
	}
	return 0
}

func (i *instructionSet) clc() uint8 {
	regSet := i.regSet

	regSet.setFlag(C, false)
	return 0
}

func (i *instructionSet) cld() uint8 {
	regSet := i.regSet

	regSet.setFlag(D, false)
	return 0
}

func (i *instructionSet) cli() uint8 {
	regSet := i.regSet

	regSet.setFlag(I, false)
	return 0
}

func (i *instructionSet) clv() uint8 {
	regSet := i.regSet

	regSet.setFlag(V, false)
	return 0
}

func (i *instructionSet) cmp() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	temp := uint16(regSet.A) - uint16(manEl.fetched)
	regSet.setFlag(C, regSet.A >= manEl.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 1
}

func (i *instructionSet) cpx() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	temp := uint16(regSet.X) - uint16(manEl.fetched)
	regSet.setFlag(C, regSet.X >= manEl.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) cpy() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	temp := uint16(regSet.Y) - uint16(manEl.fetched)
	regSet.setFlag(C, regSet.Y >= manEl.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) dec() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	temp := uint16(manEl.fetched) - 1
	mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) dex() uint8 {
	regSet := i.regSet

	regSet.X--
	regSet.setFlag(Z, regSet.X == 0x00)
	regSet.setFlag(N, (regSet.X & 0x80) > 0)
	return 0
}

func (i *instructionSet) dey() uint8 {
	regSet := i.regSet

	regSet.Y--
	regSet.setFlag(Z, regSet.Y == 0x00)
	regSet.setFlag(N, (regSet.Y & 0x80) > 0)
	return 0
}

func (i *instructionSet) eor() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.A = regSet.A ^ manEl.fetched
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A & 0x80 != 0)
	return 1
}

func (i *instructionSet) inc() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	temp := uint16(manEl.fetched) + 1
	mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) inx() uint8 {
	regSet := i.regSet

	regSet.X++
	regSet.setFlag(Z, regSet.X == 0x00)
	regSet.setFlag(N, (regSet.X & 0x80) > 0)
	return 0
}

func (i *instructionSet) iny() uint8 {
	regSet := i.regSet

	regSet.Y++
	regSet.setFlag(Z, regSet.Y == 0x00)
	regSet.setFlag(N, (regSet.Y & 0x80) > 0)
	return 0
}

func (i *instructionSet) jmp() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	regSet.Pc = manEl.addrAbs
	return 0
}

func (i *instructionSet) jsr() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	regSet.Pc--

	mBus.Write(BASE_STKP + uint16(regSet.Stkp), uint8((regSet.Pc >> 8) & 0x00FF))
	regSet.Stkp--
	mBus.Write(BASE_STKP + uint16(regSet.Stkp), uint8(regSet.Pc & 0x00FF))
	regSet.Stkp--

	regSet.Pc = manEl.addrAbs
	return 0
}

func (i *instructionSet) lda() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.A = manEl.fetched
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, (regSet.A & 0x80) > 0)
	return 1
}

func (i *instructionSet) ldx() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.X = manEl.fetched
	regSet.setFlag(Z, regSet.X == 0x00)
	regSet.setFlag(N, (regSet.X & 0x80) > 0)
	return 1
}

func (i *instructionSet) ldy() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.Y = manEl.fetched
	regSet.setFlag(Z, regSet.Y == 0x00)
	regSet.setFlag(N, (regSet.Y & 0x80) > 0)
	return 1
}

func (i *instructionSet) lsr() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	regSet.setFlag(C, manEl.fetched & 0x0001 > 0)
	temp := manEl.fetched >> 1
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := manEl.lookup[manEl.opcode].addrmodeName

	if addrmodeName == "imp" {
		regSet.A = uint8(temp & 0x00FF)
	} else {
		mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) nop() uint8 {
	switch i.manEl.opcode {
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

func (i *instructionSet) ora() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	regSet.A = regSet.A | manEl.fetched
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A == 0x80)
	return 1
}

func (i *instructionSet) pha() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	mBus.Write(BASE_STKP + uint16(regSet.Stkp), regSet.A)
	regSet.Stkp--
	return 0
}

func (i *instructionSet) php() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	mBus.Write(BASE_STKP + uint16(regSet.Stkp), regSet.Status | uint8(B) | uint8(U))
	regSet.setFlag(B, false)
	regSet.setFlag(U, false)
	regSet.Stkp--
	return 0
}

func (i *instructionSet) pla() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	regSet.Stkp++
	regSet.A = mBus.Read(BASE_STKP + uint16(regSet.Stkp), false)
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A & 0x80 > 0)
	return 0
}

func (i *instructionSet) plp() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	regSet.Stkp++
	regSet.Status = mBus.Read(BASE_STKP + uint16(regSet.Stkp), false)
	regSet.setFlag(U, true)
	return 0
}

func (i *instructionSet) rol() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	temp := (uint16(manEl.fetched) << 1) | uint16(regSet.getFlag(C))

	regSet.setFlag(C, temp & 0xFF00 > 0)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := manEl.lookup[manEl.opcode].addrmodeName

	if addrmodeName == "imp" {
		regSet.A = uint8(temp & 0x00FF)
	} else {
		mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) ror() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	manEl.fetch()
	temp := (uint16(regSet.getFlag(C)) << 7) | (uint16(manEl.fetched) >> 1)

	regSet.setFlag(C, manEl.fetched & 0x01 > 0)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := manEl.lookup[manEl.opcode].addrmodeName

	if addrmodeName == "imp" {
		regSet.A = uint8(temp & 0x00FF)
	} else {
		mBus.Write(manEl.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) rti() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	regSet.Stkp++
	regSet.Status = mBus.Read(BASE_STKP + uint16(regSet.Stkp), false)
	regSet.Status &= ^uint8(B)
	regSet.Status &= ^uint8(U)

	regSet.Stkp++
	regSet.Pc = uint16(mBus.Read(BASE_STKP + uint16(regSet.Stkp), false))
	regSet.Stkp++
	regSet.Pc |= uint16(mBus.Read(BASE_STKP + uint16(regSet.Stkp), false)) << 8
	return 0
}

func (i *instructionSet) rts() uint8 {
	regSet := i.regSet
	mBus := i.mBus

	regSet.Stkp++
	regSet.Pc = uint16(mBus.Read(BASE_STKP + uint16(regSet.Stkp), false))
	regSet.Stkp++
	regSet.Pc |= uint16(mBus.Read(BASE_STKP + uint16(regSet.Stkp), false)) << 8

	regSet.Pc++

	return 0
}

func (i *instructionSet) sbc() uint8 {
	regSet := i.regSet
	manEl := i.manEl

	manEl.fetch()
	invFetch := uint16(manEl.fetched & 0x00FF)
	val := uint16(regSet.A) + invFetch + uint16(regSet.getFlag(C))
	regSet.setFlag(C, val & 0xFF00 > 0)
	regSet.setFlag(Z, (val & 0x00FF) == 0)
	regSet.setFlag(N, (val & 0x0080) > 0)
	vFlagVal := (val ^ uint16(regSet.A)) & (val ^ invFetch) & 0x0080
	regSet.setFlag(V, vFlagVal > 0)
	regSet.A = uint8(val & 0x00FF)
	return 1
}

func (i *instructionSet) sec() uint8 {
	regSet := i.regSet

	regSet.setFlag(C, true)
	return 0
}

func (i *instructionSet) sed() uint8 {
	regSet := i.regSet

	regSet.setFlag(D, true)
	return 0
}

func (i *instructionSet) sei() uint8 {
	regSet := i.regSet

	regSet.setFlag(I, true)
	return 0
}

func (i *instructionSet) sta() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	mBus.Write(manEl.addrAbs, regSet.A)
	return 0
}

func (i *instructionSet) stx() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	mBus.Write(manEl.addrAbs, regSet.X)
	return 0
}

func (i *instructionSet) sty() uint8 {
	regSet := i.regSet
	manEl := i.manEl
	mBus := i.mBus

	mBus.Write(manEl.addrAbs, regSet.Y)
	return 0
}

func (i *instructionSet) tax() uint8 {
	regSet := i.regSet

	regSet.X = regSet.A
	regSet.setFlag(Z, regSet.X == 0x00)
	regSet.setFlag(N, regSet.X & 0x80 > 0)
	return 0
}

func (i *instructionSet) tay() uint8 {
	regSet := i.regSet

	regSet.Y = regSet.A
	regSet.setFlag(Z, regSet.Y == 0x00)
	regSet.setFlag(N, regSet.Y & 0x80 > 0)
	return 0
}

func (i *instructionSet) tsx() uint8 {
	regSet := i.regSet

	regSet.X = regSet.Stkp
	regSet.setFlag(Z, regSet.X == 0x00)
	regSet.setFlag(N, regSet.X & 0x80 > 0)
	return 0
}

func (i *instructionSet) txa() uint8 {
	regSet := i.regSet

	regSet.A = regSet.X
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A & 0x80 > 0)
	return 0
}

func (i *instructionSet) txs() uint8 {
	regSet := i.regSet

	regSet.Stkp = regSet.X
	return 0
}

func (i *instructionSet) tya() uint8 {
	regSet := i.regSet

	regSet.A = regSet.Y
	regSet.setFlag(Z, regSet.A == 0x00)
	regSet.setFlag(N, regSet.A & 0x80 > 0)
	return 0
}

func (i *instructionSet) xxx() uint8 {
	return 0
}

func CreateInstructionSet(
	mBus *bus, regSet *registerSet, manEl *managingElement,
) *instructionSet {
	i := &instructionSet{}
	i.fillInstrByName()

	i.mBus = mBus
	i.regSet = regSet
	i.manEl = manEl
	
	return i
}

