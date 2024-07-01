package utils

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

func UserArgs() string {
	urlopt := flag.String("u", "example.com", "Target URL")
	flag.Parse()

	purl, err := url.Parse(*urlopt)
	if err != nil {
		panic("URL argument messed up")
	}
	fmt.Println("url:", purl)

	checkUrl(*urlopt)
	//if url has no trailing /, add.

	return *urlopt
}
func checkUrl(urlopt string) {
	// see if url resolves
	resp, err := http.Get(urlopt)
	if err != nil {
		fmt.Println("Error resolving host")
	} else {
		fmt.Println("resolve ok: ", resp.StatusCode)
	}

}
