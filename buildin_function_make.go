package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	{
		fmt.Println("############## 初始化map用make ##############")
		var colors2 map[string]int
		fmt.Println(colors2) // map[]
		colors2 = make(map[string]int)
		colors2["pink"] = 213
		fmt.Println(colors2) // map[pink:213]
	}

	{
		response, error := http.Get("http://google.com")
		if error != nil {
			fmt.Println("Error:", error)
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
		fmt.Println("Response:", response)

		/*
			So this make function right here is a built-In function in the language that takes a type of a slice.
			And then as a second argument,
			this is the number of elements or empty spaces that we want to slice to be initialized with.
		*/
		bs := make([]byte, 99999)
		response.Body.Read(bs)
		/*
			<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en">
			<head><meta content="Search the world's information, including webpages, images, videos and more.
			Google has many special features to help you find exactly what you're looking for.
			" name="description"><meta content="noodp" name="robots"><meta content="text/html;
		*/
		fmt.Println(string(bs))
	}
}
