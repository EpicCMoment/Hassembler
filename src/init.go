package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func (a *assembler) initializeScanner() {

	if len(os.Args) < 2 {
		a.errLogger.Fatalln("no input file")
	}

	fileContent, err := os.ReadFile(os.Args[1])

	if err != nil {
		a.errLogger.Fatalf("unable to read file: %s\n", err.Error())
	}

	tempReader := strings.NewReader(string(fileContent))

	a.inputFile = bufio.NewScanner(tempReader)

}

func (a *assembler) initializeWriter() {

	fileWoutExtension := strings.Split(os.Args[1], ".")[0]

	outFileName := fileWoutExtension + ".hack"

	outFile, err := os.Create(outFileName)

	if err != nil {
		a.errLogger.Fatalln(err.Error())
	}

	a.outputFile = outFile

}

func initializeAssembler() assembler {

	ass := assembler{}

	ass.errLogger = log.New(os.Stderr, "[ERROR] ", log.Llongfile|log.LUTC)
	ass.initializeScanner()
	ass.initializeWriter()
	ass.currentLineNumber = 0
	ass.firstPassed = false
	ass.lastVariableIndex = 16
	ass.symbolTable = make(map[string]int)

	// predefined symbols
	ass.symbolTable["R0"] = 0
	ass.symbolTable["R1"] = 1
	ass.symbolTable["R2"] = 2
	ass.symbolTable["R3"] = 3

	ass.symbolTable["R4"] = 4
	ass.symbolTable["R5"] = 5
	ass.symbolTable["R6"] = 6
	ass.symbolTable["R7"] = 7

	ass.symbolTable["R8"] = 8
	ass.symbolTable["R9"] = 9
	ass.symbolTable["R10"] = 10
	ass.symbolTable["R11"] = 11

	ass.symbolTable["R12"] = 12
	ass.symbolTable["R13"] = 13
	ass.symbolTable["R14"] = 14
	ass.symbolTable["R15"] = 15

	ass.symbolTable["SP"] = 0
	ass.symbolTable["LCL"] = 1
	ass.symbolTable["ARG"] = 2
	ass.symbolTable["THIS"] = 3
	ass.symbolTable["THAT"] = 4

	ass.symbolTable["SCREEN"] = 16384
	ass.symbolTable["KBD"] = 24576

	return ass
}
