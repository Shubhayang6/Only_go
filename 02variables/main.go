package main

import "fmt"

const LoginToken string = "aabbccdd"

func main() {
	var username string = "Shubhayan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.1343434343435
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	/* deafault values and same aliases */
	var aliases int
	fmt.Println(aliases)
	fmt.Printf("Variable is of type: %T \n", aliases)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)


}
