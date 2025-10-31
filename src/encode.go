package main

import (
	"fmt"
	"strconv"
	"strings"
)

type encodingTables struct {
	initialized bool
	compTable map[string]string
	destTable map[string]string
	jumpTable map[string]string
}

var tables encodingTables


func (a *assembler) encode() string {

	if !tables.initialized {
		initializeTables()
	}

	var encodedInstruction string

	switch a.currentLine.iType {
		case A_INSTRUCTION:
			encodedInstruction = a.encodeA()

		case C_INSTRUCTION:
			encodedInstruction = a.encodeC()

		case L_INSTRUCTION:
			encodedInstruction = a.encodeL()
	}

	return encodedInstruction


}

func (a *assembler) encodeA() string {

	if a.currentLine.iType != A_INSTRUCTION {
		panic("invalid instruction type. probably wrong encode function is called")
	}

	symbolAddress, exists := a.symbolTable[a.currentLine.symbol]

	var value string

	// if not exists in the symbol table, it is a constant number
	if exists {
		value = strconv.FormatInt(int64(symbolAddress), 2)
	} else {
		intValue, _ := strconv.ParseInt(a.currentLine.symbol, 10, 32)
		value = strconv.FormatInt(intValue, 2)
	}

	// format the binary instruction
	encodedInstruction := fmt.Sprintf("0%015s", value)

	return encodedInstruction

}


func (ass *assembler) encodeC() string {

	if ass.currentLine.iType != C_INSTRUCTION {
		panic("invalid instruction type. probably wrong encode function is called")
	}

	var a string

	if strings.Contains(ass.currentLine.comp, "M") {
		a = "1"
	} else {
		a = "0"
	}

	// encoded comp
	encComp := tables.compTable[ass.currentLine.comp]
	
	// encoded dest
	encDest := tables.destTable[ass.currentLine.dest]

	// encoded jump
	encJump := tables.jumpTable[ass.currentLine.jump]

	encodedInstruction := fmt.Sprintf(
		"111%s%s%s%s",
		a,
		encComp,
		encDest,
		encJump,
	)

	return encodedInstruction

}

// L instructions doesn't encode instructions
// this function stands for syntactic beauty
func (a *assembler) encodeL() string {
	return ""
}

func initializeTables() {

	initializeCompTable()
	initializeDestTable()
	initializeJumpTable()

	tables.initialized = true

}

func initializeCompTable() {

	tables.compTable = make(map[string]string)

	tables.compTable["0"] = "101010"
	tables.compTable["1"] = "111111"
	tables.compTable["-1"] = "111010"

	tables.compTable["D"] = "001100"
	tables.compTable["A"] = "110000"
	tables.compTable["M"] = "110000"
	
	tables.compTable["!D"] = "001101"
	tables.compTable["!A"] = "110001"
	tables.compTable["!M"] = "110001"
	
	tables.compTable["-D"] = "001111"
	tables.compTable["-A"] = "110011"
	tables.compTable["-M"] = "110011"
	
	tables.compTable["D+1"] = "011111"
	tables.compTable["A+1"] = "110111"
	tables.compTable["M+1"] = "110111"
	
	tables.compTable["D-1"] = "001110"
	tables.compTable["A-1"] = "110010"
	tables.compTable["M-1"] = "110010"
	
	tables.compTable["D+A"] = "000010"
	tables.compTable["D+M"] = "000010"

	tables.compTable["D-A"] = "010011"
	tables.compTable["D-M"] = "010011"

	tables.compTable["A-D"] = "000111"
	tables.compTable["M-D"] = "000111"

	tables.compTable["D&A"] = "000000"
	tables.compTable["D&M"] = "000000"

	tables.compTable["D|A"] = "010101"
	tables.compTable["D|M"] = "010101"


}

func initializeDestTable() {

	tables.destTable = make(map[string]string)

	tables.destTable[""] = "000"
	tables.destTable["M"] = "001"
	
	tables.destTable["D"] = "010"
	tables.destTable["DM"] = "011"
	tables.destTable["MD"] = "011"

	tables.destTable["A"] = "100"
	tables.destTable["AM"] = "101"
	tables.destTable["MA"] = "101"

	tables.destTable["AD"] = "110"
	tables.destTable["DA"] = "110"

	
	tables.destTable["ADM"] = "111"
	tables.destTable["AMD"] = "111"
	tables.destTable["MDA"] = "111"
	tables.destTable["MAD"] = "111"
	tables.destTable["DAM"] = "111"
	tables.destTable["DMA"] = "111"

}

func initializeJumpTable() {

	tables.jumpTable = make(map[string]string)

	tables.jumpTable[""] = "000"
	tables.jumpTable["JGT"] = "001"
	
	tables.jumpTable["JEQ"] = "010"
	tables.jumpTable["JGE"] = "011"

	tables.jumpTable["JLT"] = "100"
	tables.jumpTable["JNE"] = "101"

	tables.jumpTable["JLE"] = "110"
	tables.jumpTable["JMP"] = "111"

}

