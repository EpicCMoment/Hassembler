package main

import (
	"errors"
	"strconv"
)

// first pass of the file
func (a *assembler) firstPass() error {

	if a.firstPassed {
		return errors.New("not the first pass")
	}

	for {

		err := a.advance()

		if err != nil {
			break
		}

		if a.currentLine.iType == L_INSTRUCTION {
			a.symbolTable[a.currentLine.symbol] = a.currentLineNumber

		}

	}

	a.initializeScanner()
	a.currentLineNumber = 0
	a.currentLine = instructionLine{}

	a.firstPassed = true

	return nil

}

func (a *assembler) secondPass() error {

	for {

		err := a.advance()

		if err != nil {
			break
		}

		if a.currentLine.iType == A_INSTRUCTION {
			symbol := a.currentLine.symbol

			_, err := strconv.Atoi(symbol)

			if err != nil {
				_, exists := a.symbolTable[symbol]

				if !exists {
					a.symbolTable[symbol] = a.lastVariableIndex
					a.lastVariableIndex++
				}
			}

		}

		encodedInstruction := a.encode()



		if encodedInstruction == "" {
			continue
		}

		a.outputFile.WriteString(encodedInstruction + "\n")
		/* for debug purposes
		otpt := fmt.Sprintf("%-15s  %s\n", a.currentLine.text, encodedInstruction)
		
		a.outputFile.WriteString(otpt)
		*/
	}



	return nil

}
