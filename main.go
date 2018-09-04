package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	log "github.com/sirupsen/logrus"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	log.Info("Starting App")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	app := widgets.NewQApplication(len(os.Args), os.Args)

	func() {
		var file = core.NewQFile2(":qml/style.qss")
		if file.Open(core.QIODevice__ReadOnly) {
			qdata := file.ReadAll()
			datastr := qdata.ConstData()
			app.SetStyleSheet(datastr)
		}
	}()

	window := widgets.NewQMainWindow(nil, 0)
	window.SetFixedSize2(250, 200)
	window.SetWindowTitle("My Qt App with changes")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	button := widgets.NewQPushButton2("all now done!", nil)
	button.ConnectClicked(func(bool) {
		button.SetDisabled(true)
	})
	widget.Layout().AddWidget(button)

	window.Show()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	app.Exec()
}
