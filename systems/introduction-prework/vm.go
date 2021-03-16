package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
	Addi  = 0x05
	Subi  = 0x06
	Jump  = 0x07
	Beqz  = 0x08
)

const instructionLength byte = 3

func compute(memory []byte) {
	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	pc := registers[0]
	for {
		op := memory[pc]

		switch op {
		case Load:
			registers[memory[pc+1]] = memory[memory[pc+2]]
		case Store:
			memory[memory[pc+2]] = registers[memory[pc+1]]
		case Add:
			registers[memory[pc+1]] = registers[memory[pc+1]] + registers[memory[pc+2]]
		case Sub:
			registers[memory[pc+1]] = registers[memory[pc+1]] - registers[memory[pc+2]]
		case Addi:
			registers[memory[pc+1]] = registers[memory[pc+1]] + memory[pc+2]
		case Subi:
			registers[memory[pc+1]] = registers[memory[pc+1]] - memory[pc+2]
		case Jump:
			pc = memory[pc+1]
			continue
		case Beqz:
			if registers[memory[pc+1]] == 0 {
				pc += memory[pc+2]
			}
		case Halt:
			return
		}

		pc += instructionLength
	}
}
