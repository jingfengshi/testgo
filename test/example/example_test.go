package example

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func ExampleMax(){
	fmt.Println(Max(1, 2))
	// Output:
	// 2
}


func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	// Output: olleh
}