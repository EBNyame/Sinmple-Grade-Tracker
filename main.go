package main

import "fmt"

// Add student grades (integers).
// View all grades.
// Calculate and print the average grade.
// Exit the program.

func main() {
	var grades []int

	for {
		fmt.Println("1. Add grades")
		fmt.Println("2. View all grades")
		fmt.Println("3. Calculate average")
		fmt.Println("4. Exit")
		fmt.Println("Choose option")

		var choice int
		fmt.Scan(&choice)

		switch choice{
		case 1:
			//Add student grades (integers).
			var grade int
			fmt.Scan(&grade)
			fmt.Println("Enter grade: ")
			grades = append(grades, grade)
			fmt.Println("Grade recorded...")

		case 2:
			// View all grades.
			fmt.Println("Grades: ", grades)

		case 3:
			//Calculate average
			if len(grades) == 0{
				fmt.Println("No grades to average")
				continue
			}
				//calculate for the total
			sum := 0
			for _, g := range grades{
				sum += g
			}
			average := float64(sum) / float64(len(grades))
			fmt.Println("Average: ", average)

		case 4:
			fmt.Println("See you again.... Goodbye!")

		default:
			fmt.Println("Invalid input... Choose again")
		}
	}

}
