// SiYuan - Build Your Eternal Digital Garden
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package conf

import (
	"github.com/siyuan-note/siyuan/kernel/util"
)

type FileTree struct {
	AlwaysSelectOpenedFile bool   `json:"alwaysSelectOpenedFile"` // 是否自动选中当前打开的文件
	OpenFilesUseCurrentTab bool   `json:"openFilesUseCurrentTab"` // 在当前页签打开文件
	RefCreateSavePath      string `json:"refCreateSavePath"`      // 块引时新建文档存储文件夹路径
	CreateDocNameTemplate  string `json:"createDocNameTemplate"`  // 新建文档名模板
	MaxListCount           int    `json:"maxListCount"`           // 最大列出数量
	AllowCreateDeeper      bool   `json:"allowCreateDeeper"`      // 允许创建超过 7 层深度的子文档

	Sort int `json:"sort"` // 排序方式
}

func NewFileTree() *FileTree {
	return &FileTree{
		AlwaysSelectOpenedFile: false,
		OpenFilesUseCurrentTab: false,
		Sort:                   util.SortModeCustom,
		CreateDocNameTemplate:  "",
		MaxListCount:           512,
		AllowCreateDeeper:      false,
	}
}
