package olcCpu

type instructionSet struct {
	cpu *olc6502
	instrByName map[string]func() uint8
}

type InstructionSet interface {
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
}


func (i *instructionSet) fillInstrByName() {
	i.instrByName = make(map[string]func() uint8)

	i.instrByName["ADC"] = i.ADC
	i.instrByName["AND"] = i.AND
	i.instrByName["ASL"] = i.ASL
	i.instrByName["BCC"] = i.BCC
	i.instrByName["BCS"] = i.BCS
	i.instrByName["BEQ"] = i.BEQ
	i.instrByName["BIT"] = i.BIT
	i.instrByName["BMI"] = i.BMI
	i.instrByName["BNE"] = i.BNE
	i.instrByName["BPL"] = i.BPL
	i.instrByName["BRK"] = i.BRK
	i.instrByName["BVC"] = i.BVC
	i.instrByName["BVS"] = i.BVS
	i.instrByName["CLC"] = i.CLC
	i.instrByName["CLD"] = i.CLD
	i.instrByName["CLI"] = i.CLI
	i.instrByName["CLV"] = i.CLV
	i.instrByName["CMP"] = i.CMP
	i.instrByName["CPX"] = i.CPX
	i.instrByName["CPY"] = i.CPY
	i.instrByName["DEC"] = i.DEC
	i.instrByName["DEX"] = i.DEX
	i.instrByName["DEY"] = i.DEY
	i.instrByName["EOR"] = i.EOR
	i.instrByName["INC"] = i.INC
	i.instrByName["INX"] = i.INX
	i.instrByName["INY"] = i.INY
	i.instrByName["JMP"] = i.JMP
	i.instrByName["JSR"] = i.JSR
	i.instrByName["LDA"] = i.LDA
	i.instrByName["LDX"] = i.LDX
	i.instrByName["LDY"] = i.LDY
	i.instrByName["LSR"] = i.LSR
	i.instrByName["NOP"] = i.NOP
	i.instrByName["ORA"] = i.ORA
	i.instrByName["PHA"] = i.PHA
	i.instrByName["PHP"] = i.PHP
	i.instrByName["PLA"] = i.PLA
	i.instrByName["PLP"] = i.PLP
	i.instrByName["ROL"] = i.ROL
	i.instrByName["ROR"] = i.ROR
	i.instrByName["RTI"] = i.RTI
	i.instrByName["RTS"] = i.RTS
	i.instrByName["SBC"] = i.SBC
	i.instrByName["SEC"] = i.SEC
	i.instrByName["SED"] = i.SED
	i.instrByName["SEI"] = i.SEI
	i.instrByName["STA"] = i.STA
	i.instrByName["STX"] = i.STX
	i.instrByName["STY"] = i.STY
	i.instrByName["TAX"] = i.TAX
	i.instrByName["TAY"] = i.TAY
	i.instrByName["TSX"] = i.TSX
	i.instrByName["TXA"] = i.TXA
	i.instrByName["TXS"] = i.TXS
	i.instrByName["TYA"] = i.TYA
	i.instrByName["XXX"] = i.XXX

}

func (i *instructionSet) ADC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu
	
	cpu.fetch()
	temp := uint16(regSet.a) + uint16(cpu.fetched) + uint16(regSet.getFlag(C))
	regSet.setFlag(C, temp > 255)
	regSet.setFlag(Z, (temp & 0x00FF) == 0)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	vFlagVal := (
		^(uint16(regSet.a) ^ uint16(cpu.fetched)) &
		(uint16(regSet.a) ^ uint16(temp))) & 0x0080
	regSet.setFlag(V, vFlagVal > 0)

	regSet.a = uint8(temp & 0x00FF)
	return 1
}

func (i *instructionSet) AND() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.a = regSet.a & cpu.fetched
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a & 0x80 != 0)
	return 1
}

func (i *instructionSet) ASL() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(cpu.fetched) << 1

	regSet.setFlag(C, temp & 0xFF00 > 0)
	regSet.setFlag(Z, temp & 0x00FF == 0x00)
	regSet.setFlag(N, temp & 0x80 > 0)

	addrmodeName := cpu.lookup[cpu.opcode].addrmodeName

	if addrmodeName == "IMP" {
		regSet.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}
	return 0
}

func (i *instructionSet) BCC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(C) == 0 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BCS() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(C) == 1 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BEQ() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(Z) == 1 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BIT() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := regSet.a & cpu.fetched
	regSet.setFlag(Z, (temp & 0x00FF) == 0x00)
	regSet.setFlag(N, cpu.fetched & (1 << 7) > 0)
	regSet.setFlag(V, cpu.fetched & (1 << 6) > 0)
	return 0
}

