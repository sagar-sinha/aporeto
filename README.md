

What does it contain?
---------------------

There are 3 programs.
- Program 1 is a bash script. It's a simple user-driven program which creates a file with the user supplied filename containing the names of the fifty US states, each in a new line. 
> Usage - ./program1 [--help|-h]
        ./program1 --create-file=<filename> [--no-prompt] [--verbose]
- Program 2 is coded in Golang. It creates an output file with duplicate lines removed from the input file . 
> Usage - program2 [--help|-h]
        program2 --file=<filename> --output=<output-filename> [—verbose]
        
- Program 3 is coded in Golang too(I think I love it). It creates an output file contaning the word-frequency table of the different words in the file who's URI the user specifies.
> Usage - program3 [-help|-h]
        program3 -urls=<comma-seperated-one-or-more-urls>

Instructions to download and run the GO problems 2 and 3 on a Unix shell
------------------------------------------------------------------------

+ go get github.com/sagar-sinha/aporeto/samples/problemX                                {replace X by 2 or 3} 
+ $GOPATH/bin/problemX                                                                  {The command above downloads                                                                                               the binary into your bin Go                                                                                                workspace's bin folder, find and                                                                                           run it }


My insights into the problems 
-----------------------------
- Problem 3 processes (downloads & creates the frequency table) the URLs in parallel using gosubroutines and a channel (which signals the main when all the subroutines have finished and it can safely exit).

- Problem 2 uses a hashmap to store the lines. The program has a complexity of O(n) but one needs to be careful with the memory usage. The files can be long and storing a hashmap in memory may not be the best idea. Would rather then change the program to store the hashmap on the disk and have only a cache in memory using the LRU scheme to handle missing entries.

- Problem 1 first removes an alreading existing file before creating a new one. Instead one could have directly overwritten the existing file. Doing this so that the verbose statement "File removed" is actually implemented. Ideally would have gotten rid of first deleting and then creating
