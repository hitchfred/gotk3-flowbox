package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatalf("error creating main window: %q\n", err)
	}
	win.SetTitle("Flow Box Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	fb, err := gtk.FlowBoxNew()
	if err != nil {
		log.Fatalf("error creating FlowBox: %q\n", err)
	}
	fb.SetMaxChildrenPerLine(5)

	for i := 0; i < 500; i++ {
		lbl, err := gtk.LabelNew(fmt.Sprintf("Label %d", i+1))
		if err != nil {
			log.Fatalf("error creating label: %q\n", err)
		}
		fb.Insert(lbl, -1)
	}

	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	scroll, _ := gtk.ScrolledWindowNew(nil, nil)
	scroll.Add(fb)
	box.PackStart(scroll, true, true, 1)

	win.Add(box)
	win.ShowAll()
	gtk.Main()
}
