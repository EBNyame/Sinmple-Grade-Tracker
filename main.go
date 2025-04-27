package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/bytedance/sonic/encoder"
)

// Add student grades (integers).
// View all grades.
// Calculate and print the average grade.
// Exit the program.

type StudentGrade struct{
	Name	string
	Grade	int
}

func saveGrades(filename string, grades []StudentGrade){
	file, err := os.Create(filename)
	if err != nil{
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(grades)
	if err != nil{
		fmt.Println("Error saving data:", err)
	}else{
		fmt.Println("Grades saved to file.")
	}
}

func loadGrades(filename string)

func main() {
	var grades []StudentGrade

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
			var name string
			fmt.Println("Enter your name: ")
			fmt.Scan(&name)
			var grade int
			fmt.Println("Enter grade: ")
			fmt.Scan(&grade)

			student := StudentGrade{Name: name, Grade: grade}
			grades = append(grades, student)
			fmt.Println("Grade added for", name)

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
