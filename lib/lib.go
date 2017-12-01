package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Mac struct is used to store useful data about MAC addresses
type Mac struct {
	Address, Company string
}

// CheckErr is a re-usable function for checking errors
func CheckErr(msg string, e error) {
	if e != nil {
		log.Fatalf("%v: %v", msg, e)
	}
}

// GetResult uses HTTP and fetches the body of the API call
func GetResult(mac Mac) string {
	base := "http://api.macvendors.com/"
	url := fmt.Sprint(base, url.QueryEscape(mac.Address))

	req, err := http.NewRequest("GET", url, nil)
	CheckErr("Couldn't construct new request", err)

	client := &http.Client{}

	res, err := client.Do(req)
	CheckErr("Couldn't complete request", err)

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "something went wrong"
	}
	b, _ := ioutil.ReadAll(res.Body)
	return string(b)
}
