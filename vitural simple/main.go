package main

import "fmt"

const (
	VM_OK     = 0
	VM_ERROR  = 1  // vm error
	STACK_MAX = 50 // limit stack

	HLT = 0 // out
	PSH = 1 // push
	POP = 2 // pop
	PUT = 3 // print
	ADD = 4 // add
	SUB = 5 // sub
	MUL = 6 // mul
	DIV = 7 // div
)

type (
	Value float64
	Code  []int
	VM    struct {
		code  Code             // syntax input
		stack [STACK_MAX]Value // stack
		top   int              // top of stack
		ip    int              // id as element of stack
	}
)

var vm VM

func initVM(code Code) {
	vm.top = 0
	vm.ip = 0
	vm.code = code // push syntax to Code array
}

// get code from array vm.code => increament vm.ip
func nextCode() int {
	code := vm.code[vm.ip]
	vm.ip++
	return code
}

func pushValue(val Value) {
	vm.stack[vm.top] = val
	vm.top++
}

func popValue() Value {
	vm.top--
	val := vm.stack[vm.top]
	return val
}

func runVM() int {
	for vm.ip < len(vm.code) {
		op := nextCode()
		debugVM(op)

		switch op {
		case HLT:
			return VM_OK
		case PSH:
			val := Value(nextCode())
			pushValue(val)
			break
		case POP:
			popValue()
			break
		case PUT:
			val := popValue()
			fmt.Printf("-> %g\n", val)
			break
		case ADD:
			b, a := popValue(), popValue()
			pushValue(a + b)
			break
		case SUB:
			b, a := popValue(), popValue()
			pushValue(a - b)
			break

		case MUL:
			b, a := popValue(), popValue()
			pushValue(a * b)
			break

		case DIV:
			b, a := popValue(), popValue()
			pushValue(a / b)
			break
		default:
			fmt.Printf("unknow opcode: %d\n", op)
			return VM_ERROR
		}
	}

	return VM_ERROR
}

func debugVM(op int) {
	switch op {

	case HLT:
		fmt.Print(">> on halt program!\n")
		break

	case PSH:
		val := vm.code[vm.ip]
		fmt.Printf(">> push %d to stack\n", val)
		break

	case POP:
		val := vm.stack[vm.top-1]
		fmt.Printf(">> pop %g from stack\n", val)
		break

	case PUT:
		val := vm.stack[vm.top-1]
		fmt.Printf(">> pop %g from stack and print it\n", val)
		break

	case ADD:
		b, a := vm.stack[vm.top-1], vm.stack[vm.top-2]
		fmt.Printf(">> add %g to %g, push %g\n", a, b, a+b)
		break

	case SUB:
		b, a := vm.stack[vm.top-1], vm.stack[vm.top-2]
		fmt.Printf(">> sub %g to %g, push %g\n", a, b, a-b)
		break

	case MUL:
		b, a := vm.stack[vm.top-1], vm.stack[vm.top-2]
		fmt.Printf(">> mul %g to %g, push %g\n", a, b, a*b)
		break

	case DIV:
		b, a := vm.stack[vm.top-1], vm.stack[vm.top-2]
		fmt.Printf(">> div %g to %g, push %g\n", a, b, a/b)
		break
	}
}

func main() {
	// 4 + 5
	code1 := []int{PSH, 4, PSH, 5, ADD, PUT, HLT}
	// 100 - 52 + 88 // 35 * 8 / 7
	code2 := []int{PSH, 100, PSH, 52, SUB, PSH, 88, ADD, PUT, PSH, 35, PSH, 8, MUL, PSH, 6, DIV, PUT, HLT}

	initVM(code1)
	if runVM() == VM_OK {
		fmt.Print("code1: OK\n\n")
	} else {
		fmt.Print("code1: ERROR\n\n")
	}

	initVM(code2)
	if runVM() == VM_OK {
		fmt.Print("code2: OK\n\n")
	} else {
		fmt.Print("code2: ERROR\n\n")
	}
}
