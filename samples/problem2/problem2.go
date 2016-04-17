package main

import (
       "fmt"
       "log"
       "io/ioutil"
       "os"
       "flag"
)
// check for error
func check (e error) {
     if e != nil {
     	log.Fatal(e)
     }
}

// print only when the verbose flag is set
func print_to_screen(s string, v bool) {
     if (v) {
     	fmt.Println(s)
     }
}
 
func main() {
    inputPtr := flag.String("file","","input file from which duplicates are to be removed")
    outputPtr := flag.String("output","","file to output the content with the duplicate removed")
    verbosePtr := flag.Bool("verbose",false,"prints to screen the messages")
    flag.Parse()
    var file_input string = *inputPtr
    var file_output string = *outputPtr
    if file_input == "" || file_output == ""{
       fmt.Println("Usage: --file=<filename> --output=<output-filename> [—verbose] ")
       os.Exit(0)
    }
    fileArray, read_err := ioutil.ReadFile(file_input)
    check(read_err)
    fOut, create_err := os.Create(file_output)
    check(create_err)
    defer fOut.Close()
    lines_encountered := map[string]bool{} // a hashmap to store lines that have been encountered

    i := 0
    i_temp := 0
    pos := 0
    isLine := false

    // iterate over the array  and add only those lines to the hashmap that have not been encountered before
    // If a line is encountered again don't write to output file
    // The lines are separated by a \n or \r or \r\n. ASCII for carraige return is 13 and for new line is 10
    // Note : One could think of using Scanner with bufio but since the files could be large and scanner has a restriction on token size, it might
    // not be the best idea to use it
    for i < len(fileArray) {
    	   if fileArray[i] == 13 && fileArray[i+1] == 10 {
	      i_temp = i
	      i = i+2
	      isLine = true
	   } else if fileArray[i] == 13 || fileArray[i] == 10 {
	     i_temp = i
	     i = i+1
	     isLine = true 
	   } else {
	     isLine = false
	     i = i + 1
	    }
	   if isLine {
	      if pos < i_temp { // check for 0 length strings which might occur due to successive new line characters
	      	   line := string(fileArray[pos:i_temp])
    	      	   if lines_encountered[line] != true {
	      	 	 lines_encountered[line] = true
	   	 	 fOut.WriteString(line+"\n")
	      	   }
	      }
	      pos = i
	   }
    }
    print_to_screen("Successfully removed duplicates",*verbosePtr)    
    fOut.Sync()

}