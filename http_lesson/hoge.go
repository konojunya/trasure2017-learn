package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	var (
		url    string
		header string
		method string
		data   string
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// header
	f.StringVar(&header, "header", "", "header flag")
	f.StringVar(&header, "h", "", "header flag")

	// method
	f.StringVar(&method, "X", "", "method flag")

	// data
	f.StringVar(&data, "d", "", "data flag")

	url = flag.Arg(0)

	if url == "" {
		fmt.Printf("URLがないよ")
		return
	}

	if method == "" {
		method = "GET"
	}

	req, _ := http.NewRequest(method, url, nil)

	fmt.Println(header)

	if header != "" {
		headers := strings.Split(header, ":")
		fmt.Print(len(headers))
		if len(headers) != 2 {
			fmt.Print(":で区切ってください")
			return
		}
		req.Header.Set(headers[0], headers[1])
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	_, _ = ioutil.ReadAll(resp.Body)
	// fmt.Printf("body = %v", string(b))
}
