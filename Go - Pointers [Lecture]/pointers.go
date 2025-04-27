package main

import "fmt"

func main() {
	age := 32 //regular variable

	age_pointer := &age // One example of pointer declaration

	// Another method for pointer declaration
	// var age_pointer *int
	// age_pointer = &age

	fmt.Println("Age", *age_pointer) // Example of dereferencing. That is the value behind the pointer.
	//adultYears := getAdultYears(age_pointer)
	// getAdultYears(age_pointer)
	editAgeToAdultYears(age_pointer)
	// fmt.Println("The age is:", adultYears)
	fmt.Println("The age is: ", age)
}

// func getAdultYears(age *int) int {
func editAgeToAdultYears(age *int) {
	// return (*age - 18)
	*age = *age - 18

}
