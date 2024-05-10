package main

import (
       "fmt"
       "os"
       //"errors"
       )

type request struct{
     req string
     //TODO error storage
}

var errorUnexpectedInput = fmt.Errorf("Expected 1 argument, got: ")

func input_chk(s []string) ([]string, error) {
    if len(s) != 2 {
       return nil, fmt.Errorf("%w %d", errorUnexpectedInput, len(s) - 1)
    }
    var slice_s []string
    slice_s = append(slice_s, s[1])
    return slice_s, nil
    	   //TODO optimize for copies, pass around an encapsulated struct
}

func main(){
    fmt.Println("Hello world!", os.Args[1])
    /*
	   input_slice := os.Args
           if len(input_slice) != 2 {
               return
           }
           r := request{req: os.Args[1]}
    */
    str, err := input_chk(os.Args) 
    if err != nil {
       fmt.Println(fmt.Errorf("Input error: %w", err))
       return 
    }
    r := request{req: str[0]}
    //fmt.Println(r)
    
//TODO: pretty struct fill
    fmt.Println("Hello world!", r)
}

//TODO1.1 read user input and react to it (filter) /v1 - cmd, v2 stdin(165)
	//Input is a single text string
//TODO1.2 store the string entered in a /v1 - file, /v2 - database

//TODO2.1 display a line from the database
//TODO2.2 open websites entered. Feed to firefox
//TODO2.3 log opened website or error