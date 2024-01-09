package oxo

import (
	"testing"
	// "fmt"
)

func TestSetXwin(t *testing.T) {

// REWRITE THIS WITH LOGS.....

	afail := SetXwin("XX       ")
	apass := SetXwin("XXX      ")
	b := IsXwin
	if afail == b {

		t.Errorf("Expected fail\t%08b got \t%08b", b, afail)
	}
	if apass != b {
		t.Errorf("Expected pass\t%08b got \t%08b", b, apass)
	}
	//   fmt.Printf("apass is \t%08b afail is \t%08b does apass==afail? %t IsXwin is \t%08b",apass,afail,apass==afail,IsXwin)
	//	fmt.Printf("afail==IsWin\t%t",afail==IsXwin)
}
func TestTableSetXwin(t *testing.T) {

	var tests = []struct {
		input    string
		expected GameStateflags
	}{
		{"O O   XXX", IsXwin},
		{"XXX   OO ", IsXwin},
		{"OO XXXO  ", IsXwin},
		{"X  X  XOO", IsXwin},
		{" X  X OXO", IsXwin},
		{"OOX  X  X", IsXwin},
		{"X   XO  X", IsXwin},
		{"  X X XOO", IsXwin},
		{"XXXXXXXXX", IsXwin},
		{"         ", Zero},
		{"OOOOOO   ", Zero},
		{"OOOOOOOOO", Zero},
		{"XOX OOOXX", Zero},
	}
	for _, test := range tests {
		if output := SetXwin(test.input); output != test.expected {

			t.Errorf("Test failed: {%s} inputted,{%08b} expected, received:{%08b}  ", test.input, test.expected, output)
		}
	}

}
