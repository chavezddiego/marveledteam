package main

import (
	"fmt"
	"net/url"
)

func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		fmt.Println(err)
	}
	return u.String()
}
func main() {
	/*
		var Url *url.URL
		Url, err := url.Parse("http://www.example.com")
		if err != nil {
			panic("boom")
		}

		Url.Path += "/some/path/or/other_with_funny_characters?_or_not/"

		fmt.Printf("Previa %q\n", Url.String())

		parameters := url.Values{}
		parameters.Add("hello", "42")
		parameters.Add("hello", "54")
		parameters.Add("vegetable", "potato")
		Url.RawQuery = parameters.Encode()
	*/
	hero := UrlEncoded("iron man")

	fmt.Println(hero)
	//fmt.Printf("Encoded URL is %q\n", Url.String())
}
