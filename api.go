package main

import "fmt"

import "time"
import "github.com/showwayer/stringutil"
import "os"

func main() {
	a := 1
	b := 3
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
	fmt.Printf(stringutil.Reverse("!ENOD") + "\n")
	fmt.Printf("%d + %d = %d", a, b, stringutil.Add(a, b))
	fmt.Println("The end")
}

func worker(myChannel chan bool) {
	fmt.Printf("working on something...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//return to main thread by setting a boolean
	myChannel <- true

}
