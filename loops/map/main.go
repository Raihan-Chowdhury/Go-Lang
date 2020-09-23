package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)
	// Assign kv
	emails["bob"] = "bob@gmail.com"
	emails["ros"] = "ros@gmail.com"
	emails["port"] = "port@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])

	// delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	// def and add kv
	DefMail := map[string]string{"bob": "bob@yahoo.com", "cot": "cot@yahoo.com"}
	fmt.Println(DefMail)
}
