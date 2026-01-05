# gobooty

Tiny Go helpers that turn factory functions into lazy, cached calls using `sync.Once`.

I created this because it's a pattern that I've been using for ages and decided to create a simple library.

## Install

```
go get github.com/Vizualni/gobooty
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/Vizualni/gobooty"
)

myBootstrappedValue = gobooty.One(func() string {
	fmt.Println("building")
	return "value"
})

myValueWithError = gobooty.Two(func() (string, error) {
	return "", fmt.Errorf("oh no")
})

func main() {

	fmt.Println(myBootstrappedValue()) // prints "building" and "value"
	fmt.Println(myBootstrappedValue()) // prints only "value"

	n, err := myValueWithError()
	fmt.Println(n, err) // prints empty string and an error
}
```

## API

- `One(func() T) func() T`: caches a single return value.
- `Two(func() (T1, T2)) func() (T1, T2)`: caches two return values.
