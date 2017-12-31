package main

import "fmt"

func main() {
	prefix1 := Prefix{
		Address: "10.0.0.0",
		Mask:    31,
	}
	fmt.Printf("prefix1 address is %v and the mask is %v\n",
		prefix1.GetAddress(), prefix1.GetMask())
	prefix1.Reset()
	fmt.Printf("prefix1 address is %v and the mask is %v\n",
		prefix1.GetAddress(), prefix1.GetMask())
}
