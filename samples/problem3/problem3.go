package main

import (
	"fmt"
       "bufio"
       "log"
       "io/ioutil"
       "strings"
       "os"
       "flag"
       "net/http"
       "strconv"
)


// report the error to the user.
func check (e error) {
     if e != nil {
     	log.Fatal(e)
     }
}

// process a URL and store the frequency table in a file 
func processUrl(url string, fileNo string, done chan bool){
        // Download the page from the given url
	
    	resp, err_resp := http.Get(url)
	check(err_resp)
	defer resp.Body.Close()
	body, err_body := ioutil.ReadAll(resp.Body)
	check(err_body)

	//initialization
	word_count := map[string]int{}
	pos := 0

	// parse the body and extract each word. A word is defined as A-Za-z0-9. Store each word in a hashmap as a (word,count) key value pair 
	for index,element := range body {
		if !((element >= 65 && element <= 90)||(element >= 97 && element <= 122)||(element >= 48 && element <= 57)) {
		   	if pos < index {
			   word := string(body[pos:index])
			   if val, ok := word_count[word];ok{
			   		word_count[word] = val + 1
			   } else {
			   	word_count[word] = 1
			   }
			}
			pos = index+1;
		}
		
	}

	// create a file and print the url in the first line, followed by each word in a new line along with the associated count
	fOut, err_file := os.Create("url"+fileNo+".txt")
 	check(err_file)
	defer fOut.Close()
	fOut.WriteString("url:	"+url)
	for key, value := range word_count {
		fOut.WriteString("\n\t"+key+":\t"+strconv.Itoa(value))
	}
	done <- true
}

func main() {
     	
	inputPtr := flag.String("urls","","comma-seperated-one-or-more-urls")
    	flag.Parse()
    	var urls string = *inputPtr
    	if urls == ""{
       	   fmt.Println("Usage: -urls=<comma-seperated-one-or-more-urls>")
      	   os.Exit(0)
   	}
	// Note: Be careful with using scanners, they have a restriction on token size. 
	// But here the url size is not expected to be too large.
	scanner := bufio.NewScanner(strings.NewReader(urls))

	 // function to split the file on ','
    	onComma :=func(data []byte, atEOF bool) (advance int, token []byte, err error) {
    	      for i:=0; i< len (data); i++ {
		  if data[i] == ',' {
		     return i+1, data[:i], nil
		  } 
	      }
	      return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	i := 0
	done := make (chan bool)
	 for scanner.Scan() {
	   i ++
    	   go processUrl(scanner.Text(), strconv.Itoa(i), done) //Executing the for-loop in parallel. Passing the file-no which will be used 
	      				 		        //to create a file in which the frequency table is to be stored. 
								//Also passing the channel used to signal when the goroutine has finished executing.
	}
	for j := 1; j<=i; j++ { // wait for all goroutines to complete before exiting 
	    <-done
	} 
}
