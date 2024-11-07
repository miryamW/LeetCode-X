package main

import (
	"fmt"
	"LeetCode-server/services/runtests"
)

func main() {
	testID := "1" 
	err := runtests.runTests(testID)
	if err != nil {
		fmt.Println("Error running tests:", err)
		return
	}

	fmt.Println("Tests succeeded")
}
