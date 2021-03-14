// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package parse

import (
	"bytes"
	"github.com/sunlightcs/lute/ast"
	"github.com/sunlightcs/lute/util"
	"strings"
)

// 判断 kramdown 块级内联属性列表（{: attrs}）是否开始。
func IALStart(t *Tree, container *ast.Node) int {
	if !t.Context.ParseOption.KramdownBlockIAL || t.Context.indented {
		return 0
	}

	if ast.NodeListItem == t.Context.Tip.Type && nil == t.Context.Tip.FirstChild { // 在列表最终化过程中处理
		return 0
	}

	if ial := t.parseKramdownBlockIAL(); nil != ial {
		t.Context.closeUnmatchedBlocks()
		t.Context.offset = t.Context.currentLineLen                    // 整行过
		if 1 < len(ial) && "type" == ial[1][0] && "doc" == ial[1][1] { // 文档块 IAL
			t.Context.rootIAL = &ast.Node{Type: ast.NodeKramdownBlockIAL, Tokens: t.Context.currentLine[t.Context.nextNonspace:]}
			t.Root.KramdownIAL = ial
			t.Root.ID = ial[0][1]
			t.ID = t.Root.ID
			return 2
		}

		lastMatchedContainer := t.Context.lastMatchedContainer
		if t.Context.allClosed {
			if ast.NodeDocument == lastMatchedContainer.Type || ast.NodeListItem == lastMatchedContainer.Type || ast.NodeBlockquote == lastMatchedContainer.Type {
				lastMatchedContainer = t.Context.Tip.LastChild // 挂到最后一个子块上
				if nil == lastMatchedContainer {
					lastMatchedContainer = t.Context.lastMatchedContainer
				}
				if ast.NodeKramdownBlockIAL == lastMatchedContainer.Type && nil != lastMatchedContainer.Parent && ast.NodeDocument == lastMatchedContainer.Parent.Type { // 两个连续的 IAL
					tokens := IAL2Tokens(ial)
					if !bytes.HasPrefix(lastMatchedContainer.Tokens, tokens) { // 有的块解析已经做过打断处理
						p := &ast.Node{Type: ast.NodeParagraph, Tokens: []byte(" ")}
						lastMatchedContainer.InsertAfter(p)
						t.Context.Tip = p
						lastMatchedContainer = p
					}
				}
			}
		}
		lastMatchedContainer.KramdownIAL = ial
		lastMatchedContainer.ID = ial[0][1]
		node := t.Context.addChild(ast.NodeKramdownBlockIAL)
		node.Tokens = t.Context.currentLine[t.Context.nextNonspace:]
		return 2
	}
	return 0
}

var openCurlyBraceColon = util.StrToBytes("{: ")
var emptyIAL = util.StrToBytes("{:}")

func IAL2Tokens(ial [][]string) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("{: ")
	for i, kv := range ial {
		buf.WriteString(kv[0])
		buf.WriteString("=\"")
		buf.WriteString(kv[1])
		buf.WriteByte('"')
		if i < len(ial)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.Bytes()
}

func Tokens2IAL(tokens []byte) (ret [][]string) {
	// tokens 开头必须是空格
	tokens = bytes.TrimRight(tokens, " \n")
	tokens = bytes.TrimPrefix(tokens, []byte("{:"))
	tokens = bytes.TrimSuffix(tokens, []byte("}"))
	for {
		valid, remains, attr, name, val := TagAttr(tokens)
		if !valid {
			break
		}

		tokens = remains
		if 1 > len(attr) {
			break
		}

		ret = append(ret, []string{util.BytesToStr(name), util.BytesToStr(val)})
	}
	return
}

func (t *Tree) parseKramdownBlockIAL() (ret [][]string) {
	tokens := t.Context.currentLine[t.Context.nextNonspace:]
	return t.Context.parseKramdownBlockIAL(tokens)
}

func (t *Tree) parseKramdownSpanIAL() {
	ast.Walk(t.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.WalkContinue
		}

		switch n.Type {
		case ast.NodeEmphasis, ast.NodeStrong, ast.NodeCodeSpan, ast.NodeStrikethrough, ast.NodeTag, ast.NodeMark, ast.NodeImage:
			break
		default:
			return ast.WalkContinue
		}

		if nil == n.Next || ast.NodeText != n.Next.Type {
			return ast.WalkContinue
		}

		tokens := n.Next.Tokens
		if pos, ial := t.Context.parseKramdownSpanIAL(tokens); 0 < len(ial) {
			n.KramdownIAL = ial
			n.Next.Tokens = tokens[pos+1:]
			if 1 > len(n.Next.Tokens) {
				n.Next.Unlink() // 移掉空的文本节点 {: ial}
			}
			spanIAL := &ast.Node{Type: ast.NodeKramdownSpanIAL, Tokens: tokens[:pos+1]}
			n.InsertAfter(spanIAL)
		}
		return ast.WalkContinue
	})
	return
}

func (context *Context) parseKramdownBlockIAL(tokens []byte) (ret [][]string) {
	if curlyBracesStart := bytes.Index(tokens, []byte("{:")); 0 == curlyBracesStart {
		tokens = tokens[curlyBracesStart+2:]
		curlyBracesEnd := bytes.Index(tokens, closeCurlyBrace)
		if 3 > curlyBracesEnd {
			return
		}

		if !bytes.Equal(tokens[curlyBracesEnd:], []byte("}\n")) { // IAL 后不能存在其他内容，必须独占一行
			return
		}
		ret = Tokens2IAL(tokens)
	}
	return
}

func (context *Context) parseKramdownSpanIAL(tokens []byte) (pos int, ret [][]string) {
	pos = bytes.Index(tokens, closeCurlyBrace)
	if curlyBracesStart := bytes.Index(tokens, []byte("{:")); 0 == curlyBracesStart && curlyBracesStart+2 < pos {
		tokens = tokens[curlyBracesStart+2:]
		curlyBracesEnd := bytes.Index(tokens, closeCurlyBrace)
		if 3 > curlyBracesEnd {
			return
		}

		tokens = tokens[:curlyBracesEnd]
		for {
			valid, remains, attr, name, val := TagAttr(tokens)
			if !valid {
				break
			}

			tokens = remains
			if 1 > len(attr) {
				break
			}

			nameStr := strings.ReplaceAll(util.BytesToStr(name), util.Caret, "")
			valStr := strings.ReplaceAll(util.BytesToStr(val), util.Caret, "")
			ret = append(ret, []string{nameStr, valStr})
		}
	}
	return
}

func (context *Context) parseKramdownIALInListItem(tokens []byte) (ret [][]string) {
	if curlyBracesStart := bytes.Index(tokens, []byte("{:")); 0 == curlyBracesStart {
		tokens = tokens[curlyBracesStart+2:]
		curlyBracesEnd := bytes.Index(tokens, closeCurlyBrace)
		if 3 > curlyBracesEnd {
			return
		}

		tokens = tokens[:bytes.Index(tokens, []byte("}"))]
		for {
			valid, remains, attr, name, val := TagAttr(tokens)
			if !valid {
				break
			}

			tokens = remains
			if 1 > len(attr) {
				break
			}

			ret = append(ret, []string{util.BytesToStr(name), util.BytesToStr(val)})
		}
	}
	return
}
