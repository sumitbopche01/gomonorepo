package subpackage1

import (
	"fmt"

	"github.com/sumitbopche01/gomonorepo/subpackage2"

)

func printPackage1(){
	fmt.Println("Hello, world")

	subpackage2.PrintPackage2()
}
