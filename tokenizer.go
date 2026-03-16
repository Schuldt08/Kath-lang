package main

import (
	"strconv"
	"strings"
)

type instruction uint64
const (
    NULL instruction = iota
    PNT
    SHIFT
    SET
    COPY
    ADD
    ADDN
    SUB
    SUBN
    IN
    OUT
    LOOP
    UNLOOP
    PROC
    EVAL
    END
)

type codeline struct {
    opcode instruction
    value int64
}

type program []codeline

func tokenize(code string) program {
    output := program{}
    for line := range strings.SplitSeq(code, "\n") {
        trimmed_line := strings.TrimSpace(line)
        args := strings.Split(trimmed_line, " ")
        
        if args[0] == "" {
            continue
        } 
        var operator string
        var value int64
        if len(args) == 2 {
            operator = args[0]
            newval, err := strconv.ParseInt(args[1], 10, 64)
            value = newval
            if err != nil {
                panic("ERROR: arg is not valid integer")
            }
        } else {
            panic("ERROR: invalid amount of args")
        }
        
        var current_inst instruction = 0
        switch operator {
        default: panic("ERROR: invalid opcode")
        case "pnt":
            current_inst = PNT
        case "shift":
            current_inst = SHIFT
        case "set":
            current_inst = SET
        case "copy":
            current_inst = COPY
        case "add":
            current_inst = ADD
        case "addn":
            current_inst = ADDN
        case "sub":
            current_inst = SUB
        case "subn":
            current_inst = SUBN
        case "in":
            current_inst = IN
        case "out":
            current_inst = OUT
        case "loop":
            current_inst = LOOP
        case "unloop":
            current_inst = UNLOOP
        case "proc":
            current_inst = PROC
        case "eval":
            current_inst = EVAL
        case "end":
            current_inst = END
        }

        current_codeline := codeline{current_inst, value}
        output = append(output, current_codeline)
    }

    return output
}
