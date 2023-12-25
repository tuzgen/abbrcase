# abbrcase
abbrcase is a very simple Go linter that fails and reports when abbreviations in identifiers are written in all caps

## Install
```
go install github.com/tuzgen/abbrcase@latest
```

## Rules
The rules below apply to abbreviations defined under the abbrs flag.

- Any identifier that matches an abbreviation should be either all uppercase or all lowercase
e.g. http or HTTP, not Http
### Example
When abbrs = http
```go
package http // package name should not be reported as an error

import "fmt"

func findOrderByOrderID(orderId string) struct{} {
	return struct{}{}
}

func createHttpServer() {
    //...
}

func createHTTPServer() {
    //...
}

func main() {
	orderID := "asd"
	orderId := "Ads"

	createHTTPServer()
	createHttpServer()
	val := findOrderByOrderID(orderID)
	fmt.Println(orderID, orderId, val)
}
``` 

#### Output
```
abbrcase 
on  main [?] via  v1.21.3 
❮ ./abbrcase -abbrs=http ./...
/Users/oguztuzgen/Development/personal/abbrcase/examples/example.go:9:6: use all caps abbreviations: Http should be HTTP
/Users/oguztuzgen/Development/personal/abbrcase/examples/example.go:22:2: use all caps abbreviations: Http should be HTTP

```

### Why do you need this?
Any form of static code analysis aims to improve code quality and avoid samey code review comments. This project also aims to tackle the same issues. It is just a niche problem I encountered in my personal experience. 

