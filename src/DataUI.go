package main

import (
	"errors"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	_ "github.com/json-iterator/go"
	"log"
	_ "redisManager/redis"
	redisUtil "redisManager/redis"
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

func main() {

	gtk.Init(nil)
	var err error
	imageOK, err = gdk.PixbufNewFromFile("reg.png")
	errorCheck(err)

	// 从glade中获取build对象
	builder, err := gtk.BuilderNewFromFile("redisManager.glade")
	errorCheck(err)

	// 获取id是"mainWindows"的对象
	obj, err := builder.GetObject("mainWindows")

	errorCheck(err)
	win, err := isWindow(obj)

	treeObj, err := builder.GetObject("gtkListTree")
	errorCheck(err)
	treeView, err := isTreeView(treeObj)
	errorCheck(err)
	textObject, err := builder.GetObject("gtkTextView")
	textView, _ = isTextView(textObject)
	buffer, err := textView.GetBuffer()
	// start,end:= buffer.GetBounds()
	// text,err:=buffer.GetText(start,end,true)
	buffer.SetText("asdasldandjkanefonadfasdnfalsd;fna;lsdfnasdlfnasdkjfnas;fdnl")
	//窗体销毁时调用的方法
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	showDB(win, treeView)

}

func showDB(win *gtk.Window, treeView *gtk.TreeView) {
	var iter1, iter2 *gtk.TreeIter
	treeStore, err := gtk.TreeStoreNew(glib.TYPE_OBJECT, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("创建treeview失败:", err)
	}

	treeView.AppendColumn(createImageColumn("图标", COLUMN_ICON))
	treeView.AppendColumn(createTextColumn("内容", COLUMN_TEXT))
	treeView.SetModel(treeStore)
	iter1 = addTreeRow(treeStore, imageOK, "数据库0")

	mp := redisUtil.buildDbStr("139.196.38.232:6379", "adminfeng@.", 0)
	for key, value := range mp {
		if value.value != "" {
			iter2 = addSubRow(treeStore, iter1, imageOK, key)
		} else {
			for i := range value.list {
				addSubRow(treeStore, iter2, imageOK, value.list[i])
			}

		}

	}

	// Add some rows to the tree store

	iter2 = addSubRow(treeStore, iter1, imageOK, "第二层")
	iter2 = addSubRow(treeStore, iter1, imageOK, "这是个有想法的值")
	addSubRow(treeStore, iter2, imageOK, "什么人")
	addSubRow(treeStore, iter2, imageOK, "这是什么情况")
	addSubRow(treeStore, iter2, imageOK, "哈哈哈")
	iter2 = addSubRow(treeStore, iter1, imageOK, "优美的语言")
	iter1 = addTreeRow(treeStore, imageOK, "新的一层")
	iter2 = addSubRow(treeStore, iter1, imageOK, "值")
	iter2 = addSubRow(treeStore, iter1, imageOK, "又是一个值")
	iter2 = addSubRow(treeStore, iter1, imageOK, "还是一个值")
	addSubRow(treeStore, iter2, imageOK, "这个值不会说")
	addSubRow(treeStore, iter2, imageOK, "好说好说")

	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("不能获取选择的对象")
	}
	selection.SetMode(gtk.SELECTION_SINGLE)
	selection.Connect("changed", treeSelectionChangedCB)

	win.ShowAll()
	gtk.Main()
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("类型不是 *gtk.Window")
}

func isTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.TreeView); ok {
		return win, nil
	}
	return nil, errors.New("该类型不是 *gtk.TreeView")
}

func isTextView(obj glib.IObject) (*gtk.TextView, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.TextView); ok {
		return win, nil
	}
	return nil, errors.New("该类型不是 *gtk.TreeView")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// Add a column to the tree view (during the initialization of the tree view)
// We need to distinct the type of data shown in either column.
func createTextColumn(title string, id int) *gtk.TreeViewColumn {
	// In this column we want to show text, hence create a text renderer
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("无法创建text项:", err)
	}

	// Tell the renderer where to pick input from. Text renderer understands
	// the "text" property.
	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("无法创建列：", err)
	}

	return column
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

// Handle selection
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
		valeu, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
		str, _ := valeu.GetString()
		//log.Print(str)
		//log.Printf("treeSelectionChangedCB: 选中的路径是: %s\n", tpath)
		set_text_in_tview(textView, tpath.String()+":"+str)
	}
}
