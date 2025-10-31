package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	a := initializeAssembler()

	err := a.firstPass()

	if err != nil {
		a.errLogger.Fatalln(err.Error())
	}

	err = a.secondPass()

	if err != nil {
		a.errLogger.Fatalln(err.Error())
	}

	/* for debug purposes
	a.initializeScanner()
	a.currentLineNumber = 0

	fmt.Printf("%4s %-10s   %-15s %10s %10s %10s%10s%20s\n", "LINE", "INSTR", "TYPE", "SYMBOL", "DEST", "COMP", "JUMP", "ENCODED")

	for {

		err := a.advance()

		if err != nil {
			break
		}

		fmt.Printf("%-4d %-10s   %-15s %10s %10s %10s%10s%20s\n", a.currentLineNumber, a.currentLine.text, a.currentLine.iType, a.currentLine.symbol, a.currentLine.dest, a.currentLine.comp, a.currentLine.jump, a.encode())

	}

	*/

	outFileName := strings.Split(os.Args[1], ".")[0] + ".hack"

	fmt.Println("Assembly is done.")
	fmt.Printf("Output file: %s\n", outFileName)


}
