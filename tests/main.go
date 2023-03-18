/**
  @author: Zero
  @date: 2023/3/18 12:08:06
  @desc:

**/

package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Duration(time.Duration(1).Hours() * 24)
	fmt.Println(a.String())
}
