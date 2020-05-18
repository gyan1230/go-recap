package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.OpenFile("/tmp/test", os.O_RDWR, 0644)
	checkErr(err)
	defer f.Close()

	n, err := f.WriteString("Hello world!!!!\n")
	checkErr(err)
	fmt.Println("no of bytes written : ", n)
	someText := "How are you?\n"
	n, err = f.Write([]byte(someText))
	checkErr(err)

	fmt.Println("no of bytes written : ", n)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// b := make([]byte, 100)
	// n3, err := f.Read(b)
	// checkErr(err)

	b, err := ioutil.ReadFile("/tmp/test")
	checkErr(err)
	stringversion := string(b)
	fmt.Println("The content of file: ", stringversion)
}

func checkErr(e error) {
	if e != nil {
		fmt.Println("error: ", e)
		os.Exit(1)
	}
}
