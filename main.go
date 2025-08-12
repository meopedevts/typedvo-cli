package main

import (
	"os"

	"github.com/meopedevts/typedvo-cli/cmd"
)

func main() {
	code := cmd.Execute()
	os.Exit(int(code))
}
