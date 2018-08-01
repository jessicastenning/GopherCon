package main

import "fmt"

// declare a struct called Movie
// add the following fields:
// - Title (string)
// - Released (bool)
// - Length (int)

type Movie struct {
  title string
  released bool
  length int
}

func main() {
	// declare a variable called "movie" of type "Movie"

    var movie Movie

	// Set the Title to "Wizard of Oz"
	// Set the Released variable to "true"
	// Set Length to 125

    movie.title = "Wizard of Oz"
    movie.released = true
    movie.length = 125

	// Print the value of "movie" out
	// hint: you can use fmt.Println(movie)

  fmt.Println(movie)
}
