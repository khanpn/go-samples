package main

import (
	"fmt"
	"regexp"
)

var digitRegexp = regexp.MustCompile(`\d`)

func main() {
	str := "asddk23242k3423j423423hb42h34g5465323"

	fmt.Println(digitRegexp.ReplaceAllString(str, ""))
}