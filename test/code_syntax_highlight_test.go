// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sunlightcs/lute"
)

func TestCodeSyntaxHighlightIssue17(t *testing.T) {
	// 语法高亮支持内联样式 https://github.com/b3log/lute/issues/17

	caseName := "code-syntax-highlight-issue17"
	data, err := ioutil.ReadFile(caseName + ".md")
	if nil != err {
		t.Fatalf("read case failed: %s", err)
	}

	luteEngine := lute.New()
	luteEngine.SetCodeSyntaxHighlightInlineStyle(true)
	luteEngine.SetCodeSyntaxHighlightLineNum(true)
	style := "monokai"
	luteEngine.SetCodeSyntaxHighlightStyleName(style)
	htmlBytes := luteEngine.Markdown(caseName, data)
	html := string(htmlBytes)
	expected := `<pre style="color: #f8f8f2; background-color: #272822"><code class="language-go"><span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">1</span><span style="color:#f92672">package</span> <span style="color:#a6e22e">main</span>
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">2</span>
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">3</span><span style="color:#f92672">import</span> <span style="color:#e6db74">&#34;fmt&#34;</span>
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">4</span>
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">5</span><span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">6</span>	<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#e6db74">&#34;Hello, 世界&#34;</span>)
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">7</span>}
</code></pre>
`
	if expected != html {
		t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\n", caseName, expected, html)
	}

	data, err = ioutil.ReadFile(caseName + ".tpl")
	if nil != err {
		t.Fatalf("read template failed: %s", err)
	}
	template := string(data)
	template = strings.ReplaceAll(template, "${style}", style)
	template = strings.ReplaceAll(template, "${code}", html)
	ioutil.WriteFile(caseName+".html", []byte(template), 0644)
}

var codeSyntaxHighlightLineNumTests = []parseTest{

	{"0", "```java\nint i;\n```\n", "<pre><code class=\"language-java highlight-chroma\"><span class=\"highlight-ln\">1</span><span class=\"highlight-kt\">int</span> <span class=\"highlight-n\">i</span><span class=\"highlight-o\">;</span>\n</code></pre>\n"},
}

func TestCodeSyntaxHighlightLineNum(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.RenderOptions.CodeSyntaxHighlightLineNum = true

	for _, test := range codeSyntaxHighlightLineNumTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

var codeSyntaxHighlightTests = []parseTest{

	{"0", "```java\nint i;\n```\n", "<pre><code class=\"language-java highlight-chroma\"><span class=\"highlight-kt\">int</span> <span class=\"highlight-n\">i</span><span class=\"highlight-o\">;</span>\n</code></pre>\n"},
}

func TestCodeSyntaxHighlight(t *testing.T) {
	luteEngine := lute.New()

	for _, test := range codeSyntaxHighlightTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

var codeSyntaxHighlightInlineTests = []parseTest{

	{"0", "```java\nint i;\n```\n", "<pre style=\"color: #f8f8f2; background-color: #282a36\"><code class=\"language-java\"><span style=\"color:#8be9fd\">int</span> i<span style=\"color:#ff79c6\">;</span>\n</code></pre>\n"},
}

func TestCodeSyntaxHighlightInline(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.RenderOptions.CodeSyntaxHighlightInlineStyle = true
	luteEngine.RenderOptions.CodeSyntaxHighlightStyleName = "dracula"

	for _, test := range codeSyntaxHighlightInlineTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

var codeSyntaxHighlightStyleTests = []parseTest{

	{"0", "```java\nint i;\n```\n", "<pre><code class=\"language-java highlight-chroma\"><span class=\"highlight-kt\">int</span> <span class=\"highlight-n\">i</span><span class=\"highlight-o\">;</span>\n</code></pre>\n"},
}

func TestCodeSyntaxHighlightStyle(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.RenderOptions.CodeSyntaxHighlightStyleName = "monokai"

	for _, test := range codeSyntaxHighlightStyleTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

var codeSyntaxHighlightOffTests = []parseTest{

	{"0", "```java\nint i;\n```\n", "<pre><code class=\"language-java\">int i;\n</code></pre>\n"},
}

func TestCodeSyntaxHighlightOff(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.RenderOptions.CodeSyntaxHighlight = false

	for _, test := range codeSyntaxHighlightOffTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
