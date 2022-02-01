package links

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// type URLList struct {
// 	links []string `json:"links"`
// }

//Will write a newly updated list and return said list
func BuildList() {

	basicList := []string{
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

	//checks for existing file, creates if not found
	if _, err := os.Stat("links.txt"); errors.Is(err, os.ErrNotExist) { //TODO: Run this at app launch
		// path/to/whatever does not exist
		file, err := os.Create("links.txt")
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		for i, _ := range basicList { //TODO seperate written strings in file by line
			_, err := file.WriteString(basicList[i])
			if err != nil {
				log.Fatal("Error writing to file:", err)
			}
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
// }
