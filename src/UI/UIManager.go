package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	redisUtil "redisManager/redis"
	"strings"
)

const (
	COLUMN_ICON = iota
	COLUMN_TEXT
)

var (
	imageOK   *gdk.Pixbuf = nil
	imageFAIL *gdk.Pixbuf = nil
)
var textView *gtk.TextView = nil

var data map[string]redisUtil.DataObj

var keyMap map[string]string = make(map[string]string)

func setTextView(view *gtk.TextView) {
	textView = view
}
func setData(redisData map[string]redisUtil.DataObj) {
	data = redisData
}
func setKeyMap(redisKey map[string]string) {
	keyMap = redisKey
}

/**
刷新数据到UI上
*/
func showDB(win *gtk.Window, treeView *gtk.TreeView) {
	imageOK, _ = gdk.PixbufNewFromFile("reg.png")
	var iter1, iter2 *gtk.TreeIter
	treeStore, err := gtk.TreeStoreNew(glib.TYPE_OBJECT, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("创建treeview失败:", err)
	}

	treeView.AppendColumn(createImageColumn("图标", COLUMN_ICON))
	treeView.AppendColumn(createTextColumn("内容", COLUMN_TEXT))
	treeView.SetModel(treeStore)
	iter1 = addTreeRow(treeStore, imageOK, "数据库0")

	data = redisUtil.BuildDbStr("139.196.38.232:6379", "adminfeng@.", 0)
	for key, value := range data {

		if value.Value != "" {
			strs := strings.Split(key, ":")
			fmt.Print("字符串", strs)
			if len(strs) > 1 {
				iter2 = addSubRow(treeStore, iter1, imageOK, strs[0])
				//fmt.Print("当前路径",path)
				for i, v := range strs {
					if v != "" && i > 0 {
						trmp := addSubRow(treeStore, iter2, imageOK, v)

						path, err := treeStore.ToTreeModel().GetPath(trmp)
						if nil != err {
							fmt.Print(err)
						}
						keyMap[path.String()] = key
					}
				}
			} else {
				path, _ := treeStore.ToTreeModel().GetPath(iter2)
				keyMap[path.String()] = key
			}

		} else {
			iter2 = addSubRow(treeStore, iter1, imageOK, key)
			path, _ := treeStore.ToTreeModel().GetPath(iter2)
			keyMap[path.String()] = key
		}

	}
	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("不能获取选择的对象")
	}
	selection.SetMode(gtk.SELECTION_SINGLE)
	selection.Connect("changed", treeSelectionChangedCB)

	win.ShowAll()
	gtk.Main()
}

func createImageColumn(title string, id int) *gtk.TreeViewColumn {

	cellRenderer, err := gtk.CellRendererPixbufNew()
	if err != nil {
		log.Fatal("无法创建pixbuf:", err)
	}
	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "pixbuf", id)
	if err != nil {
		log.Fatal("无法创建列:", err)
	}

	return column
}

func createTextColumn(title string, id int) *gtk.TreeViewColumn {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("无法创建text项:", err)
	}

	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("无法创建列：", err)
	}

	return column
}

//添加顶层的行
func addTreeRow(treeStore *gtk.TreeStore, icon *gdk.Pixbuf, text string) *gtk.TreeIter {
	return addSubRow(treeStore, nil, icon, text)
}

// 添加子行
func addSubRow(treeStore *gtk.TreeStore, iter *gtk.TreeIter, icon *gdk.Pixbuf, text string) *gtk.TreeIter {
	// 末尾的新行添加迭代器
	i := treeStore.Append(iter)

	//设置迭代器表示的树存储行的内容
	err := treeStore.SetValue(i, COLUMN_ICON, icon)
	if err != nil {
		log.Fatal("未能设置值:", err)
	}
	err = treeStore.SetValue(i, COLUMN_TEXT, text)
	if err != nil {
		log.Fatal("未能设置值:", err)
	}
	return i
}
func set_text_in_tview(tv *gtk.TextView, text string) {
	buffer := get_buffer_from_tview(tv)
	buffer.SetText(text)
}

func get_buffer_from_tview(tv *gtk.TextView) *gtk.TextBuffer {
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}
	return buffer
}

func treeSelectionChangedCB(selection *gtk.TreeSelection) {
	var iter *gtk.TreeIter
	var model gtk.ITreeModel
	var ok bool
	model, iter, ok = selection.GetSelected()
	if ok {

		tpath, err := model.(*gtk.TreeModel).GetPath(iter)
		if err != nil {
			log.Printf("treeSelectionChangedCB:无法获取路径: %s\n", err)
			return
		}
		//fmt.Print("----------------------->",tpath)
		kstr := keyMap[tpath.String()]
		//valeu, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
		model.(*gtk.TreeModel).GetPath(iter)
		//str, _ := valeu.GetString()
		//log.Print(str)
		//log.Printf("treeSelectionChangedCB: 选中的路径是: %s\n", tpath)
		value := data[kstr]
		if value.Value != "" {
			set_text_in_tview(textView, value.Key+":"+value.Value)
		} else {
			list := data[kstr].List
			str := ""
			for i := range list {
				str += "," + list[i]
			}
			set_text_in_tview(textView, str)
		}

	}
}
