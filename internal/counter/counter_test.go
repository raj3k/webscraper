package counter

import "testing"

func TestMostFrequentWords(t *testing.T) {
	tests := []struct {
		input    string
		num      int
		expected []WordCount
	}{
		{
			input: "This is a sample string. This string contains sample words, and this is a sample sentence.",
			num:   3,
			expected: []WordCount{
				{"sample", 3},
				{"This", 2},
				{"is", 2},
			},
		},
		{
			input: "Hello world! Hello Go! Go is awesome!",
			num:   2,
			expected: []WordCount{
				{"Hello", 2},
				{"Go", 2},
			},
		},
	}

	for _, test := range tests {
		result := MostFrequentWords(test.input, test.num)

		if len(result) != len(test.expected) {
			t.Errorf("Expected %d most frequent words, but got %d", len(test.expected), len(result))
		}

		for i, expected := range test.expected {
			if result[i] != expected {
				t.Errorf("Expected %s: %d, but got %s: %d", expected.Word, expected.Count, result[i].Word, result[i].Count)
			}
		}
	}
}
