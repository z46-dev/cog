package runner

const BYTES_PER_WORD = 8

const (
	NUM_REGISTERS   = 16
	REGISTER_RAX    = iota // Accumulator
	REGISTER_RBX           // Base
	REGISTER_RCX           // Counter
	REGISTER_RSP           // Stack Pointer
	REGISTER_RBP           // Base Pointer
	REGISTER_RSI           // Source Index
	REGISTER_RDI           // Destination Index
	REGISTER_RDX           // Data
	REGISTER_R8            // General Purpose
	REGISTER_R9            // General Purpose
	REGISTER_R10           // General Purpose
	REGISTER_R11           // General Purpose
	REGISTER_R12           // General Purpose
	REGISTER_R13           // General Purpose
	REGISTER_R14           // General Purpose
	REGISTER_EFLAGS        // Flags
)

const (
	FLAG_CARRY     = 1 << iota // Flag for carry
	FLAG_ZERO                  // Flag for zero
	FLAG_SIGN                  // Flag for sign
	FLAG_OVERFLOW              // Flag for overflow
	FLAG_PARITY                // Flag for parity
	FLAG_DIRECTION             // Flag for direction
	FLAG_INTERRUPT             // Flag for interrupt
	FLAG_IOPL                  // Flag for I/O privilege level
	FLAG_NT                    // Flag for nested task
	FLAG_RF                    // Flag for resume
	FLAG_VM                    // Flag for virtual 8086 mode
	FLAG_AC                    // Flag for alignment check
	FLAG_VIF                   // Flag for virtual interrupt
	FLAG_VIP                   // Flag for virtual interrupt pending
	FLAG_ID                    // Flag for identification
)

const (
	OPCODE_NOP     = iota // No operation (do nothing)
	OPCODE_MOV            // Move <src> <dst>
	OPCODE_ADD            // Add <src> <dst>
	OPCODE_SUB            // Subtract <src> <dst>
	OPCODE_MUL            // Multiply <src> <dst>
	OPCODE_DIV            // Divide <src> <dst>
	OPCODE_AND            // Bitwise AND <src> <dst>
	OPCODE_OR             // Bitwise OR <src> <dst>
	OPCODE_XOR            // Bitwise XOR <src> <dst>
	OPCODE_SHL            // Shift left <src> <dst>
	OPCODE_SHR            // Shift right <src> <dst>
	OPCODE_CMP            // Compare <src> <dst>
	OPCODE_JMP            // Jump <label>
	OPCODE_JE             // Jump if equal <label>
	OPCODE_JNE            // Jump if not equal <label>
	OPCODE_JG             // Jump if greater <label>
	OPCODE_JGE            // Jump if greater or equal <label>
	OPCODE_JL             // Jump if less <label>
	OPCODE_JLE            // Jump if less or equal <label>
	OPCODE_PUSH           // Push <src>
	OPCODE_POP            // Pop <dst>
	OPCODE_CALL           // Call <label>
	OPCODE_RET            // Return
	OPCODE_HLT            // Halt
	OPCODE_SYSCALL        // System call <num> <...>
	OPCODE_SYSEXIT        // System exit
)

/**
 * Processor architecture:
 *
 * Use little-endian for all data
 * Data can be represented by taking up the number of bytes needed in the 8-byte word
 *
 * Instructions are formatted as follows:
 *  - 1 byte opcode
 *  - 1-8 bytes for each operand, depending on the instruction
 *
 * Registers are formatted as follows:
 *  - 8 bytes for each register
 *  - 16 registers
 */

type Processor struct {
	InUse, SupportsSIMD     bool
	Registers, Instructions []uint8
}

func NewProcessor(supportsSIMD bool, numRegisters int) *Processor {
	return &Processor{
		InUse:        false,
		SupportsSIMD: supportsSIMD,
		Registers:    make([]uint8, numRegisters*BYTES_PER_WORD),
	}
}

type Memory struct {
	RawBytes []byte
}

type Runner struct {
	Processors []*Processor
	Memory     *Memory
}
