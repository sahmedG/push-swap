package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/smslash/push-swap/check/handle"
	"github.com/smslash/push-swap/check/proof"
	"github.com/smslash/push-swap/check/valid"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arr := make([]int, 0)
	a := make([]int, 0)
	b := make([]int, 0)
	s := strings.Fields(os.Args[1])
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			handle.Error()
		}
		arr = append(arr, n)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())
		res := ""

		for i := 0; i < len(str); i++ {
			res += valid.RemoveNewLine(str[i]) + " "
		}

		f := strings.Fields(res)
		valid.IsValid(f)
		arr, b = proof.Start(arr, b, f)
		if proof.IsSorted(arr) && len(b) == 0 {
			fmt.Println("OK")
			os.Exit(0)
		}
	}

	if proof.IsSorted(a) && len(b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	// scanner.Scan()
	// str := strings.Fields(scanner.Text())
	// res := ""

	// for i := 0; i < len(str); i++ {
	// 	res += valid.RemoveNewLine(str[i]) + " "
	// }

	// f := strings.Fields(res)
	// fmt.Println("Result: ", f)
	// valid.IsValid(f)
	// proof.Start(arr, f)
}
