package main

import "fmt"
import "github.com/showwayer/stringutil"
import "os"

func main() {
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
}
