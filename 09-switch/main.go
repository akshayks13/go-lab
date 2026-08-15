package main

import (
	"fmt"
)

func main(){
	var day string 
	fmt.Print("Enter the day of the week (e.g., Monday, Tuesday, etc.): ")
	fmt.Scan(&day)

	if day == ""{
		fmt.Println("No day entered, please enter a valid day of the week.")
		return
	}

	switch day {
	case "Monday":
		fmt.Println("It's Monday, start of the week!")
	case "Tuesday":
		fmt.Println("It's Tuesday, keep going!")
	case "Wednesday":
		fmt.Println("It's Wednesday, halfway through!")
	case "Thursday":
		fmt.Println("It's Thursday, almost there!")
	case "Friday":
		fmt.Println("It's Friday, the weekend is near!")
	case "Saturday":
		fmt.Println("It's Saturday, enjoy your day!")
	case "Sunday":
		fmt.Println("It's Sunday, relax and recharge!")
	default:
		fmt.Println("Invalid day entered, please enter a valid day of the week.")
	}
}