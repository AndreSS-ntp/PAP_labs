package arguments

import (
	"fmt"
	"os"
)

func GiveArguments() {
	fmt.Println("Given Arguments ", os.Args[1:])
}
