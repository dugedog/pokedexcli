package main
import ("testing")

func TestCleanInput(t *testing.T) {
			
	cases := []struct {},
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{ 
			input: "bob harry tim jim",
			expected: []string{"bob","harry","tim","jim"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
		t.Errorf("Error, actual length (%v) not matching target (%v)", len(actual), len(c.expected))
 		// error and continue here
	}
	length := len(actual)
  	for i:=0; i<length; i++ {
  		word := actual[i]
		if i >= len(c.expected){
			t.Errorf("Out of Range of expected slice")
			break
			}
  		expectedWord := c.expected[i]
  		if word != expectedWord {
  			t.Errorf("Actual Word (%s), does not match expected word (%v)", word, expectedWord)
  			}
  		// Check each word in the slice
  		// if they don't match, use t.Errorf to print an error message
  		// and fail the test
  	}
}

}
