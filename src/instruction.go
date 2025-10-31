package main

import "strings"

type instructionType string

const (
	A_INSTRUCTION instructionType = "A_INSTRUCTION"
	C_INSTRUCTION instructionType = "C_INSTRUCTION"
	L_INSTRUCTION instructionType = "L_INSTRUCTION"
)

type instructionLine struct {
	text   string
	iType  instructionType
	symbol string
	dest   string
	comp   string
	jump   string
}

func (a *assembler) parseInstruction() {

	a.currentLine.parseIType()
	a.currentLine.parseSymbol()
	a.currentLine.parseJump()
	a.currentLine.parseDest()
	a.currentLine.parseComp()

}

func (i *instructionLine) parseJump() {

	switch i.iType {

	case C_INSTRUCTION:

		if strings.Contains(i.text, ";") {
			i.jump = strings.SplitN(i.text, ";", 2)[1]
		}

	default:
		i.jump = ""

	}

}

func (i *instructionLine) isComment() bool {

	return strings.Contains(i.text, "//")

}

func (i *instructionLine) parseDest() {

	switch i.iType {

	case C_INSTRUCTION:

		if strings.Contains(i.text, "=") {
			i.dest = strings.SplitN(i.text, "=", 2)[0]
		}

	default:
		i.dest = ""

	}

}

// returns the computation part
func (i *instructionLine) parseComp() {

	switch i.iType {

	case C_INSTRUCTION:

		if strings.Contains(i.text, ";") {
			i.comp = strings.SplitN(i.text, ";", 2)[0]

		} else {
			i.comp = strings.SplitN(i.text, "=", 2)[1]

		}

	default:
		i.comp = ""

	}

}

func (i *instructionLine) parseSymbol() {

	switch i.iType {

	case A_INSTRUCTION:
		i.symbol = strings.TrimPrefix(i.text, "@")

	case L_INSTRUCTION:
		i.symbol = strings.Trim(i.text, "()")

	default:
		i.symbol = ""

	}

}

func (i *instructionLine) parseIType() {

	line := i.text

	if strings.Contains(line, "@") {
		i.iType = A_INSTRUCTION
	} else if strings.Contains(line, "(") {
		i.iType = L_INSTRUCTION
	} else {
		i.iType = C_INSTRUCTION
	}

}
