package links

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Creates a new list of links if there isnt one already, runs at startup
func BuildList() {

	basicList := []string{
		"http://google.com\n",
		"http://facebook.com\n",
		"http://twitter.com\n",
		"http://golang.org\n",
		"http://github.com\n",
		"http://amazon.com\n",
		"http://stackoverflow.com\n",
		"http://apple.com\n",
		"http://netflix.com\n",
	}

	//checks for existing file, creates if not found
	if _, err := os.Stat("links.txt"); errors.Is(err, os.ErrNotExist) { //TODO: Run this at app launch
		file, err := os.Create("links.txt")
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		defer file.Close()

		//writes the basicList to the text file
		w := bufio.NewWriter(file)
		for i := range basicList {
			_, err := w.WriteString(basicList[i])
			if err != nil {
				log.Fatal("Error writing to file:", err)
			}
			w.Flush()
		}
	}
}

func GetLinks() []string {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://golang.org",
		"http://github.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://apple.com",
		"http://netflix.com",
	}
	return links
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	func(link string) {
		status := fmt.Sprintf("200 OK - %s", link)
		c <- status
	}(link)
}

func LoopLinks() {
	list := GetLinks()

	c := make(chan string)

	for _, l := range list {
		go CheckLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l)
	}
}

// func AddURL(l string) []string {
// 	//add to file
//check for proper formatting
//add \n at the end
// }
