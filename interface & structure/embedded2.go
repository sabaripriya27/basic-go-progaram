// Go program to illustrate the concept
// of the embedding interfaces
package main

import "fmt"

// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
	picked()
}

// Interface 3
// Interface 3 embedded with interface
// 1's method and interface 2
// And also contain its own method
type FinalDetails interface {
	details()
	AuthorArticles
	cdeatils()
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
	cid       int
	post      string
	pick      int
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

// Implementing methods
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

func (a author) picked() {

	fmt.Printf("\nTotal number of picked articles: %d", a.pick)
}

// Implementing the method
// of the embedded interface
func (a author) cdeatils() {

	fmt.Printf("\nAuthor Id: %d", a.cid)
	fmt.Printf("\nPost: %s", a.post)
}

// Main value
func main() {

	// Assigning values to the structure
	values := author{

		a_name:    "abc",
		branch:    "ECE",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
		cid:       3087,
		post:      "asdf",
		pick:      58,
	}

	// Accessing the methods
	// of the interface 1 and 2
	// Using FinalDetails interface
	var f FinalDetails = values
	f.details()
	f.articles()
	f.picked()
	f.cdeatils()
}
