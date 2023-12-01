package main

import (
	"fmt"
)

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true

*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	l := len(s)
	ss := make(map[byte]byte)
	tt := make(map[byte]byte)

	for i := 0; i < l; i++ {
		si := s[i]
		ti := t[i]
		_, ok1 := ss[si]
		_, ok2 := tt[ti]

		if !ok1 && !ok2 {
			ss[si] = ti
			tt[ti] = si
		} else if ss[si] != ti || tt[ti] != si {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Testing isIsomorphic...")
	checkIsIsomorphic(true, "egg", "add")
	checkIsIsomorphic(false, "foo", "bar")
	checkIsIsomorphic(true, "paper", "title")
	checkIsIsomorphic(false, "badc", "baba")
	fmt.Println("Testing isIsomorphic...OK")
}

func checkIsIsomorphic(expected bool, s1, s2 string) {
	actual := isIsomorphic(s1, s2)
	if actual != expected {
		panic(fmt.Errorf("Oops wrong response. found %d, expected", actual, expected))
	}
}
