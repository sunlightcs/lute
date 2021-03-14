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
	"testing"

	"github.com/sunlightcs/lute"
)

var kramdownBlockIALTests = []parseTest{

	{"22", "* {: id=\"20210221200613-7vpmc8h\"}foo\n  {: id=\"20210221195351-x5tgalq\" updated=\"20210221201411\"}\n{: id=\"20210221195349-czsad7f\" updated=\"20210221195351\"}\n\n\n{: id=\"20210215183533-l36k5mo\" type=\"doc\"}", "<ul id=\"20210221195349-czsad7f\" updated=\"20210221195351\">\n<li id=\"20210221200613-7vpmc8h\">foo</li>\n</ul>\n"},
	{"21", "* {: id=\"fooid\"}foo\n\n  > bar\n  {: id=\"barid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"fooid\">\n<p>foo</p>\n<blockquote id=\"barid\">\n<p>bar</p>\n</blockquote>\n</li>\n</ul>\n"},
	{"20", "* {: id=\"fooid\"} foo\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"fooid\">foo</li>\n</ul>\n"},
	{"19", "* foo\n\n  bar\n  {: id=\"barid\"}\n\n  > baz\n  {: id=\"bazid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n<p id=\"barid\">bar</p>\n<blockquote id=\"bazid\">\n<p>baz</p>\n</blockquote>\n</li>\n</ul>\n"},
	{"18", "* foo\n\n  bar\n  {: id=\"barid\"}\n\n  baz\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n<p id=\"barid\">bar</p>\n<p>baz</p>\n</li>\n</ul>\n"},
	{"17", "> * foo\n>   * bar\n>     * baz\n>\n>       bazz\n>       {: id=\"bazzid\"}\n>     {: id=\"bazid\"}\n>   {: id=\"barid\"}\n> {: id=\"fooid\"}\n{: id=\"id\"}", "<blockquote id=\"id\">\n<ul id=\"fooid\">\n<li id=\"20060102150405-1a2b3c4\">foo\n<ul id=\"barid\">\n<li id=\"20060102150405-1a2b3c4\">bar\n<ul id=\"bazid\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>baz</p>\n<p id=\"bazzid\">bazz</p>\n</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n</blockquote>\n"},
	{"16", "> foo\n> {: id=\"fooid\"}\n> * bar\n> {: id=\"barid\"}\n{: id=\"id\"}", "<blockquote id=\"id\">\n<p id=\"fooid\">foo</p>\n<ul id=\"barid\">\n<li id=\"20060102150405-1a2b3c4\">bar</li>\n</ul>\n</blockquote>\n"},
	{"15", "> foo\n>\n> * bar\n> {: id=\"barid\"}\n{: id=\"id\"}", "<blockquote id=\"id\">\n<p>foo</p>\n<ul id=\"barid\">\n<li id=\"20060102150405-1a2b3c4\">bar</li>\n</ul>\n</blockquote>\n"},
	{"14", "foo\n{: id=\"fooid\"}\nbar\n{: id=\"barid\"}", "<p id=\"fooid\">foo</p>\n<p id=\"barid\">bar</p>\n"},
	{"13", "foo\n{: id=\"fooid\"}\nbar", "<p id=\"fooid\">foo</p>\n<p>bar</p>\n"},
	{"12", "* foo\n\n  > bar\n  {: id=\"bqid\"}\n  > baz\n  > {: id=\"bazid\"}\n* baz\n  {: id=\"bazid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n<blockquote id=\"bqid\">\n<p>bar</p>\n</blockquote>\n<blockquote>\n<p id=\"bazid\">baz</p>\n</blockquote>\n</li>\n<li id=\"20060102150405-1a2b3c4\">\n<p id=\"bazid\">baz</p>\n</li>\n</ul>\n"},
	{"11", "* foo\n  * bar\n  * baz\n  {: id=\"subid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">foo\n<ul id=\"subid\">\n<li id=\"20060102150405-1a2b3c4\">bar</li>\n<li id=\"20060102150405-1a2b3c4\">baz</li>\n</ul>\n</li>\n</ul>\n"},
	{"10", "* foo\n  * bar\n  * baz\n  {: id=\"subid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">foo\n<ul id=\"subid\">\n<li id=\"20060102150405-1a2b3c4\">bar</li>\n<li id=\"20060102150405-1a2b3c4\">baz</li>\n</ul>\n</li>\n</ul>\n"},
	{"9", "* foo\n\n  > bar\n  > {: id=\"barid\"}\n  {: id=\"bqid\"}\n\n  baz\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n<blockquote id=\"bqid\">\n<p id=\"barid\">bar</p>\n</blockquote>\n<p>baz</p>\n</li>\n</ul>\n"},
	{"8", "* foo\n\n  > bar\n  {: id=\"bqid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n<blockquote id=\"bqid\">\n<p>bar</p>\n</blockquote>\n</li>\n</ul>\n"},
	{"7", "* > foo\n  {: id=\"bqid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<blockquote id=\"bqid\">\n<p>foo</p>\n</blockquote>\n</li>\n</ul>\n"},
	{"6", "* foo\n\n* bar\n  {: id=\"barid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">\n<p>foo</p>\n</li>\n<li id=\"20060102150405-1a2b3c4\">\n<p id=\"barid\">bar</p>\n</li>\n</ul>\n"},
	{"5", "* foo\n  {: id=\"fooid\"}\n{: id=\"id\"}", "<ul id=\"id\">\n<li id=\"20060102150405-1a2b3c4\">foo</li>\n</ul>\n"},
	{"4", "* foo\n{: id=\"fooid\"}", "<ul id=\"fooid\">\n<li id=\"20060102150405-1a2b3c4\">foo</li>\n</ul>\n"},
	{"3", "> foo\n> {: id=\"fooid\"}\n>\n> baz\n> {: id=\"bazid\"}\n>\n{: id=\"bqid\"}", "<blockquote id=\"bqid\">\n<p id=\"fooid\">foo</p>\n<p id=\"bazid\">baz</p>\n</blockquote>\n"},
	{"2", "> foo\n> {: id=\"fooid\"}\n{: id=\"bqid\"}", "<blockquote id=\"bqid\">\n<p id=\"fooid\">foo</p>\n</blockquote>\n"},
	{"1", "> foo\n> {: id=\"fooid\" name=\"bar\"}", "<blockquote>\n<p id=\"fooid\" name=\"bar\">foo</p>\n</blockquote>\n"},
	{"0", "foo\n{: id=\"fooid\" class=\"bar\"}", "<p id=\"fooid\" class=\"bar\">foo</p>\n"},
}

