package main

import "fmt"

// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Structure
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// Implementing method
// of the interface 1
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)

}

// Implementing method
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

// Main value
func main() {

	// Assigning values
	// to the structure
	fmt.Println("first person details")
	values := author{
		a_name:    "sabari",
		branch:    "ECE",
		college:   "XYZ",
		year:      2020,
		salary:    18000,
		particles: 100,
		tarticles: 139,
	}

	// Accessing the method
	// of the interface 1
	var i1 AuthorDetails = values
	i1.details()

	// Accessing the method
	// of the interface 2
	var i2 AuthorArticles = values
	i2.articles()

}
