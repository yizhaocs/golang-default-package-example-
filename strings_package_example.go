package main

import (
	"fmt"
	"strings"
)

func main(){
	ex1 := "Ace of Spades 1"
	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 输出 *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** Println ***************** ")
		fmt.Println(ex1) // Ace of Spades 1
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 是否存在 *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** Contains ***************** ")
		/*
			Contains reports whether substr is within s.
		 */
		fmt.Println(strings.Contains("seafood", "foo")) // true
		fmt.Println(strings.Contains("seafood", "bar")) // false
		fmt.Println(strings.Contains("seafood", "")) // true
		fmt.Println(strings.Contains("", "")) // true

		fmt.Println("***************** ContainsAny ***************** ")
		/*
			ContainsAny reports whether any Unicode code points in chars are within s.
		 */
		fmt.Println(strings.ContainsAny("team", "i")) // false
		fmt.Println(strings.ContainsAny("fail", "ui")) // true
		fmt.Println(strings.ContainsAny("ure", "ui")) // true
		fmt.Println(strings.ContainsAny("failure", "ui")) // true
		fmt.Println(strings.ContainsAny("foo", "")) // false
		fmt.Println(strings.ContainsAny("", "")) // false
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 统计 *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** Count ***************** ")
		/*
			Count counts the number of non-overlapping instances of substr in s.
			If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
		 */
		fmt.Println(strings.Count("cheese", "e")) // 3
		fmt.Println(strings.Count("five", "")) // before & after each rune // 5
	}


	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 对比 *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** Compare ***************** ")
		/*
			Compare returns an integer comparing two strings lexicographically.
			The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
		*/
		fmt.Println(strings.Compare("a", "b")) // -1
		fmt.Println(strings.Compare("a", "a")) // 0
		fmt.Println(strings.Compare("b", "a")) // 1
		fmt.Println(strings.Compare("c", "a")) // 1

		fmt.Println("***************** EqualFold ***************** ")
		/*
			EqualFold reports whether s and t,
			interpreted as UTF-8 strings,
			are equal under Unicode case-folding,
			which is a more general form of case-insensitivity.
		*/
		fmt.Println(strings.EqualFold("Go", "Go")) // true
		fmt.Println(strings.EqualFold("Go", "go")) // true
		fmt.Println(strings.EqualFold("Go", "Goo")) // false
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 切分 *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** Fields ***************** ")
		/*
			Fields splits the string s around each instance of one or more consecutive white space characters,
			as defined by unicode.IsSpace,
			returning a slice of substrings of s or an empty slice if s contains only white space.
		 */
		fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]

	}
}
