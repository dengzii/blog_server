package tools

import "fmt"

func Log(l ...interface{}) {

	for _, v := range l {
		fmt.Print(v)
	}
	fmt.Println()
}
