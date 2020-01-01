# go-ms-utils

CheckToken usage

```go
package main

import (
	"fmt"
	oauth "github.com/alapierre/go-ms-utils/auth"
)

func main() {
	
	res, err := oauth.CheckToken("http://localhost:9001/oauth", "2e4228e4-ffd7-41ff-abf5-aa5d105abd79")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Printf("%#v\n", res)
	
}
```

GetToken usage

```go
package main

import (
	"fmt"
	oauth "github.com/alapierre/go-ms-utils/auth"
)

func main() {

	res, err := oauth.GetToken("http://localhost:9001/oauth", "user", "pass", "...", "...")

	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}
```
