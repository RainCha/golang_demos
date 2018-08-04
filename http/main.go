package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func getMooc()  {
	resp , err := http.Get("http://www.imooc.com")

	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

func getMobileMooc(){

	request, err:= http.NewRequest(http.MethodGet,"https://www.imooc.com/", nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Mobile Safari/537.36")

	resp , err := http.DefaultClient.Do(request)

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", s)
}


func main() {
	//getMooc()

	getMobileMooc()
}

