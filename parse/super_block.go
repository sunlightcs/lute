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
	"github.com/sunlightcs/lute/lex"
	"github.com/sunlightcs/lute/util"
)

// 判断超级块（{{{ blocks }}}）是否开始。
func SuperBlockStart(t *Tree, container *ast.Node) int {
	if !t.Context.ParseOption.SuperBlock || t.Context.indented {
		return 0
	}

	if ok, layout := t.parseSuperBlock(); ok {
		t.Context.closeUnmatchedBlocks()
		t.Context.addChild(ast.NodeSuperBlock)
		t.Context.addChildMarker(ast.NodeSuperBlockOpenMarker, nil)
		t.Context.addChildMarker(ast.NodeSuperBlockLayoutMarker, layout)
		t.Context.offset = t.Context.currentLineLen - 1 // 整行过
		return 1
	}
	return 0
}

func SuperBlockContinue(superBlock *ast.Node, context *Context) int {
	if context.isSuperBlockClose(context.currentLine[context.nextNonspace:]) {
		level := 0
		for p := context.Tip; nil != p; p = p.Parent {
			if ast.NodeSuperBlock == p.Type {
				level++
			}
		}
		if 1 < level {
			return 4 // 嵌套层闭合
		}
		return 3 // 顶层闭合
	}
	return 0
}

func (context *Context) superBlockFinalize(superBlock *ast.Node) {
	// 最终化所有子块
	for child := superBlock.FirstChild; nil != child; child = child.Next {
		if child.Close {
			continue
		}
		context.finalize(child)
	}
}

func (t *Tree) parseSuperBlock() (ok bool, layout []byte) {
	marker := t.Context.currentLine[t.Context.nextNonspace]
	if lex.ItemOpenBrace != marker {
		return
	}

	fenceChar := marker
	var fenceLen int
	for i := t.Context.nextNonspace; i < t.Context.currentLineLen && fenceChar == t.Context.currentLine[i]; i++ {
		fenceLen++
	}

	if 3 != fenceLen {
		return
	}

	layout = t.Context.currentLine[t.Context.nextNonspace+fenceLen:]
	layout = lex.TrimWhitespace(layout)
	if !bytes.EqualFold(layout, nil) && !bytes.EqualFold(layout, []byte("row")) && !bytes.EqualFold(layout, []byte("col")) {
		return
	}
	return true, layout
}

func (context *Context) isSuperBlockClose(tokens []byte) (ok bool) {
	tokens = lex.TrimWhitespace(tokens)
	if bytes.Equal(tokens, []byte(util.Caret+"}}}")) {
		p := &ast.Node{Type: ast.NodeParagraph, Tokens: util.CaretTokens}
		context.TipAppendChild(p)
	}
	endCaret := bytes.HasSuffix(tokens, util.CaretTokens)
	tokens = bytes.ReplaceAll(tokens, util.CaretTokens, nil)
	if !bytes.Equal([]byte("}}}"), tokens) {
		return
	}
	if endCaret {
		paras := context.Tip.ChildrenByType(ast.NodeParagraph)
		if length := len(paras); 0 < length {
			lastP := paras[length-1]
			lastP.Tokens = append(lastP.Tokens, util.CaretTokens...)
		}
	}
	return true
}
