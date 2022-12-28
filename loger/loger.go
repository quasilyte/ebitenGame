package loger

import "github.com/sirupsen/logrus"

var (
	L = logrus.New()
)

func ConfigLoger() {
	L.Level = 6
	//Log.Trace("6")
	//Log.Debug("5")
	//Log.Info("4")
	//Log.Warning("3")
	//Log.Error("2")
	//Log.Fatal("1")
	//Log.Panic("0")
}
