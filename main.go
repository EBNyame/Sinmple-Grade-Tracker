package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

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

func loadGrades(filename string) ([]StudentGrade, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var grades []StudentGrade
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&grades)
    if err != nil {
        return nil, err
    }
    return grades, nil
}

func main() {
	var grades []StudentGrade
	filename := "grades.json"

	grades, err := loadGrades(filename)
	if err != nil{
		fmt.Println("No previous grades found, starting fresh.")
	}

	for {
		fmt.Println("\n--- Grade Tracker ---")
		fmt.Println("1. Add grades")
		fmt.Println("2. View all grades")
		fmt.Println("3. Calculate average")
		fmt.Println("4. Sort Grades")
		fmt.Println("4. Save and Exit")
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
				sum += g.Grade
			}
			average := float64(sum) / float64(len(grades))
			fmt.Println("Average: ", average)

		case 4:
            fmt.Println("Sort by:")
            fmt.Println("1. Grade")
            fmt.Println("2. Name")
            fmt.Print("Choose an option: ")
            var sortChoice int
            fmt.Scan(&sortChoice)

            if sortChoice == 1 {
                sort.Slice(grades, func(i, j int) bool {
                    return grades[i].Grade < grades[j].Grade
                })
                fmt.Println("Grades sorted by grade.")
            } else if sortChoice == 2 {
                sort.Slice(grades, func(i, j int) bool {
                    return grades[i].Name < grades[j].Name
                })
                fmt.Println("Grades sorted by name.")
            } else {
                fmt.Println("Invalid option.")
            }

        case 5:
            saveGrades(filename, grades)
            fmt.Println("Goodbye!")
            return

        default:
            fmt.Println("Invalid option.")
        }
    }
}
