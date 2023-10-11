package parse

import (
	"testing"
)

const testHTML = `
<html>
  <head>
    <title>Sample "Hello, World" Application</title>
  </head>
  <body>
	<h1>Testing</h1>
	<script>
		console.log("test");
	<script>
  </body>
</html>
`

func TestFullText(t *testing.T) {
	//given
	var doc = HTMLParse(testHTML)
	expected := "sample \"hello, world\" application testing"

	//when
	got := doc.FullText()

	//then
	if got != expected {
		t.Errorf("HTMLParse() = %v, want: %v", got, expected)
	}
}
