package tokenizer

import "testing"

const testHTML = `
<html>
  <head>
    <title>Sample "Hello, World" Application</title>
	<style>
	  body {
		width: 100%;	
	  }
	</style>
  </head>
  <body>
	<h1>Testing</h1>
	<script>
		console.log("test");
	<script>
  </body>
</html>
`

func TestParseHTML(t *testing.T) {
	// given
	expected := "sample \"hello, world\" application testing "
	// when
	got := ParseHTML(testHTML)
	// then
	if got != expected {
		t.Errorf("HTMLParse() = %v, want: %v", got, expected)
	}
}
