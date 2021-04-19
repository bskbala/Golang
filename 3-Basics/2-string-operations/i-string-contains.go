package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n ====stringContains====")
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println()

	// String.ContainsAny

	fmt.Println("\n ====stringContainsany====")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println()

	// String ContainsRune

	fmt.Println("\n ====StringContainsRune====")
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	fmt.Println()

	// String Count

	fmt.Println("\n ====StringCount====")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", ""))
	fmt.Println()

	// String EqualFold

	fmt.Println("\n ====StringEqualFold====")
	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("GO", "GO"))
	fmt.Println()

	// String HasPrefix
	fmt.Println("\n ====StringHasPrefix====")
	fmt.Println(strings.HasPrefix("Gopher", "go"))
	fmt.Println(strings.HasPrefix("Gopher", "go"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println()

	// String HasSuffix

	fmt.Println("\n ====StringHasPrefix====")
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasPrefix("Amigo", ""))
	fmt.Println()

	// String Index
	fmt.Println("\n ====StringIndex====")
	fmt.Println(strings.Index("chicken", "c"))
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println()


}
