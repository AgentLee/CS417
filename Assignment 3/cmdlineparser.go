// used to figure out how to parse command line input
// not used for project itself

package main

import "flag"
import "fmt"

func blah() {

	urlPtr		:=	flag.String("url", "", "input url")
    methodPtr 	:= 	flag.String("method", "", "create, remove, etc.")
    yearPtr 	:= 	flag.String("year", "", "remove students with this year")
    dataPtr		:= 	flag.String("data", "", "JSON data for new students")

    flag.Parse()

    fmt.Println("url:", *urlPtr)
    fmt.Println("method:", *methodPtr)
    fmt.Println("year:", *yearPtr)
    fmt.Println("data:", *dataPtr)
}
