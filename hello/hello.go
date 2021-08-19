//go install
//~/go/bin/hello

package main

import (
	"fmt"
	"log"
	"os"
)

func ar2str(arr []byte, countt int) string {
	var str string = ""
	for i := 0; i < countt; i++ {
		//fmt.Printf("%q", count, data[:count])
		//byte2 code point rune
		char := rune(arr[i])
		fmt.Println("rune q=%q", char)
		fmt.Println("rune q=%v", char)
		//str = str +
	}

	return str
}

func main() {
	//fmt.Println("Hello, world 2.")

	//https://pkg.go.dev/os
	//file, err := os.Open("file.go") // For read access.
	file, err := os.Open("file.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	////////////////////
	//fmt.Printf("read %d bytes: %q\n", count, data[:count])
	fmt.Printf("%q", data[:count])

	//fmt.Printf("%q\n", data.Split("a,b,c", ","))

}
