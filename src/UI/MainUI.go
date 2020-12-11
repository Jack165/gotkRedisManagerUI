package main

import (
	"errors"
	_ "github.com/go-redis/redis/v8"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	redisUtil "redisManager/redis"
)

const appId = "com.feng.RedisManager"

func main() {
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	imageOK, err = gdk.PixbufNewFromFile("redis.jfif")
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	application.Connect("activate", func() {
		log.Println("application activate")

		builder, err := gtk.BuilderNewFromFile("UI/newRedis.glade")
		errorCheck(err)

		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("mainWindows")
		errorCheck(err)
		treeObj, err := builder.GetObject("gtkTreeView")
		errorCheck(err)
		treeView, err := isTreeView(treeObj)
		win, err := isWindow(obj)
		errorCheck(err)
		textObject, err := builder.GetObject("gtkTextView")
		textView, _ = isTextView(textObject)
		win.SetTitle("RedisManager")
		win.SetIcon(imageOK)
		win.Show()
		application.AddWindow(win)
		rdb := redisUtil.GetRedisDb("139.196.38.232:6379", "adminfeng@.", 0)
		redisClient = rdb
		keys := redisUtil.KeyList(rdb)
		flushKeys(treeView, keys)
		//showDB(win,treeView)
	})

	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	os.Exit(application.Run(os.Args))
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
