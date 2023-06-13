# Initial config for Go

we create a config file where intalling all dependencies

```sh
go mod init
```

this created `go.mod` and we could create a file with some code of go like `main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hola desde el main")
}
```

And now we can type the next on terminal

```sh
touch main.go
go run .
go build main.go
```