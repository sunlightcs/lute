// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"io/ioutil"
	"os"
	"runtime/pprof"

	"github.com/sunlightcs/lute"
)

func main() {
	spec := "test/commonmark-spec"
	bytes, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		panic(err)
	}

	luteEngine := lute.New()
	luteEngine.SetGFMTaskListItem(false)
	luteEngine.SetGFMTable(false)
	luteEngine.SetGFMAutoLink(false)
	luteEngine.SetGFMStrikethrough(false)
	luteEngine.SetSoftBreak2HardBreak(false)
	luteEngine.SetCodeSyntaxHighlight(false)
	luteEngine.SetFootnotes(false)
	luteEngine.SetAutoSpace(false)
	luteEngine.SetFixTermTypo(false)
	luteEngine.SetEmoji(false)
	luteEngine.SetBlockRef(false)
	luteEngine.SetMark(false)

	cpuProfile, _ := os.Create("pprof/cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	for i := 0; i < 1024; i++ {
		luteEngine.Markdown("pprof "+spec, bytes)
	}
	pprof.StopCPUProfile()
}
