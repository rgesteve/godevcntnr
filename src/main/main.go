package main

import (
	"fmt"
	//"runtime"
	"encoding/csv"
	"log"
	//"strings"
	"io"
	"embed"
	"fakedep"
)

//go:embed data.csv
var fs embed.FS

func main() {
/*
     in := `first_name, last_name, username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
*/
     fmt.Println("Hello, World from main")
     fakedep.PrintHelloWorld()
     // fmt.Printf("Go version is: %s\n", runtime.Version())

     // Open the embedded CSV file
     file, err := fs.Open("data.csv")
     if err != nil {
     	log.Fatal("Error opening embedded file:", err)
	return
     }

     // Create a new CSV reader from the opened file
     r := csv.NewReader(file)

//     r := csv.NewReader(strings.NewReader(in))

     for {
     	 record, err := r.Read()
	 if err == io.EOF {
	    break
	 }
	 if err != nil {
	    log.Fatal(err)
	 }
	 fmt.Println(record)
     }

     fmt.Println("done!")
}
