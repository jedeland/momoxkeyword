package testing

import (
	"flag"
	"fmt"
)

type flagOutput struct {
	isTesting bool
} 

func InitialiseBasicFlags() *bool  {
	testingFlag := flag.Bool("testing", true, "Testing, true or false")

	flag.Parse()
	flagResult := flagOutput{isTesting: *testingFlag}
	fmt.Println(flagResult.isTesting)
	// Alterntive could be an interface of values relating to the flags
	return testingFlag
}