package casify

import (
	"fmt"
	"strings"
	"testing"
)

// Generate fuzzing tests
//go:generate go run fuzzing_gen.go

func ExampleConvert() {
	// Here is a few basic examples
	fmt.Println(Convert("helloWorld", "_", strings.ToLower))          // camelCase to snake_case
	fmt.Println(Convert("http_status_code", "", strings.Title))       // snake_case to CamelCase
	fmt.Println(Convert("XRequestID", "-", strings.ToUpper))          // CamelCase with acronyms to kebab-case
	fmt.Println(Convert("AnHTTPRequestWithID", "_", strings.ToLower)) // A little bit harder acronym case (pun intended)

	// Some advanced examples
	fmt.Println(Convert("anyString with ArbitraryCaseChanges, punctuation!and  \t\n space-separation",
		"-", strings.ToLower))
	fmt.Println(Convert("ПОДДЕРЖКА___юникода", "", func(s string) string {
		return strings.Title(strings.ToLower(s))
	}))
	fmt.Println(Convert("Any non-alphanumeric symbols: !@#$ are skipped%^&*()", "", strings.Title))
	fmt.Println(Convert("Thus, strings with numeric (123) identifiers appended321 work as intended.", "", strings.Title))
	// Output:
	// hello_world
	// HttpStatusCode
	// X-REQUEST-ID
	// an_http_request_with_id
	// any-string-with-arbitrary-case-changes-punctuation-and-space-separation
	// ПоддержкаЮникода
	// AnyNonAlphanumericSymbolsAreSkipped
	// ThusStringsWithNumeric123IdentifiersAppended321WorkAsIntended
}

func ExampleUntitle() {
	fmt.Println(Untitle("STRING"))
	fmt.Println(Untitle("A String With Words In It"))
	// Output:
	// sTRING
	// a string with words in it
}

func ExampleSnakeCase() {
	fmt.Println(SnakeCase("camelCase"))
	fmt.Println(SnakeCase("CAMELCaseWithACRONYMS"))
	fmt.Println(SnakeCase("kebab-case"))
	fmt.Println(SnakeCase("UPPER-KEBAB-CASE"))
	fmt.Println(SnakeCase("snake_case"))
	fmt.Println(SnakeCase("UPPER_SNAKE_CASE"))
	fmt.Println(SnakeCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	// Output:
	// camel_case
	// camel_case_with_acronyms
	// kebab_case
	// upper_kebab_case
	// snake_case
	// upper_snake_case
	// with_mixed_caps_and_mixed_separators_also
}

func ExampleUpperSnakeCase() {
	fmt.Println(UpperSnakeCase("camelCase"))
	fmt.Println(UpperSnakeCase("CAMELCaseWithACRONYMS"))
	fmt.Println(UpperSnakeCase("kebab-case"))
	fmt.Println(UpperSnakeCase("UPPER-KEBAB-CASE"))
	fmt.Println(UpperSnakeCase("snake_case"))
	fmt.Println(UpperSnakeCase("UPPER_SNAKE_CASE"))
	fmt.Println(UpperSnakeCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	// Output:
	// CAMEL_CASE
	// CAMEL_CASE_WITH_ACRONYMS
	// KEBAB_CASE
	// UPPER_KEBAB_CASE
	// SNAKE_CASE
	// UPPER_SNAKE_CASE
	// WITH_MIXED_CAPS_AND_MIXED_SEPARATORS_ALSO
}

func ExampleKebabCase() {
	fmt.Println(KebabCase("camelCase"))
	fmt.Println(KebabCase("CAMELCaseWithACRONYMS"))
	fmt.Println(KebabCase("kebab-case"))
	fmt.Println(KebabCase("UPPER-KEBAB-CASE"))
	fmt.Println(KebabCase("snake_case"))
	fmt.Println(KebabCase("UPPER_SNAKE_CASE"))
	fmt.Println(KebabCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	// Output:
	// camel-case
	// camel-case-with-acronyms
	// kebab-case
	// upper-kebab-case
	// snake-case
	// upper-snake-case
	// with-mixed-caps-and-mixed-separators-also
}

func ExampleUpperKebabCase() {
	fmt.Println(UpperKebabCase("camelCase"))
	fmt.Println(UpperKebabCase("CAMELCaseWithACRONYMS"))
	fmt.Println(UpperKebabCase("kebab-case"))
	fmt.Println(UpperKebabCase("UPPER-KEBAB-CASE"))
	fmt.Println(UpperKebabCase("snake_case"))
	fmt.Println(UpperKebabCase("UPPER_SNAKE_CASE"))
	fmt.Println(UpperKebabCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	// Output:
	// CAMEL-CASE
	// CAMEL-CASE-WITH-ACRONYMS
	// KEBAB-CASE
	// UPPER-KEBAB-CASE
	// SNAKE-CASE
	// UPPER-SNAKE-CASE
	// WITH-MIXED-CAPS-AND-MIXED-SEPARATORS-ALSO
}

func ExampleCamelCase() {
	fmt.Println(CamelCase("camelCase"))
	fmt.Println(CamelCase("CAMELCaseWithACRONYMS"))
	fmt.Println(CamelCase("kebab-case"))
	fmt.Println(CamelCase("UPPER-KEBAB-CASE"))
	fmt.Println(CamelCase("snake_case"))
	fmt.Println(CamelCase("UPPER_SNAKE_CASE"))
	fmt.Println(CamelCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	fmt.Println(CamelCase("aa4N7"))
	// Output:
	// camelCase
	// camelCaseWithAcronyms
	// kebabCase
	// upperKebabCase
	// snakeCase
	// upperSnakeCase
	// withMixedCapsAndMixedSeparatorsAlso
	// aa4N7
}

func ExampleUpperCamelCase() {
	fmt.Println(UpperCamelCase("camelCase"))
	fmt.Println(UpperCamelCase("CAMELCaseWithACRONYMS"))
	fmt.Println(UpperCamelCase("kebab-case"))
	fmt.Println(UpperCamelCase("UPPER-KEBAB-CASE"))
	fmt.Println(UpperCamelCase("snake_case"))
	fmt.Println(UpperCamelCase("UPPER_SNAKE_CASE"))
	fmt.Println(UpperCamelCase("WITH-mixedCAPS_and   MixedSeparators, \t also!"))
	// Output:
	// CamelCase
	// CamelCaseWithAcronyms
	// KebabCase
	// UpperKebabCase
	// SnakeCase
	// UpperSnakeCase
	// WithMixedCapsAndMixedSeparatorsAlso
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("Thus, strings with numeric (123) identifiers appended321 work as intended.", "", strings.Title)
	}
}

func BenchmarkConvertSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("CAMELCaseWithACRONYMS", "", strings.Title)
	}
}