func TestKramdownBlockIALs(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.ParseOptions.KramdownBlockIAL = true
	luteEngine.RenderOptions.KramdownBlockIAL = true

	for _, test := range kramdownBlockIALTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}

var kramdownBlockIALIDNAmeTests = []parseTest{

	{"4", "```\nfoo\n```\n{: id=\"20201105103725-3ad5wcz\"}", "<pre data-block-id=\"20201105103725-3ad5wcz\"><code class=\"highlight-chroma\">foo\n</code></pre>\n"},
	{"3", "$$\nfoo\n$$\n{: id=\"20201105103725-3ad5wcz\"}", "<div class=\"language-math\" data-block-id=\"20201105103725-3ad5wcz\">foo</div>\n"},
	{"2", "> foo\n> {: id=\"fooid\" name=\"bar\"}", "<blockquote>\n<p data-block-id=\"fooid\" name=\"bar\">foo</p>\n</blockquote>\n"},
	{"1", "# foo\n{: id=\"fooid\" class=\"bar\"}", "<h1 id=\"foo\" data-block-id=\"fooid\" class=\"bar\">foo</h1>\n"},
	{"0", "# foo", "<h1 id=\"foo\">foo</h1>\n"},
}

func TestKramdownBlockIALIDName(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.ParseOptions.KramdownBlockIAL = true
	luteEngine.RenderOptions.KramdownBlockIAL = true
	luteEngine.RenderOptions.KramdownIALIDRenderName = "data-block-id"

	for _, test := range kramdownBlockIALIDNAmeTests {
		html := luteEngine.MarkdownStr(test.name, test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal markdown text\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
