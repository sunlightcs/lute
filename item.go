// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

// +build !srcmap

package lute

// item 描述了词法分析的一个 token。
type item byte

// items 定义了 token 数组。
type items []item

// nilItem 返回一个空值 token。
func nilItem() item {
	return item(0)
}

// isNilItem 判断 item 是否为空值。
func isNilItem(item item) bool {
	return 0 == item
}

// newItem 构造一个 token。
func newItem(term byte, ln, col int) item {
	return item(term)
}

// term 返回 item 的词素。
func term(item item) byte {
	return byte(item)
}

func setTerm(tokens *items, i int, term byte) {
	(*tokens)[i] = item(term)
}

// strToItems 将 str 转为 items。
func strToItems(str string) (ret items) {
	length := len(str)
	ret = make(items, 0, 256)
	for i := 0; i < length; i++ {
		ret = append(ret, item(str[i]))
	}
	return
}

// itemsToStr 将 items 转为 string。
func itemsToStr(items items) string {
	return string(itemsToBytes(items))
}

// itemsToBytes 将 items 转为 []byte。
func itemsToBytes(items items) (ret []byte) {
	length := len(items)
	ret = make([]byte, 0, 256)
	for i := 0; i < length; i++ {
		ret = append(ret, byte(items[i]))
	}
	return
}

// bytesToItems 将 bytes 转为 items。
func bytesToItems(bytes []byte) (ret items) {
	length := len(bytes)
	ret = make(items, 0, 256)
	for i := 0; i < length; i++ {
		ret = append(ret, item(bytes[i]))
	}
	return
}
