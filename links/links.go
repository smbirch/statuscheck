package links

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	if _, err := os.Stat("links.txt"); errors.Is(err, os.ErrNotExist) {
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

func ReadList() []string {
	file, err := os.Open("links.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var list []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal("Error reading file:", err)
		}
		list = append(list, s.Text())
	}
	return list
}

//This is now not needed, functions should make calls to ReadList now
//TODO: Delete this function and rework all calls to it
func GetLinks() []string {
	links := ReadList()
	return links
}

func CheckLink(link string, c chan string) {
	b, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link
		return
	}
	func(link string) {

		fmt.Println(b.Status + " " + link)
		c <- link
	}(link)
}

func LoopLinks() {
	list := ReadList()

	c := make(chan string)
	for _, l := range list {
		l = strings.TrimSpace(l)
		go CheckLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(1 * time.Second)

			CheckLink(link, c)
		}(l)
	}
}
