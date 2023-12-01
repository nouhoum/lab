package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("simplify path")

	path := "/home//toto//./b/"

	fmt.Println("final result=", simplifyPath(path))
}

func simplifyPath(path string) string {
	cleanDoubleSlash := func(path string) string {
		return strings.ReplaceAll(path, "//", "/")
	}

	withOutDoubleSlashes := cleanDoubleSlash(path)
	es := strings.Split(withOutDoubleSlashes, "/")
	fmt.Println("===> es=", es)
	fmt.Println("===> es joined=", strings.Join(es, "/"))

	cleanDots := func(path string) string {
		return strings.ReplaceAll(path, "/.", "")
	}
	newPath, _ := strings.CutSuffix(cleanDots(withOutDoubleSlashes), "/")

	return newPath
}
