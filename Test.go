package main

import "fmt"
import "math"
import "rsc.io/quote"
import "github.com/sambhav13/Repo1/mypkg1"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(math.Max(73.15, 92.46))
	fmt.Println(quote.Hello())
	fmt.Println(mypkg1.GetPkg1String())

}
