package go_utilities

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
}

func TestTools_RandomString2(t *testing.T) {
	var testTools Tools
	var testToolsArray []string
	n := 20
	for i := 0; i < n; i++ {
		//testToolsArray[i] = testTools.RandomString(n)
		s := testTools.RandomString(10)
		testToolsArray = append(testToolsArray, s)
	}
	for i := 0; i < len(testToolsArray)-1; i++ {
		for j := i + 1; j < len(testToolsArray); j++ {
			if testToolsArray[i] == testToolsArray[j] {
				t.Error("not all values are different")
			}
		}
	}
}
