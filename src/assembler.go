package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type assembler struct {
	errLogger         *log.Logger
	inputFile       *bufio.Scanner
	outputFile		*os.File
	currentLineNumber int
	currentLine       instructionLine
	symbolTable       map[string]int
	firstPassed       bool
	lastVariableIndex int
}

func (a *assembler) advance() error {

	a.currentLineNumber++

	for {

		haveNext := a.inputFile.Scan()

		// if there isn't more lines
		if !haveNext {

			// if end of the file is reached
			if a.inputFile.Err() == nil {
				return errors.New("EOF")

				// if an actual error has occured
			} else {
				return a.inputFile.Err()
			}
		}

		a.currentLine.text = a.inputFile.Text()
		a.currentLine.text = strings.TrimSpace(a.currentLine.text)

		// if line contains only the whitespace
		if a.currentLine.text == "" || a.currentLine.isComment() {
			continue
		}

		// if line isn't empty
		break

	}

	a.parseInstruction()

	return nil

}
