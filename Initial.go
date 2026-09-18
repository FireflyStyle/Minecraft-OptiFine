package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Classroom struct {
	Students []Student
}

func (c *Classroom) AddStudent(name string, age int, score float64) {
	c.Students = append(c.Students, Student{
		Name:  name,
		Age:   age,
		Score: score,
	})
}

func (c *Classroom) SortByScore() {
	sort.Slice(c.Students, func(i, j int) bool {
		return c.Students[i].Score > c.Students[j].Score
	})
}

func (c *Classroom) AverageScore() float64 {
	if len(c.Students) == 0 {
		return 0
	}

	total := 0.0

	for _, student := range c.Students {
		total += student.Score
	}

	return total / float64(len(c.Students))
}

func (c *Classroom) PrintReport() {
	fmt.Println("Classroom Report")
	fmt.Println("================")

	for index, student := range c.Students {
		fmt.Printf(
			"%d. %s | Age: %d | Score: %.1f\n",
			index+1,
			student.Name,
			student.Age,
			student.Score,
		)
	}

	fmt.Println("================")
	fmt.Printf("Students: %d\n", len(c.Students))
	fmt.Printf("Average Score: %.2f\n", c.AverageScore())

	if len(c.Students) > 0 {
		fmt.Printf("Top Student: %s\n", c.Students[0].Name)
		fmt.Printf("Top Score: %.1f\n", c.Students[0].Score)
	}
}

func main() {
	classroom := Classroom{}

	classroom.AddStudent("Alice", 20, 94.5)
	classroom.AddStudent("Brian", 21, 87.0)
	classroom.AddStudent("Clara", 19, 98.0)
	classroom.AddStudent("David", 22, 91.5)

	classroom.SortByScore()
	classroom.PrintReport()
}