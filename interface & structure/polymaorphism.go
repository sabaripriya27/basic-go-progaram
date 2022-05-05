// Go program to illustrate the concept
// of polymorphism using interfaces
package main

import "fmt"

// Interface
type student interface {
	mark() int
	paper() string
}

// Structure 1
type team1 struct {
	mark1   int
	paper_1 string
}

// Methods of employee interface are
// implemented by the team1 structure
func (t1 team1) mark() int {
	return t1.mark1
}

func (t1 team1) paper() string {
	return t1.paper_1
}

// Structure 2
type team2 struct {
	mark_2  int
	paper_2 string
}

// Methods of employee interface are
// implemented by the team2 structure
func (t2 team2) mark() int {
	return t2.mark_2
}

func (t2 team2) paper() string {
	return t2.paper_2
}

func finaldevelop(i []student) {

	mark := 0
	var total int
	for _, stu := range i {

		fmt.Printf("\nstudent name = %s\n ", stu.paper())
		fmt.Printf("Total number of mark %d\n ", stu.mark())
		mark += stu.mark()
		total = mark / 2
	}
	fmt.Printf("\n total Tamil paper mark = %d", total)
}

// Main function
func main() {

	res1 := team1{mark1: 80,
		paper_1: "Tamil paper-I"}

	res2 := team2{mark_2: 76,
		paper_2: "Tamil paper-II"}

	final := []student{res1, res2}
	finaldevelop(final)

}
