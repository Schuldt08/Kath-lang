package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
    fmt.Println("usage: kath [file]")
}

func main() {
    args := os.Args
    if len(args) != 2 {
        usage()
        return
    }
    f, err := os.ReadFile(args[1])
    if err != nil {
        log.Fatal("INVALID FILE PATH")
    }
    code := string(f)
    run(tokenize(code))
}

