package links

import (
	"fmt"
	"net/http"
	"time"
)

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