func (i *instructionSet) BMI() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(N) == 1 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BNE() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(Z) == 0 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BPL() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(N) == 0 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BRK() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	//regSet.pc++

	regSet.setFlag(I, true)
	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8((regSet.pc >> 8) & 0x00FF))
	regSet.stkp--
	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8((regSet.pc) & 0x00FF))
	regSet.stkp--

	regSet.setFlag(B, true)
	cpu.Write(BASE_STKP + uint16(regSet.stkp), regSet.status)
	regSet.stkp--
	regSet.setFlag(B, false)

	regSet.pc = uint16(cpu.Read(0xFFFE)) | (uint16(cpu.Read(0xFFFF)) << 8)

	return 0
}

func (i *instructionSet) BVC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(V) == 0 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) BVS() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	if regSet.getFlag(V) == 1 {
		cpu.cycles++
		cpu.addrAbs = regSet.pc + cpu.addrRel

		if (cpu.addrAbs & 0xFF00) != (regSet.pc & 0xFF00) {
			cpu.cycles++
		}

		regSet.pc = cpu.addrAbs
	}
	return 0
}

func (i *instructionSet) CLC() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(C, false)
	return 0
}

func (i *instructionSet) CLD() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(D, false)
	return 0
}

func (i *instructionSet) CLI() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(I, false)
	return 0
}

func (i *instructionSet) CLV() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(V, false)
	return 0
}

func (i *instructionSet) CMP() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(regSet.a) - uint16(cpu.fetched)
	regSet.setFlag(C, regSet.a >= cpu.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 1
}

func (i *instructionSet) CPX() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(regSet.x) - uint16(cpu.fetched)
	regSet.setFlag(C, regSet.x >= cpu.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) CPY() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(regSet.y) - uint16(cpu.fetched)
	regSet.setFlag(C, regSet.y >= cpu.fetched)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) DEC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(cpu.fetched) - 1
	cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) DEX() uint8 {
	regSet := i.cpu.regSet

	regSet.x--
	regSet.setFlag(Z, regSet.x == 0x00)
	regSet.setFlag(N, (regSet.x & 0x80) > 0)
	return 0
}

func (i *instructionSet) DEY() uint8 {
	regSet := i.cpu.regSet

	regSet.y--
	regSet.setFlag(Z, regSet.y == 0x00)
	regSet.setFlag(N, (regSet.y & 0x80) > 0)
	return 0
}

func (i *instructionSet) EOR() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.a = regSet.a ^ cpu.fetched
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a & 0x80 != 0)
	return 1
}

func (i *instructionSet) INC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := uint16(cpu.fetched) + 1
	cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)
	return 0
}

func (i *instructionSet) INX() uint8 {
	regSet := i.cpu.regSet

	regSet.x++
	regSet.setFlag(Z, regSet.x == 0x00)
	regSet.setFlag(N, (regSet.x & 0x80) > 0)
	return 0
}

func (i *instructionSet) INY() uint8 {
	regSet := i.cpu.regSet

	regSet.y++
	regSet.setFlag(Z, regSet.y == 0x00)
	regSet.setFlag(N, (regSet.y & 0x80) > 0)
	return 0
}

func (i *instructionSet) JMP() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.pc = cpu.addrAbs
	return 0
}

func (i *instructionSet) JSR() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.pc--

	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8((regSet.pc >> 8) & 0x00FF))
	regSet.stkp--
	cpu.Write(BASE_STKP + uint16(regSet.stkp), uint8(regSet.pc & 0x00FF))
	regSet.stkp--

	regSet.pc = cpu.addrAbs
	return 0
}

func (i *instructionSet) LDA() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.a = cpu.fetched
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, (regSet.a & 0x80) > 0)
	return 1
}

func (i *instructionSet) LDX() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.x = cpu.fetched
	regSet.setFlag(Z, regSet.x == 0x00)
	regSet.setFlag(N, (regSet.x & 0x80) > 0)
	return 1
}

func (i *instructionSet) LDY() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.y = cpu.fetched
	regSet.setFlag(Z, regSet.y == 0x00)
	regSet.setFlag(N, (regSet.y & 0x80) > 0)
	return 1
}

func (i *instructionSet) LSR() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.setFlag(C, cpu.fetched & 0x0001 > 0)
	temp := cpu.fetched >> 1
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := cpu.lookup[cpu.opcode].addrmodeName

	if addrmodeName == "IMP" {
		regSet.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) NOP() uint8 {
	switch i.cpu.opcode {
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

func (i *instructionSet) ORA() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	regSet.a = regSet.a | cpu.fetched
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a == 0x80)
	return 1
}

func (i *instructionSet) PHA() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.Write(BASE_STKP + uint16(regSet.stkp), regSet.a)
	regSet.stkp--
	return 0
}

func (i *instructionSet) PHP() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.Write(BASE_STKP + uint16(regSet.stkp), regSet.status | uint8(B) | uint8(U))
	regSet.setFlag(B, false)
	regSet.setFlag(U, false)
	regSet.stkp--
	return 0
}

func (i *instructionSet) PLA() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.stkp++
	regSet.a = cpu.Read(BASE_STKP + uint16(regSet.stkp))
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a & 0x80 > 0)
	return 0
}

