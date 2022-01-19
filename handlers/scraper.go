package handlers

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var client http.Client

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FetchURL(url string) {

	response, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of bytes copied to STDOUT:", n)
}

func FetchAndSave(url string, fileExt string) {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("./media/" + strconv.Itoa(rand.Int()) + "." + fileExt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}

func Sanitizer(url string) string {
	//Must find the last / and then check for the next .
	//after that, if it finds any forbiden simbol then deletes everything after that
	domainName := GetDomainName(url)
	var match string
	switch domainName {
	case "www.twitter.com":
		var re = regexp.MustCompile(`.+status/\d*/?`)
		match = re.FindString(url)

	case "www.pixiv.com":
		var re = regexp.MustCompile(`.+artworks/\d*/?`)
		match = re.FindString(url)

	case "www.artstation.com":
		var re = regexp.MustCompile(`.+artwork/.*/?`)
		match = re.FindString(url)

	case "www.instagram.com":
		var re = regexp.MustCompile(`.+\/p\/.*\/`)
		match = re.FindString(url)
	}

	return match
}

func GetFileExtension(url string) string {

	var re = regexp.MustCompile(`format=[a-z]*`)
	match := re.FindString(url)

	return match[7:]
}

func GetDomainName(url string) string {
	var re = regexp.MustCompile(`[a-zA-Z]*\.[a-zA-Z]+\.[a-zA-Z]+`)
	match := re.FindString(url)

	return match
}
