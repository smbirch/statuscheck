package links

import (
	"fmt"
	"net/http"
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

func LinkBuffer() {

}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up - 200 OK")
	c <- link
}

func AddLinks(links []string) []string {
	//Adding more links
	return links
}
