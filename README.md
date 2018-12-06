# DEADLINE!

[![Build Status](https://travis-ci.org/k3rn3l-p4n1c/deadline.svg?branch=master)](https://travis-ci.org/k3rn3l-p4n1c/deadline)
[![Go Report Card](https://goreportcard.com/badge/github.com/k3rn3l-p4n1c/deadline)](https://goreportcard.com/report/github.com/k3rn3l-p4n1c/deadline)

Go package for running functions in a hardly limited time.

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
