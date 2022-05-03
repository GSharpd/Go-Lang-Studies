package main

import "fmt"

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	case 7:
		return "sunday"
	default:
		return "invalid"
	}
}

func dayOfTheWeek2(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Sunday"
	default:
		weekDay = "Whatever"
	}

	return weekDay
}

func dayOfTheWeek3(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Sunday"
		fallthrough
	default:
		weekDay = "Whatever"
	}

	return weekDay
}

func main() {
	fmt.Println("Switch statements")

	day := dayOfTheWeek(4)

	fmt.Println(day)

	day2 := dayOfTheWeek(10)

	fmt.Println(day2)

	day3 := dayOfTheWeek2(1)
	day4 := dayOfTheWeek2(2)

	fmt.Println(day3, day4)

	day5 := dayOfTheWeek3(1)
	fmt.Println(day5) // results in "whatever", since falltrhough clause makes code go into the next statement of the switch

	// go doesn't demand the break clause
}
