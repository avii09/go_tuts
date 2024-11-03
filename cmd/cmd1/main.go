package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

// this is a comment
// 1. go to terminal and create module : go mod init <module name>, mod contains list of packages, packages contain list of files.
// 2. go build <file path/name>, this is done to compile the code into an executable file which runs directly on the OS
// 3. go run <file path/name>, the go file is executed.
// 4. every file needs a package?
// when importing main package, it indicates the starting point of the prog, thus while using the main package we need to also specify the main func.
// the main func indicates as a starting point of the prog.
