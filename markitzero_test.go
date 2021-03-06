package main

import (
	"strings"
	"testing"
)

func doonezerotest(t *testing.T, input, output string) {
	result := markitzero(input)
	if result != output {
		t.Errorf("\nexpected:\n%s\noutput:\n%s", output, result)
	}
}

func TestBasictest(t *testing.T) {
	input := `link to https://example.com/ with **bold** text`
	output := `link to <a href="https://example.com/">https://example.com/</a> with <b>bold</b> text`
	doonezerotest(t, input, output)
}

func TestMultibold(t *testing.T) {
	input := `**in** out **in**`
	output := `<b>in</b> out <b>in</b>`
	doonezerotest(t, input, output)
}

func TestLinebreak1(t *testing.T) {
	input := "hello\n> a quote\na comment"
	output := "hello<blockquote>a quote</blockquote><p>a comment"
	doonezerotest(t, input, output)
}

func TestLinebreak2(t *testing.T) {
	input := "hello\n\n> a quote\n\na comment"
	output := "hello<br><blockquote>a quote</blockquote><p>a comment"
	doonezerotest(t, input, output)
}

func TestLinebreak3(t *testing.T) {
	input := "hello\n\n```\nfunc(s string)\n```\n\ndoes it go?"
	output := "hello<br><pre><code>func(s string)</code></pre><p>does it go?"
	doonezerotest(t, input, output)
}

func TestCodeStyles(t *testing.T) {
	input := "hello\n\n```go\nfunc(s string)\n```\n\ndoes it go?"
	output := "hello<br><pre><code><span class=kw>func</span><span class=op>(</span>s <span class=tp>string</span><span class=op>)</span></code></pre><p>does it go?"
	doonezerotest(t, input, output)
}

func TestSimplelink(t *testing.T) {
	input := "This is a [link](https://example.com)."
	output := `This is a <a href="https://example.com">link</a>.`
	doonezerotest(t, input, output)
}

func TestSimplelink2(t *testing.T) {
	input := "See (http://example.com) for examples."
	output := `See (<a href="http://example.com">http://example.com</a>) for examples.`
	doonezerotest(t, input, output)
}

func TestWikilink(t *testing.T) {
	input := "I watched [Hackers](https://en.wikipedia.org/wiki/Hackers_(film))"
	output := `I watched <a href="https://en.wikipedia.org/wiki/Hackers_(film)">Hackers</a>`
	doonezerotest(t, input, output)
}

func TestQuotedlink(t *testing.T) {
	input := `quoted "https://example.com/link" here`
	output := `quoted "<a href="https://example.com/link">https://example.com/link</a>" here`
	doonezerotest(t, input, output)
}

func TestLinkinQuote(t *testing.T) {
	input := `> a quote and https://example.com/link`
	output := `<blockquote>a quote and <a href="https://example.com/link">https://example.com/link</a></blockquote><p>`
	doonezerotest(t, input, output)
}

func TestBoldLink(t *testing.T) {
	input := `**b https://example.com/link b**`
	output := `<b>b <a href="https://example.com/link">https://example.com/link</a> b</b>`
	doonezerotest(t, input, output)
}

func TestHonklink(t *testing.T) {
	input := `https://en.wikipedia.org/wiki/Honk!`
	output := `<a href="https://en.wikipedia.org/wiki/Honk!">https://en.wikipedia.org/wiki/Honk!</a>`
	doonezerotest(t, input, output)
}

func TestImagelink(t *testing.T) {
	input := `an image <img alt="caption" src="https://example.com/wherever"> and linked [<img src="there">](example.com)`
	output := `an image <img alt="caption" src="https://example.com/wherever"> and linked <a href="example.com"><img src="there"></a>`
	doonezerotest(t, input, output)
}

func TestLists(t *testing.T) {
	input := `hello
+ a list
+ the list

para

- singleton`
	output := `hello<ul><li>a list<li>the list</ul><p>para<ul><li>singleton</ul><p>`
	doonezerotest(t, input, output)
}

func TestTables(t *testing.T) {
	input := `hello

| col1 | col 2 |
| row2 | cell4 |

para
`
	output := `hello<table><tr><td>col1<td>col 2<tr><td>row2<td>cell4</table><p>para`
	doonezerotest(t, input, output)
}

var benchData, simpleData string

func init() {
	benchData = strings.Repeat("hello there sir. It is a **fine** day for some testing!\n", 100)
	simpleData = strings.Repeat("just a few words\n", 100)
}

func BenchmarkModerateSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		markitzero(benchData)
	}
}

func BenchmarkSimpleData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		markitzero(simpleData)
	}
}
