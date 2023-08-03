package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &str)

	str = strings.Replace(str, "c=", "!", -1)
	str = strings.Replace(str, "c-", "@", -1)
	str = strings.Replace(str, "dz=", "#", -1)
	str = strings.Replace(str, "d-", "$", -1)
	str = strings.Replace(str, "lj", "%", -1)
	str = strings.Replace(str, "nj", "^", -1)
	str = strings.Replace(str, "s=", "&", -1)
	str = strings.Replace(str, "z=", "*", -1)

	var alphabets []string
	for i := 0; i < len(str); i++ {
		alphabets = append(alphabets, string(str[i]))
	}

	fmt.Println(len(alphabets))
}