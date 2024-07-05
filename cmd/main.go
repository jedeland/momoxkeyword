package main

import (
	"fmt"
	"momoxkeyword/internal/scraping"
	"momoxkeyword/internal/testing"
)

func main() {
	basicFlags := testing.InitialiseBasicFlags()
	fmt.Println("Is this a test? ", *basicFlags )
	
	scraping.DetermineUrlsTest()
	fmt.Println("heyo!")
}