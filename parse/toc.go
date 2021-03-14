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

func (context *Context) parseToC(paragraph *ast.Node) *ast.Node {
	lines := lex.Split(paragraph.Tokens, lex.ItemNewline)
	if 1 != len(lines) {
		return nil
	}

	content := bytes.TrimSpace(lines[0])
	if context.ParseOption.VditorWYSIWYG || context.ParseOption.VditorIR || context.ParseOption.VditorSV {
		content = bytes.ReplaceAll(content, util.CaretTokens, nil)
	}
	if !bytes.EqualFold(content, []byte("[toc]")) {
		return nil
	}
	return &ast.Node{Type: ast.NodeToC}
}
