package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	printResponseStatus(response)
	printResponseBody_method1(response)
	printResponseBody_method2(response)


}

func printResponseStatus(response *http.Response){
	fmt.Println("response.Status:", response.Status) // response.Status: 200 OK
}

func printResponseBody_method1(response *http.Response){
	/*
			Response: &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, max-age=0] Content-Type:[text/html; charset=ISO-8859-1] Date:[Thu, 13 Jan 2022 18:01:27 GMT] Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."] Server:[gws] Set-Cookie:[1P_JAR=2022-01-13-18; expires=Sat, 12-Feb-2022 18:01:27 GMT; path=/; domain=.google.com; Secure NID=511=hdnakqOzQCRYut8ZwUqa8r4pbPsmMAiXNxMAlqSCfJ4tN5VkXaSoYbcjcArEA6bwjyTtUZDAQUg19ZRjZ-OnZf1pK_IVl88w2YPxyaNqcyG3NluyGrRe3IgIP28d1uNefaREI2os7y9Ccsk1iXlpJ5LsIUllXJXxNsymdiFvmmU; expires=Fri, 15-Jul-2022 18:01:27 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN] X-Xss-Protection:[0]] 0xc0000be180 -1 [] false true map[] 0xc000216000 <nil>}
		<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en"><head><meta content="Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for." name="description"><meta content="noodp" name="robots"><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title>
	*/
	io.Copy(os.Stdout, response.Body)
}

func printResponseBody_method2(response *http.Response){
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
