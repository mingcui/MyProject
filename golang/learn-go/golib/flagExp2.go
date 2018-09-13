// use flag example two
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		usageText := `
Usage flagExp2 [OPTION]
      -S, --S       Service Name
       -I, --I        Image Name
       -Ri, --Ri     Remove Image
       -R, --R       Run Service
example: 
	flagExp2 -S service1 -R     # run service
	flagExp2 -I image1 -Ri    #remove image`
		fmt.Fprintf(os.Stderr, " %s\n", usageText)
	}
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
