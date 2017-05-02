//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
// Unit tests for full document parsing and rendering
//

package blackfriday

import (
	"testing"
)

func runConfluence(input string) string {
	renderer := ConfluenceRenderer(0)
	extensions := 0
	extensions |= EXTENSION_FENCED_CODE
	extensions |= EXTENSION_TABLES
	extensions |= EXTENSION_STRIKETHROUGH
	extensions |= EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK
  return string(Markdown([]byte(input), renderer, extensions))
}

func doTestsConfluence(t *testing.T, tests []string) {
	// catch and report panics
	var candidate string
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("\npanic while processing [%#v]: %s\n", candidate, err)
		}
	}()

	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		candidate = input
		expected := tests[i+1]
		actual := runConfluence(candidate)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}

		// now test every substring to stress test bounds checking
		if !testing.Short() {
			for start := 0; start < len(input); start++ {
				for end := start + 1; end <= len(input); end++ {
					candidate = input[start:end]
					_ = runConfluence(candidate)
				}
			}
		}
	}
}

func TestConfluence(t *testing.T) {
	var tests = []string{
		"# h1",
		"h1. h1\n\n",

		"```go\nfunc foo() bool {\n\treturn true;\n}\n```\n",
		"{code:language=go}func foo() bool {\n\treturn true;\n}\n{code}\n",

		"```\n<div class=\"hoge\">hoge</div>\n```\n",
		"{code}<div class=\"hoge\">hoge</div>\n{code}\n",

		"> foo",
		"{quote}foo\n\n{quote}",

		"---",
		"----",

		"* List\n# Header\n* List\n",
		"* List\n\nh1. Header\n* List\n",

		"* List\n" +
			" * shallow indent\n" +
			"  * part of second list\n" +
			"   * still second\n" +
			"    * almost there\n" +
			"     * third level\n",

		"* List\n* shallow indent\n* part of second list\n* still second\n* almost there\n* third level\n",

		"1. Ting\n\n2. Bong\n3. Goo\n",
		"# Ting\n# Bong\n# Goo\n",

		"||heading 1||heading 2||heading 3||\n|cell A1|cell A2|cell A3|\n|cell B1|cell B2|cell B3|",
		"||heading 1||heading 2||heading 3||\n|cell A1|cell A2|cell A3|\n|cell B1|cell B2|cell B3|\n\n",

		"**strong**",
		"*strong*\n\n",

		"__strong__",
		"*strong*\n\n",

		"~~deleted~~",
		"-deleted-\n\n",

		"![](http://www.host.com/image.gif)",
		"!http://www.host.com/image.gif!\n\n",

		"[Google](https://www.google.co.jp/)",
		"[Google|https://www.google.co.jp/]\n\n",
	}
	doTestsConfluence(t, tests)
}
