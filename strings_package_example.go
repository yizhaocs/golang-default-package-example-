package main

import (
	"fmt"
	"strings"
	"unicode"
)
/*
	RFE:
		https://pkg.go.dev/strings@go1.17.6
 */
func main() {
	ex1 := "Ace of Spades 1"
	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 输出 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Println ***************** ")
			fmt.Println(ex1) // Ace of Spades 1
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* index 位置 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** index ***************** ")
			/*
				Index returns the index of the first instance of substr in s,
				or -1 if substr is not present in s.
			*/
			fmt.Println(strings.Index("chicken", "ken")) // 4
			fmt.Println(strings.Index("chicken", "dmr")) // -1
		}
		{
			fmt.Println("***************** IndexAny ***************** ")
			/*
				IndexAny returns the index of the first instance of any Unicode code point from chars in s,
				or -1 if no Unicode code point from chars is present in s.
			*/
			fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
			fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1
		}
		{
			fmt.Println("***************** LastIndex ***************** ")
			/*
				LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
			*/
			fmt.Println(strings.Index("go gopher", "go"))         // 0
			fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
			fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
		}
		{
			fmt.Println("***************** LastIndexAny ***************** ")
			/*
				LastIndexAny returns the index of the last instance of any Unicode code point from chars in s,
				or -1 if no Unicode code point from chars is present in s.
			*/
			fmt.Println(strings.LastIndexAny("go gopher", "go"))     // 4
			fmt.Println(strings.LastIndexAny("go gopher", "rodent")) // 8
			fmt.Println(strings.LastIndexAny("go gopher", "fail"))   // -1
		}
	}
	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 是否存在 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Contains ***************** ")
			/*
				Contains reports whether substr is within s.
			*/
			fmt.Println(strings.Contains("seafood", "foo")) // true
			fmt.Println(strings.Contains("seafood", "bar")) // false
			fmt.Println(strings.Contains("seafood", ""))    // true
			fmt.Println(strings.Contains("", ""))           // true
		}
		{
			fmt.Println("***************** ContainsAny ***************** ")
			/*
				ContainsAny reports whether any Unicode code points in chars are within s.
			*/
			fmt.Println(strings.ContainsAny("team", "i"))     // false
			fmt.Println(strings.ContainsAny("fail", "ui"))    // true
			fmt.Println(strings.ContainsAny("ure", "ui"))     // true
			fmt.Println(strings.ContainsAny("failure", "ui")) // true
			fmt.Println(strings.ContainsAny("foo", ""))       // false
			fmt.Println(strings.ContainsAny("", ""))          // false
		}
		{
			fmt.Println("***************** HasPrefix ***************** ")
			fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
			fmt.Println(strings.HasPrefix("Gopher", "C"))  // false
			fmt.Println(strings.HasPrefix("Gopher", ""))   // true
		}
		{
			fmt.Println("***************** HasSuffix ***************** ")
			fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
			fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
			fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
			fmt.Println(strings.HasSuffix("Amigo", ""))    // true
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 统计 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Count ***************** ")
			/*
				Count counts the number of non-overlapping instances of substr in s.
				If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
			*/
			fmt.Println(strings.Count("cheese", "e")) // 3
			fmt.Println(strings.Count("five", ""))    // before & after each rune // 5
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 对比 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Compare ***************** ")
			/*
				Compare returns an integer comparing two strings lexicographically.
				The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
			*/
			fmt.Println(strings.Compare("a", "b")) // -1
			fmt.Println(strings.Compare("a", "a")) // 0
			fmt.Println(strings.Compare("b", "a")) // 1
			fmt.Println(strings.Compare("c", "a")) // 1
		}
		{
			fmt.Println("***************** EqualFold ***************** ")
			/*
				EqualFold reports whether s and t,
				interpreted as UTF-8 strings,
				are equal under Unicode case-folding,
				which is a more general form of case-insensitivity.
			*/
			fmt.Println(strings.EqualFold("Go", "Go"))  // true
			fmt.Println(strings.EqualFold("Go", "go"))  // true
			fmt.Println(strings.EqualFold("Go", "Goo")) // false
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 切分 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Split ***************** ")
			/*
				Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

				If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.

				If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.

				It is equivalent to SplitN with a count of -1.
			*/
			fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
			fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
			fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
			fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]
		}
		{
			fmt.Println("***************** SplitAfter ***************** ")
			/*
				SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.

				If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.

				If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.

				It is equivalent to SplitAfterN with a count of -1.
			*/
			fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
		}
		{
			fmt.Println("***************** SplitAfterN ***************** ")
			/*
				SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.

				The count determines the number of substrings to return:
					n > 0: at most n substrings; the last substring will be the unsplit remainder.
					n == 0: the result is nil (zero substrings)
					n < 0: all substrings
				Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.
			*/
			fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
		}
		{
			fmt.Println("***************** SplitN ***************** ")
			/*
				SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.

				The count determines the number of substrings to return:

					n > 0: at most n substrings; the last substring will be the unsplit remainder.
					n == 0: the result is nil (zero substrings)
					n < 0: all substrings
				Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for Split.
			*/
			fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
			z := strings.SplitN("a,b,c", ",", 0)
			fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
		}
		{
			fmt.Println("***************** Fields ***************** ")
			/*
				Fields splits the string s around each instance of one or more consecutive white space characters,
				as defined by unicode.IsSpace,
				returning a slice of substrings of s or an empty slice if s contains only white space.
			*/
			fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]
		}
		{
			fmt.Println("***************** FieldsFunc ***************** ")
			/*
				FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s.
				If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.

				FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.
			*/
			f := func(c rune) bool {
				return !unicode.IsLetter(c) && !unicode.IsNumber(c)
			}
			fmt.Printf("Fields are: %q", strings.FieldsFunc("  abc;defg,xy!z...", f)) // ["abc" "defg" "xy" "z"]
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 合并 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Join ***************** ")
			/*
				Join concatenates the elements of its first argument to create a single string.
				The separator string sep is placed between elements in the resulting string.
			*/
			s := []string{"foo", "bar", "baz"}
			fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 更换 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** Replace ***************** ")
			/*
				Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
				If old is empty, it matches at the beginning of the string and after each UTF-8 sequence,
				yielding up to k+1 replacements for a k-rune string.
				If n < 0, there is no limit on the number of replacements.
			*/
			fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
			fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
		}
		{
			fmt.Println("***************** ReplaceAll ***************** ")
			/*
				ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
				If old is empty, it matches at the beginning of the string and after each UTF-8 sequence,
				yielding up to k+1 replacements for a k-rune string.
			*/
			fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) // moo moo moo
		}
	}

	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* 大小写转换 *******************************")
		fmt.Println("******************************************************************** ")
		{
			fmt.Println("***************** ToLower ***************** ")
			/*
				ToLower returns s with all Unicode letters mapped to their lower case.
			*/
			fmt.Println(strings.ToLower("Gopher")) // gopher
		}
		{
			fmt.Println("***************** ToUpper ***************** ")
			/*
				ToUpper returns s with all Unicode letters mapped to their upper case.
			*/
			fmt.Println(strings.ToUpper("Gopher")) // GOPHER
		}
	}
}
