package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(&os.Args)

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
	for i := 0; i < 500; i++ {
		lbl, err := gtk.LabelNew(fmt.Sprintf("Label %d", i+1))
		if err != nil {
			log.Fatalf("error creating label: %q\n", err)
		}
		fb.Insert(lbl, -1)
	}

	text, _ := gtk.TextViewNew()
	buffer, _ := text.GetBuffer()
	text.SetEditable(false)

	fb.SetActivateOnSingleClick(false)
	fb.SetMaxChildrenPerLine(5)
	fb.SetSelectionMode(gtk.SELECTION_MULTIPLE)
	fb.Connect("selected-children-changed", func() {
		buffer.SetText("")
		for _, child := range fb.GetSelectedChildren() {
			// The child widget of the FlowBoxChild is one of the labels we added earlier
			w, _ := child.GetChild()
			lbl, _ := w.GetProperty("label")
			s := fmt.Sprintf("child selected: %p: index %d, label %q\n", child, child.GetIndex(), lbl)
			buffer.InsertAtCursor(s)
		}
	})
	fb.Connect("child-activated", func(fb *gtk.FlowBox, child *gtk.FlowBoxChild) {
		w, _ := child.GetChild()
		lbl, _ := w.GetProperty("label")
		s := fmt.Sprintf("child activated: %p: index %d, label %q\n", child, child.GetIndex(), lbl)
		buffer.SetText(s)
	})
	paned, _ := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)
	scrollFlow, _ := gtk.ScrolledWindowNew(nil, nil)
	scrollFlow.Add(fb)
	scrollLog, _ := gtk.ScrolledWindowNew(nil, nil)
	scrollLog.Add(text)
	paned.Pack1(scrollFlow, true, true)
	paned.Pack2(scrollLog, true, true)

	win.Add(paned)
	win.SetDefaultSize(640, 480)
	win.ShowAll()
	gtk.Main()
}
