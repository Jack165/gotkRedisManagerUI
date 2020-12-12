package main

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
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
var keyTreeView *gtk.TreeView
var redisClient *redis.Client

var data map[string]redisUtil.DataObj

var keyMap = make(map[string]string)

/**
重新加载所以的key列表
*/
func flushKeys(treeView *gtk.TreeView, keys []string) {
	keyTreeView = treeView
	imageOK, _ = gdk.PixbufNewFromFile("Resource/reg.png")
	keyTreeView.AppendColumn(createImageColumn("图标", COLUMN_ICON))
	keyTreeView.AppendColumn(createTextColumn("内容", COLUMN_TEXT))
	treeStore, err := gtk.TreeStoreNew(glib.TYPE_OBJECT, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("创建treeView失败:", err)
	}
	keyTreeView.SetModel(treeStore)
	iter1 := addTreeRow(treeStore, imageOK, "数据库")
	for key, _ := range keys {
		strs := strings.Split(keys[key], ":")
		appendKeyTree(strs, keys[key], treeStore, iter1)
	}
	//keyTreeView.ExpandAll()

	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("不能获取选择的对象")
	}
	selection.SetMode(gtk.SELECTION_SINGLE)
	selection.Connect("changed", showValue)

}

func appendKeyTree(keys []string, redisKey string, treeStore *gtk.TreeStore, iter *gtk.TreeIter) {
	if len(keys) > 2 {
		iter = addSubRow(treeStore, iter, imageOK, keys[0])
		appendKeyTree(keys[1:len(keys)], redisKey, treeStore, iter)
	} else {
		var ter2 *gtk.TreeIter
		if len(keys) > 1 {
			ter2 = addSubRow(treeStore, iter, imageOK, keys[0]+":"+keys[1])
		} else {
			ter2 = addSubRow(treeStore, iter, imageOK, keys[0])
		}
		path, _ := treeStore.ToTreeModel().GetPath(ter2)
		keyMap[path.String()] = redisKey
	}
}

//创建treeView泰坦
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

	//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GBK.NewEncoder()))
	//bs:=[]byte(text)

	buffer.SetText(text)

}

func get_buffer_from_tview(tv *gtk.TextView) *gtk.TextBuffer {
	buffer, err := tv.GetBuffer()

	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}
	return buffer
}

func showValue(selection *gtk.TreeSelection) {
	var ok bool
	var iter *gtk.TreeIter
	var model gtk.ITreeModel
	model, iter, ok = selection.GetSelected()
	if ok {

		tpath, err := model.(*gtk.TreeModel).GetPath(iter)
		if err != nil {
			log.Printf("treeSelectionChangedCB:无法获取路径: %s\n", err)
			return
		}
		redisKey := keyMap[tpath.String()]
		keyTreeView.ExpandRow(tpath, false)
		if redisKey != "" {
			text := redisUtil.GetRedisValue(redisKey, redisClient)
			set_text_in_tview(textView, text)
		}

	}
}
