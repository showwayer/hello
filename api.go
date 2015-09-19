package main

import "fmt"

import "time"
import "github.com/showwayer/stringutil"
import "os"

func main() {
	a := 1
	b := 3
	ow := "Google"
	fmt.Println("Starting API Server...")
	fmt.Println("Creating temp file...")
	myFile, err := os.Create("tmp.txt")
	if err != nil {
		fmt.Println("Fail to create file!")
		os.Exit(-1)
	}
	myFile.WriteString("sample text goes into this tmp file")
	err = myFile.Close()
	//here we call a function in stringutil
	fmt.Printf("Reversing %s...%s...\n", ow, stringutil.Reverse(ow))
	fmt.Printf("%d + %d = %d\n", a, b, stringutil.Add(a, b))

	myChannel := make(chan bool, 1)
	go worker(myChannel)

	//read and block the amin thread
	<-myChannel

	fmt.Println("The end")
}

func worker(myChannel chan bool) {
	fmt.Printf("working on something...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//return to main thread by setting a boolean
	myChannel <- true

}
