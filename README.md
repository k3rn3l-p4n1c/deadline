# DEADLINE!

[![Build Status](https://travis-ci.org/k3rn3l-p4n1c/deadline.svg?branch=master)](https://travis-ci.org/k3rn3l-p4n1c/deadline)

Go package for executing function with specified deadline.


## Example

```go
package main

import (
	"context"
	"time"
	"fmt"
	
	"github.com/k3rn3l-p4n1c/deadline"
	)


func main() {
	timeout := 2 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := deadline.Run(ctx, func(innerCtx context.Context) {
	    // do something that may take times
	})
	
	if err != nil {
		fmt.Println("timeout exceeded")
	}
}
```
