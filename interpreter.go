package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func find_proc(proc_index int64, program program) (uint64, error) {
    var line uint64 = 0
    for line < uint64(len(program)) {
        if program[line].opcode == PROC && program[line].value == proc_index {
            return line, nil
        }
        line++
    }
    return 0, errors.New("attempt to evaluate nonexistent procedure")
}

func run(input program) {
    // idk why i chose 4096 specifically honestly, the length of this array doesnt really matter
    memory := [4096]int64{}
    var ptr uint16
    call_stack := []uint64{}
    

    line, err := find_proc(0, input)
    if err != nil {
        log.Fatal("unable to find entrypoint proc 0")
    }
    line ++

    for line < uint64(len(input)){
        statement := input[line]
        switch statement.opcode {
        // pointer movement
        case PNT:
            if statement.value < 0 || statement.value > 65535 {
                log.Fatal("ERROR: attempt to point to invalid memory position")
            } 
            
            ptr = uint16(statement.value)
        case SHIFT:
            new_ptr := int64(ptr) + statement.value
            if new_ptr < 0 || new_ptr > 65535 {
                log.Fatal("ERROR: attempt to shift pointer to invalid memory position")
            }
            
            ptr = uint16(new_ptr)
        
        // math
        case SET:
            memory[ptr] = statement.value
        case COPY:
            if statement.value < 0 || statement.value > 65535 {
                log.Fatal("ERROR: attempt to copy from invalid memory position")
            }
            
            memory[ptr] = memory[statement.value]
        case ADD:
            if statement.value < 0 || statement.value > 65535 {
                log.Fatal("ERROR: attempt to read from invalid memory position")
            }
            memory[ptr] += memory[statement.value]
        case ADDN:
            memory[ptr] += statement.value
        case SUB:
            if statement.value < 0 || statement.value > 65535 {
                log.Fatal("ERROR: attempt to read from invalid memory position")
            }
            memory[ptr] -= memory[statement.value]
        case SUBN:
            memory[ptr] -= statement.value

        // basic I/O
        case IN:
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            err := scanner.Err()
            if err != nil {
                log.Fatal(err)
            } 
            
            num, err := strconv.ParseInt(scanner.Text(), 10, 64)
            if err != nil {
                log.Fatal("ERROR: input is not a valid integer")
            }
            memory[ptr] = num
        case OUT:
            fmt.Println(memory[ptr])

        // loops
        case LOOP:
            if memory[ptr] == statement.value {
                count := 1
                for count > 0 && line < uint64(len(input)) {
                    line++
                    if input[line].opcode == LOOP {count++}
                    if input[line].opcode == UNLOOP {count--}
                }
            }
        case UNLOOP:
            if memory[ptr] != statement.value {
                count := 1
                for count > 0  && line < uint64(len(input)) {
                    line-- 
                    if input[line].opcode == LOOP {count--}
                    if input[line].opcode == UNLOOP {count++}
                }
            }

        // procedures
        case PROC:
            count := 1
            for count > 0 && line < uint64(len(input)) {
                line++
                if input[line].opcode == PROC {count++}
                if input[line].opcode == END {count--}
            }
        case EVAL:
            proc_line, err := find_proc(statement.value, input)
            if err != nil {
                log.Fatal(err)
            }
            call_stack = append(call_stack, line)
            line = proc_line
        case END:
            if len(call_stack) > 0 {
                stack_line := call_stack[len(call_stack) - 1]
                call_stack = call_stack[:len(call_stack) - 1]
                line = stack_line
            } else {
                return
            }
        }

    line++
    }
}
