// flag use examle
package main

import (
	"flag"
	"fmt"
	//"os"
)

func main() {
	S := flag.String("S", "nil", "Service name")
	I := flag.String("I", "nil", "Image name")
	R := flag.Bool("R", false, "Start service")
	Ri := flag.Bool("Ri", false, "Remove Image")
	flag.Parse()
	fmt.Println("value of S:", *S)
	fmt.Println("value of I:", *I)
	fmt.Println("value of R:", *R)
	fmt.Println("value of RI:", *Ri)
}
