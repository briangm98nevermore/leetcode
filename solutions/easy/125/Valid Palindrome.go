package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	A phrase is a palindrome if, after converting all uppercase letters into
	lowercase letters
	and removing all non-alphanumeric characters, it reads the same forward and
	backward.
	Alphanumeric characters include letters and numbers.

	Given a string s, return true if it is a palindrome, or false otherwise.

	Example 1:
	Input: s = "A man, a plan, a canal: Panama"
	Output: true
	Explanation: "amanaplanacanalpanama" is a palindrome.

	Example 2:
	Input: s = "race a car"
	Output: false
	Explanation: "raceacar" is not a palindrome.

	Example 3:
	Input: s = " "
	Output: true
	Explanation: s is an empty string "" after removing non-alphanumeric characters.
	Since an empty string reads the same forward and backward, it is a palindrome.

	Constraints:

		1 <= s.length <= 2 * 105
		s consists only of printable ASCII characters.
*/

func isPalindrome(s string) bool {

	result := clean(s)

	for i,j := 0, len(result) -1; i < j; i,j = i+1, j-1 {

		if result[i] != result[j] {
			return false
		}
	}
	return true
}

func clean(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return unicode.ToLower(r)
        }
        return -1
    }, s)
}

func main()  {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}