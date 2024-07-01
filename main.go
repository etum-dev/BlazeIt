package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/nekopizza/BlazeIt/utils"
)

func main() {
	url := utils.UserArgs()
	visitPage(url)
}
func readLists() /*([]string)*/ {
	// TODO: allow user to specify file
	fwrk_files, err := os.Open("lists/framework.txt")
	if err != nil {
		panic(err)
	}
	defer fwrk_files.Close()
	line := bufio.NewScanner(fwrk_files)
	for line.Scan() {
		fmt.Println(line.Text())
	}

}

func visitPage(urlopt string) {
	url_endpoints := "_framework/blazor.boot.json"
	fmt.Println(urlopt + url_endpoints)
	res, err := http.Get(urlopt + url_endpoints)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

}
