package main

import "fmt"

func main() {
	var username string = "ratnesh"

	fmt.Print(username)
	fmt.Printf("Variable is of type:  %T \n ", username)

	var isloggedin bool = true
	fmt.Print(isloggedin)
	fmt.Printf("Variable is of type:  %T \n ", isloggedin)

	var smallvalue int = 255
	fmt.Print(smallvalue)
	fmt.Printf("Variable is of type:  %T \n ", smallvalue)

	var anothervariable int 
	fmt.Print(anothervariable)
	fmt.Printf("Variable is of type:  %T \n ", anothervariable)


	withoutvar := "ratnesh"
	fmt.Print(withoutvar)
}
