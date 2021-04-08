package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func updateLog(p *widgets.Paragraph, text string) {
	w, h := ui.TerminalDimensions()
	p.Text = text
	p.PaddingLeft = (w - len(text)) / 2
	p.SetRect(0, h-3, w, h)

	ui.Render(p)
}

func updateTitle(p *widgets.Paragraph, text string) {
	w, _ := ui.TerminalDimensions()
	p.Text = text
	p.PaddingLeft = (w - len(text)) / 2
	p.SetRect(0, 0, w, 3)

	ui.Render(p)
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ptitle := widgets.NewParagraph()
	updateTitle(ptitle, "Hello World!")

	plog := widgets.NewParagraph()
	updateLog(plog, "log...")

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "q" {
				break
			} else {
				updateLog(plog, string(e.ID))
				//plog.Text = string(e.ID)
				//ui.Render(plog)
				//fmt.Println("key :", e.ID)
			}
		}
	}
}
