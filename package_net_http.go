package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	/*
		Response:
		&{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, max-age=0]
		Content-Type:[text/html; charset=ISO-8859-1]
		Date:[Wed, 12 Jan 2022 23:29:32 GMT]
		Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."]
		Server:[gws]
		Set-Cookie:[1P_JAR=2022-01-12-23;
		expires=Fri, 11-Feb-2022 23:29:32 GMT; path=/;
		domain=.google.com;
		Secure NID=511=nov42aH7gIzFK78pFw4ZA9EB8SXSMfyoeVsvZ6AMx3XeXYVHGvOB8Jp7ykNLq6ev6AaSwEBPTWmpdJZkxABk-iP2mO_5V3OXzxLPfSrxZqjQakrkCT0vwFr4QgvTjWdEIcNa6h-t9j0Cg7NKItPlO0DaXOWL7PJbIlcLA0-VcA0; expires=Thu, 14-Jul-2022 23:29:32 GMT; path=/; domain=.google.com; HttpOnly]
		X-Frame-Options:[SAMEORIGIN]
		X-Xss-Protection:[0]] 0xc0000c8240 -1 [] false true map[] 0xc0000fc100 <nil>}
	*/
	fmt.Println("Response:", resp)
}