func (i *instructionSet) PLP() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.stkp++
	regSet.status = cpu.Read(BASE_STKP + uint16(regSet.stkp))
	regSet.setFlag(U, true)
	return 0
}

func (i *instructionSet) ROL() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := (uint16(cpu.fetched) << 1) | uint16(regSet.getFlag(C))

	regSet.setFlag(C, temp & 0xFF00 > 0)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := cpu.lookup[cpu.opcode].addrmodeName

	if addrmodeName == "IMP" {
		regSet.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) ROR() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	temp := (uint16(regSet.getFlag(C)) << 7) | (uint16(cpu.fetched) >> 1)

	regSet.setFlag(C, cpu.fetched & 0x01 > 0)
	regSet.setFlag(Z, (temp & 0x00FF) == 0x0000)
	regSet.setFlag(N, (temp & 0x0080) > 0)

	addrmodeName := cpu.lookup[cpu.opcode].addrmodeName

	if addrmodeName == "IMP" {
		regSet.a = uint8(temp & 0x00FF)
	} else {
		cpu.Write(cpu.addrAbs, uint8(temp & 0x00FF))
	}

	return 0
}

func (i *instructionSet) RTI() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.stkp++
	regSet.status = cpu.Read(BASE_STKP + uint16(regSet.stkp))
	regSet.status &= ^uint8(B)
	regSet.status &= ^uint8(U)

	regSet.stkp++
	regSet.pc = uint16(cpu.Read(BASE_STKP + uint16(regSet.stkp)))
	regSet.stkp++
	regSet.pc |= uint16(cpu.Read(BASE_STKP + uint16(regSet.stkp))) << 8
	return 0
}

func (i *instructionSet) RTS() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	regSet.stkp++
	regSet.pc = uint16(cpu.Read(BASE_STKP + uint16(regSet.stkp)))
	regSet.stkp++
	regSet.pc |= uint16(cpu.Read(BASE_STKP + uint16(regSet.stkp))) << 8

	regSet.pc++

	return 0
}

func (i *instructionSet) SBC() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.fetch()
	invFetch := uint16(cpu.fetched & 0x00FF)
	val := uint16(regSet.a) + invFetch + uint16(regSet.getFlag(C))
	regSet.setFlag(C, val & 0xFF00 > 0)
	regSet.setFlag(Z, (val & 0x00FF) == 0)
	regSet.setFlag(N, (val & 0x0080) > 0)
	vFlagVal := (val ^ uint16(regSet.a)) & (val ^ invFetch) & 0x0080
	regSet.setFlag(V, vFlagVal > 0)
	regSet.a = uint8(val & 0x00FF)
	return 1
}

func (i *instructionSet) SEC() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(C, true)
	return 0
}

func (i *instructionSet) SED() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(D, true)
	return 0
}

func (i *instructionSet) SEI() uint8 {
	regSet := i.cpu.regSet

	regSet.setFlag(I, true)
	return 0
}

func (i *instructionSet) STA() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.Write(cpu.addrAbs, regSet.a)
	return 0
}

func (i *instructionSet) STX() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.Write(cpu.addrAbs, regSet.x)
	return 0
}

func (i *instructionSet) STY() uint8 {
	regSet := i.cpu.regSet
	cpu := i.cpu

	cpu.Write(cpu.addrAbs, regSet.y)
	return 0
}

func (i *instructionSet) TAX() uint8 {
	regSet := i.cpu.regSet

	regSet.x = regSet.a
	regSet.setFlag(Z, regSet.x == 0x00)
	regSet.setFlag(N, regSet.x & 0x80 > 0)
	return 0
}

func (i *instructionSet) TAY() uint8 {
	regSet := i.cpu.regSet

	regSet.y = regSet.a
	regSet.setFlag(Z, regSet.y == 0x00)
	regSet.setFlag(N, regSet.y & 0x80 > 0)
	return 0
}

func (i *instructionSet) TSX() uint8 {
	regSet := i.cpu.regSet

	regSet.x = regSet.stkp
	regSet.setFlag(Z, regSet.x == 0x00)
	regSet.setFlag(N, regSet.x & 0x80 > 0)
	return 0
}

func (i *instructionSet) TXA() uint8 {
	regSet := i.cpu.regSet

	regSet.a = regSet.x
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a & 0x80 > 0)
	return 0
}

func (i *instructionSet) TXS() uint8 {
	regSet := i.cpu.regSet

	regSet.stkp = regSet.x
	return 0
}

func (i *instructionSet) TYA() uint8 {
	regSet := i.cpu.regSet

	regSet.a = regSet.y
	regSet.setFlag(Z, regSet.a == 0x00)
	regSet.setFlag(N, regSet.a & 0x80 > 0)
	return 0
}

func (i *instructionSet) XXX() uint8 {
	return 0
}

func CreateInstructionSet(cpu *olc6502) *instructionSet {
	i := &instructionSet{}
	i.cpu = cpu
	i.fillInstrByName()

	i.cpu.connectInstructionSet(i)

	return i
}

