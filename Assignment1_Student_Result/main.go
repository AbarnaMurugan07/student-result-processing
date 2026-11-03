package main

import "fmt"

func main() {
	var s Student

	fmt.Println("===== STUDENT RESULT PROCESSING SYSTEM =====")

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&s.Name)

	fmt.Print("Enter Register Number: ")
	fmt.Scan(&s.RegisterNo)

	subjects := []string{"Tamil", "English", "Computer", "Maths", "Science"}

	s.Marks = make([]int, len(subjects))

	for i := 0; i < len(subjects); i++ {
		fmt.Print("Enter ", subjects[i], " Mark: ")
		fmt.Scan(&s.Marks[i])

		if s.Marks[i] < 0 || s.Marks[i] > 100 {
			fmt.Println("Invalid mark! Enter marks between 0 and 100.")
			return
		}
	}

	total := calculateTotal(s.Marks)
	average := calculateAverage(total, len(s.Marks))
	grade := calculateGrade(average)

	result := "PASS"

	for _, mark := range s.Marks {
		if mark < 35 {
			result = "FAIL"
			break
		}
	}

	fmt.Println("\n========== STUDENT RESULT ==========")
	fmt.Println("Student Name :", s.Name)
	fmt.Println("Register No  :", s.RegisterNo)

	for i, mark := range s.Marks {
		fmt.Println(subjects[i], ":", mark)
	}

	fmt.Println("Total        :", total)
	fmt.Printf("Average      : %.2f\n", average)
	fmt.Println("Grade        :", grade)
	fmt.Println("Result       :", result)
}
