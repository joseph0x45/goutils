package goutils_test

import (
	"fmt"
	"testing"

	"github.com/joseph0x45/goutils"
)

func TestParsedDotenv(t *testing.T) {
	dummyContent := `PORT=8080
GRUH=asdasd
BRUH=askdads
`
	parsed := goutils.ParseSimpleDotenv(dummyContent)
	fmt.Printf("%+v\n", parsed)
	port := parsed.GetKey("PORT")
	gruh := parsed.GetKey("GRUH")
	bruh := parsed.GetKey("BRUH")
	fmt.Printf("%s %s %s\n", port, gruh, bruh)
	parsed.SetKey("PORT", "6969")
	fmt.Printf("%s\n", parsed.GetKey("PORT"))
	fmt.Println(parsed.Write())

}
