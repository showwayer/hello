package main

import "fmt"
import "github.com/showwayer/stringutil"
import "os"

func main() {
	fmt.Printf("Hello, World!\n")
	fmt.Println("Creating temp file...")
	myFile, err := os.Create("tmp.txt")
	if err != nil {
		fmt.Println("Fail to create file!")
	}
	myFile.WriteString("sample text goes into this tmp file")
	myFile.Close()
	//here we call a function in stringutil
	fmt.Printf(stringutil.Reverse("!DOOG") + "\n")
}
