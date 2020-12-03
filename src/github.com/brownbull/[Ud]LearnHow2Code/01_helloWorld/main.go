// # GitBash - Windows
// cd "src\github.com\brownbull\[Ud]LearnHow2Code\01_helloWorld"
// # RUN
// go run main.go
// # BUILD
// go build
// # EXEC
// ./01_helloWorld.exe
// # INSTALL - GOBIN must be a valid env var, add to path too for direct exec
// go install main.go
// main.exe
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
}
