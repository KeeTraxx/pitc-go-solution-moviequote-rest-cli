package main

import "net/http"
import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"

// MovieQuote represents a quote from a movie
type MovieQuote struct {
	Movie     string `json:"movie"`
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Actor     string `json:"actor"`
	Year      uint   `json:"year"`
}

func main() {

	// Make a http GET Request to http://pitc-moviequote.ose3-lab.puzzle.ch/v1/moviequotes/random
	resp, err := http.Get("http://pitc-moviequote.ose3-lab.puzzle.ch/v1/moviequotes/random")

	// Handle errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read HTTP response body
	body, err := ioutil.ReadAll(resp.Body)

	// Handle errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Declare a varable
	var quote MovieQuote

	// Convert http body to struct
	err = json.Unmarshal(body, &quote)

	// Handle errors during conversion
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out the result
	fmt.Printf("\"%v\"\n", quote.Quote)
	fmt.Printf("\tCharacter: %30v\n\tActor:     %30v\n\tMovie:     %23v (%v)\n\n", quote.Character, quote.Actor, quote.Movie, quote.Year)

}
