// this program just for rzzl deploy
// go version 1.10.3

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// set gloab var
var usageText = `

Usage flagExp2 [OPTION]

       -S,  --S      Service Name
       -Ri, --Ri     Remove Image
       -R, --R       Run Service

example: 

	flagExp2 -R -S ServiceName    # run service

	flagExp2 -Ri -S ServiceName    #remove image`

func RmImage(SerName string) (ImageName string) {
	//  delete images
	comands := "/usr/bin/docker images | grep " + SerName + "| awk -F ' ' '{print $(NF-5)}'"
	cmd := exec.Command("/bin/bash", "-c", comands)
	if ImageName, err := cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(ImageName))
	}

	return ImageName

}

func RunService(SerName string) {
	// start image with given serivce
}

func main() {
	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, " %s\n", usageText)

	}
	ServiceName := flag.String("S", "nil", "Service name")

	R := flag.Bool("R", false, "Start service")

	Ri := flag.Bool("Ri", false, "Remove Image")

	flag.Parse()

	fmt.Println("value of ServiceName:", *ServiceName)

	fmt.Println("value of R:", *R)

	fmt.Println("value of Ri:", *Ri)

	if *ServiceName == "" {
		fmt.Println("No Service Name Give!")
		fmt.Println(usageText)
		os.Exit(1)
	}
	/*
		        length := len(os.Args)
				arrys := os.Args[1:]
				ServiceName = arrys[-1]
				println(ServiceName)


		for _, argu := range arrys {
				switch argu {
				case "-Ri":
					RmImage(*ServiceName)
				case "-R":
					RunService(*ServiceName)
				default:
					fmt.Println(usageText)
					os.Exit(2)
				}
			}
	*/
	if (*Ri) && (*R == false) {
		fmt.Println("value of Ri:", *Ri)
		RmImage(*ServiceName)
	} else if (*R) && (*Ri == false) {
		fmt.Println("value of Ri:", *Ri)
		RunService(*ServiceName)
	} else if *R && *Ri {
		fmt.Println("value of R:", *R)
		fmt.Println("value of Ri:", *Ri)
		RmImage(*ServiceName)
		RunService(*ServiceName)
	}
}
