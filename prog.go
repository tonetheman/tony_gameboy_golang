package main

import (
	"fmt"
	"io/ioutil"
)

type CPU struct {
	a, b, c, d, e, f byte
	pc               int  // program counter
	zf               bool // zero flag
	sf               bool
	hcf              bool
	cf               bool
}

func main() {

	data, err := ioutil.ReadFile("tetris_world.gb")
	if err != nil {
		fmt.Println("could not read file", err)
		return
	}

	fmt.Println("len of data read from file", len(data))

	fmt.Println("byte0", data[0])
	var cpu CPU
	cpu.pc = 0x100

}
