package main

import (
	"fmt"
	"io/ioutil"
)

type CPU struct {
	a, b, c, d, e, f byte
	h, l             byte
	pc               int  // program counter
	zf               bool // zero flag
	sf               bool
	hcf              bool
	cf               bool
	cycles           int
}

func interp(cpu *CPU, data []byte) {
	count := 0
	for {
		curInstr := data[cpu.pc]

		switch curInstr {
		case 0x00:
			// nop
			cpu.cycles += 1
			cpu.pc++

			// LD nn,n
		case 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e:
			cpu.pc++
			if curInstr == 0x06 {
				cpu.b = data[cpu.pc]
			} else if curInstr == 0x0e {
				cpu.c = data[cpu.pc]
			} else if curInstr == 0x16 {
				cpu.d = data[cpu.pc]
			} else if curInstr == 0x1e {
				cpu.e = data[cpu.pc]
			} else if curInstr == 0x26 {
				cpu.h = data[cpu.pc]
			} else if curInstr == 0x2e {
				cpu.l = data[cpu.pc]
			}
			cpu.pc++
			// JP nn
		case 0xc3:
			// LS byte first
			cpu.pc++
			byte0 := data[cpu.pc]
			cpu.pc++
			byte1 := data[cpu.pc]
			cpu.pc++
			var jmpVal = int(byte0) | int(byte1)<<8
			fmt.Printf("jmp nn %d %d %d\n", byte0, byte1, jmpVal)

		default:
		} // end of switch

		if count == 0 {
			break
		}
	} // end of main loop
}

func main() {

	data, err := ioutil.ReadFile("tetris_world.gb")
	if err != nil {
		fmt.Println("could not read file", err)
		return
	}

	fmt.Println("len of data read from file", len(data))

	fmt.Printf("byte0: %0x\n", data[0x100])
	var cpu CPU
	cpu.pc = 0x100
	interp(&cpu, data)
}
