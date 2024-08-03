package main

import (
	"fmt"
	"net/url"
)

// const url string = "https://www.amazon.com/SAMSUNG-Android-Included-Expandable-Exclusive/dp/B0CWS8MNW1?ref=dlx_deals_dg_dcl_B0CWS8MNW1_dt_sl14_7e"
const myurl string = "https://www.amazon.com:8000/learn?coursename=reactjs&paymentid=lkdfsadafk"

func main() {
	fmt.Println("my Handling in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.amazon.com",
		Path:    "/testpath",
		RawPath: "user=testrawdata",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
