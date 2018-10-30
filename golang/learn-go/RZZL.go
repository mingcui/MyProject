// this program just for rzzl deploy
// go version 1.10.3

package main

import (
 "flag"
 "fmt"
 "io/ioutil"
 "os"
 "os/exec"
 "strings"
)

// set gloab var
var usageText = `
This Promgram should follow below order to run:
1. user Ri to delete old version images
2. user R to run

Usage RZZL_Dep [OPTION]

       -S,  --S      Service Name
       -Ri, --Ri     Remove Image
       -R, --R       Run Service

example: 

 RZZL_Dep -R -S ServiceName    # run service

 RZZL_Dep -Ri -S ServiceName    #remove image`

func RmImage(SerName string) {
 //  DELETE docker
 ImageName := ""
 // exec.Command("docker", "stop", SerName)
 comandsDelDocker := "/usr/bin/docker rm -f " + SerName
 cmdDelDocker := exec.Command("/bin/bash", "-c", comandsDelDocker)
 if _, err := cmdDelDocker.Output(); err != nil {
  fmt.Println("Error occur when run command delete docker")
  fmt.Println(err)
  os.Exit(2)
 } else {
  fmt.Printf("Del Docker for %s is Successufl !", SerName)
 }

 //  find image
 comandsFindImg := "/usr/bin/docker images | grep " + SerName + "| awk -F ' ' '{print $(NF-5)}'"
 fmt.Println(comandsFindImg)
 cmdFindImg := exec.Command("/bin/bash", "-c", comandsFindImg)
 if Img, err := cmdFindImg.Output(); err != nil {
  fmt.Println("Error occur when run command find image:")
  fmt.Println(err)
  os.Exit(1)
 } else {
  ImageName = string(Img)
  fmt.Print(ImageName)
 }

 //   delete Image
 comandsDelImg := "/usr/bin/docker rmi " + ImageName
 fmt.Print(comandsDelImg)
 cmdDelImg := exec.Command("/bin/bash", "-c", comandsDelImg)
 if _, err := cmdDelImg.Output(); err != nil {
  fmt.Println("Error occur when run command delete image")
  fmt.Println(err)
  os.Exit(2)
 } else {
  fmt.Printf("Del Iamge for %s is Successufl !\n", SerName)
  os.Exit(0)
 }
}

func RunService(SerName string) {
 // start image with given serivce
 data, err := ioutil.ReadFile("dockerRunCommand.txt")
 if err != nil {
  fmt.Println(err)
  os.Exit(3)
 }
 // fmt.Println(string(data))
 datas := strings.Split(string(data), "\n")
 for _, comandsRunDocker := range datas {
  fmt.Println(comandsRunDocker)
  if strings.Contains(comandsRunDocker, SerName) {
   fmt.Printf("The comands: [ %s ] contian service: %s \n", comandsRunDocker, SerName)
   fmt.Printf("Now run the command! \n")
   cmdRun := exec.Command("/bin/bash", "-c", comandsRunDocker)
   if _, err := cmdRun.Output(); err != nil {
    fmt.Println("Error occur when run docker")
    fmt.Println(err)
    os.Exit(4)
   } else {
    fmt.Printf("Run Docker for %s is Successufl !\n", SerName)
    break
   }
  } else {
   fmt.Println("Can't find this service in command list \n")
  }
 }
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
 } else {
  fmt.Println(usageText)
 }
}
