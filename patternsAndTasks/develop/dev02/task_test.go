package main

import "testing"

func TestUnpackString(t *testing.T) {
	t.Log("Test unpackString behavior on different correct strings.")
	{
		inputOutput := make(map[string]string)
		inputOutput["a4bc2d5e"] = "aaaabccddddde"
		inputOutput["abcd"] = "abcd"
		inputOutput["abc2yelp3"] = "abcсyelppp" // last is digit

		testID := 0
		{
			for input, expectedOutput := range inputOutput {
				t.Logf("\tTest %d:\t`%s`.", testID, input)
				testUnpack(input, expectedOutput, true, t)
				testID++
			}
		}
	}
}

func testUnpack(input, expectedOutput string, correctResultExpected bool, t *testing.T) {
	str, err := unpackString(input)
	if str != expectedOutput {
		t.Errorf("\tExpect: `%s`.\n\t\tResult: `%s`.", expectedOutput, str)
	}
	if err != nil && correctResultExpected == true {
		t.Error("Expected no errors")
	}
}
