package do_test

import (
	"fmt"

	"github.com/aquilax/do"
)

func ExampleRun() {
	result := do.Run(func() string { return "Hello from inside" })
	meaning := 42
	answer := do.Run(func() string { return fmt.Sprintf("It must be %d", meaning) })
	fmt.Println(result)
	fmt.Println(answer)

	// Output:
	// Hello from inside
	// It must be 42
}
