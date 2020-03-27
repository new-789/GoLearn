package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b string
	var c float32
	//fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	fmt.Fscan(os.Stdin, &a, &b, &c)
	//fmt.Fscanln(os.Stdin, &a, &b, &c)
	//fmt.Fprintf(os.Stdout, "a=%d b=%s c=%f\n" ,a, b, c)
	//fmt.Fprintln(os.Stdout, a, b, c)
	fmt.Fprint(os.Stdout, a, b, c)
}
