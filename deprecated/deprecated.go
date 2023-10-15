package deprecated

import "fmt"

// HelloNotDeprecated prints a message.
// It's not deprecated at all and shouldn't have any DEPRECATED tag on pkg.go.dev.
func HelloNotDeprecated() {
	fmt.Println("Hello, World!")
}

// HelloDeprecatedWithNewLine prints a message.
// This function is deprecated with a special comment below,
// preceeded by a newline.
// I expect that it will have a DEPRECATED tag on pkg.go.dev.
//
// Deprecated: use HelloNotDeprecated() instead
func HelloDeprecatedWithNewline() {
	fmt.Println("Hello, World!")
}

// HelloDeprecatedWithoutNewLine prints a message.
// This function is deprecated with a special comment below.
// However, because it's not preceeded by a newline,
// I don't expect go.pkg.mod to actually show a DEPRECATED tag for it.
// Deprecated: use HelloNotDeprecated() instead
func HelloDeprecatedWithoutNewline() {
	fmt.Println("Hello, World!")
}
