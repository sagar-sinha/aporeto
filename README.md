
Instructions to download and run the GO problems 2 and 3
--------------------------------------------------------

- go get github.com/sagar-sinha/aporeto/samples/problemX (replace X by 2 or 3)
- $GOPATH/bin/problemX (The command above downloads the binary into your bin Go workspace's bin folder, find and run it )


Insights into the problems
--------------------------
Problem 3 processes the URLs in parallel using gosubroutines and a channel (which signals the main when all the subroutines have finished and it can safely exit) 

Problem 2 uses a hashmap to store the lines. The program has a complexity of O(n) but one needs to be careful with the memory usage. The files can be long and storing a hashmap in memory may not be the best idea. Would rather then change the program to store the hashmap on the disk and have only a cache in memory using the LRU scheme to handle missing entries.

Problem 1 first removes an alreading existing file before creating a new one. Instead one could have directly overwritten the existing file. Doing this so that the verbose statement "File removed" is actually implemented. Ideally would have gotten rid of first deleting and then creating.
