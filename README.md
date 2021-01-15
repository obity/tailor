### tailor

tailor returns only the map of the specified keys

#### install
```
go get github.com/wandercn/tailor
```

#### example
```
package main

import (
	"fmt"

	"github.com/wandercn/tailor"
)

func main() {
	doc := tailor.Object{
		"A": 1,
		"B": tailor.Object{
			"age": 20},
		"C": "c",
		"D": 0.99,
	}
	fmt.Printf("before: %v\n", doc)

	keys := []string{"C", "D"}
	result := tailor.ClipOne(keys, doc)

	fmt.Printf("after: %v\n", result)

}

```
