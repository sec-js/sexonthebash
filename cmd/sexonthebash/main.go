package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//go routine which read stdin permently (and print it for now);
// Maybe use mutex, safe read to avoid race condition
// Once it is ok the same thing for stdout

// see:
//https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/
//StdinPipe for exec command
//https://zetcode.com/golang/pipe/
//https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
//https://stackoverflow.com/questions/50788805/how-to-read-from-stdin-with-goroutines-in-golang

func main() {
	fmt.Println("Before bash")

	bash := exec.Command("/bin/bash", "-s")
	var outBuffer bytes.Buffer
	// var errBuffer bytes.Buffer
	var inBuffer bytes.Buffer

	mwOut := io.MultiWriter(os.Stdout, &outBuffer)
	io.TeeReader(os.Stdin, &inBuffer)
	// mwErr := io.MultiWriter(os.Stderr, &errBuffer)
	// mrIn := io.MultiReader(os.Stdin, &inBuffer)

	bash.Stdout = mwOut
	bash.Stderr = os.Stderr
	bash.Stdin = os.Stdin

	bash.Run()

	fmt.Println()
	fmt.Println("After bash")
	fmt.Println("Captured output", outBuffer.String())
	// fmt.Println("Captured error", errBuffer.String())
	fmt.Println("Captured input", inBuffer.String())
}
