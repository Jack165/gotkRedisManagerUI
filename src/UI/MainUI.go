package main

import (
	"errors"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	redisUtil "redisManager/redis"
	"strconv"
)

const appId = "com.feng.RedisManager"

var rdb *redis.Client
var urlText *gtk.Entry
var pwdText *gtk.Entry
var comboBox *gtk.ComboBoxText
var win *gtk.ApplicationWindow

var connet link

type link struct {
	name string
	url  string
	pwd  string
}

func main() {
	showMain()
	//showDB()
}

func showMain() {
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	imageOK, err = gdk.PixbufNewFromFile("Resource/redis.jfif")
	errorCheck(err)
	application.Connect("startup", func() {

		log.Println("application startup")
	})
	application.Connect("activate", func() {
		log.Println("application activate")

		builder, err := gtk.BuilderNewFromFile("Resource/selectDB.glade")
		errorCheck(err)

		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("gtkAppWindow")
		errorCheck(err)
		win, _ = isWindow(obj)
		btnLoginObj, err := builder.GetObject("btnLogin")
		errorCheck(err)
		btn, err := isButton(btnLoginObj)
		btn.Connect("clicked", loginBtnClicked)
		btnDetailObj, err := builder.GetObject("btnDetail")
		errorCheck(err)
		btnShowDb, err := isButton(btnDetailObj)
		btnShowDb.Connect("clicked", showDbDetail)

		errorCheck(err)
		txtUrl, err := builder.GetObject("entryUrl")
		urlText, _ = isEntry(txtUrl)
		urlText.SetText("139.196.38.232:6379")
		txtPwd, err := builder.GetObject("entryPwd")
		pwdText, _ = isEntry(txtPwd)
		pwdText.SetVisibility(false)
		pwdText.SetInvisibleChar('*')
		pwdText.SetText("adminfeng@.")
		gtkComboBoxObj, err := builder.GetObject("checkDb")
		errorCheck(err)
		comboBox, err = isComboBox(gtkComboBoxObj)
		errorCheck(err)
		//rdb := redisUtil.GetRedisDb("139.196.38.232:6379", "adminfeng@.", 0)
		rdb := redisUtil.GetRedisDb(connet.url, connet.pwd, comboBox.GetActive())
		redisClient = rdb
		win.SetTitle("RedisManager")
		win.SetIcon(imageOK)

		win.Show()
		application.AddWindow(win)
	})
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	os.Exit(application.Run(os.Args))
}

func showDB() {

	db, _ := strconv.Atoi(comboBox.GetActiveText())
	rdb := redisUtil.GetRedisDb(connet.url, connet.pwd, db)
	redisClient = rdb
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	imageOK, err = gdk.PixbufNewFromFile("Resource/redis.jfif")

	application.Connect("startup", func() {
		log.Println("application startup")
	})

	application.Connect("activate", func() {
		log.Println("application activate")

		builder, err := gtk.BuilderNewFromFile("Resource/newRedis.glade")
		errorCheck(err)

		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("mainWindows")
		detailWin, err := isWindow(obj)
		errorCheck(err)
		treeObj, err := builder.GetObject("gtkTreeView")
		errorCheck(err)
		treeView, err := isTreeView(treeObj)
		errorCheck(err)
		textObject, err := builder.GetObject("gtkTextView")
		textView, _ = isTextView(textObject)
		detailWin.SetTitle("RedisManager-数据库" + comboBox.GetActiveText())
		detailWin.SetIcon(imageOK)
		detailWin.Show()
		application.AddWindow(win)

		keys := redisUtil.KeyList(rdb)
		flushKeys(treeView, keys)

	})

	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})
	application.Run(os.Args)
	//os.Exit()
}

func loginBtnClicked() {
	url, _ := urlText.GetText()
	pwd, _ := pwdText.GetText()
	connet.url = url
	connet.pwd = pwd

	rdb = redisUtil.GetRedisDb(url, pwd, 0)
	dbSize := redisUtil.GetDbSize(rdb)
	var texts = make([]string, dbSize)
	comboBox.RemoveAll()
	for i := 0; i < dbSize; i++ {
		texts[i] = strconv.Itoa(i)
	}
	boxCom(texts)
}

func showDbDetail() {
	showDB()
}

func boxCom(text []string) {
	for i := 0; i < len(text); i++ {
		num := text[i]
		comboBox.AppendText(num)
	}
	comboBox.SetActive(0)

}

func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
func errorCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func isWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}
func isTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	if win, ok := obj.(*gtk.TreeView); ok {
		return win, nil
	}
	return nil, errors.New("该类型不是 *gtk.TreeView")
}

func isTextView(obj glib.IObject) (*gtk.TextView, error) {
	if win, ok := obj.(*gtk.TextView); ok {
		return win, nil
	}
	return nil, errors.New("该类型不是 *gtk.TreeView")
}

func isButton(obj glib.IObject) (*gtk.Button, error) {
	if button, ok := obj.(*gtk.Button); ok {
		return button, nil
	}
	return nil, errors.New("该类型不是 *gtk.Button")
}
func isEntry(obj glib.IObject) (*gtk.Entry, error) {
	if entry, ok := obj.(*gtk.Entry); ok {
		return entry, nil
	}
	return nil, errors.New("该类型不是 *gtk.TreeView")
}
func isComboBox(obj glib.IObject) (*gtk.ComboBoxText, error) {
	if entry, ok := obj.(*gtk.ComboBoxText); ok {
		return entry, nil
	}
	return nil, errors.New("该类型不是 *gtk.ComboBox")
}
