package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func updateLog(p *widgets.Paragraph, text string) {
	w, h := ui.TerminalDimensions()
	p.Text = text
	p.TextStyle.Fg = ui.ColorYellow
	//p.TextStyle.Bg = ui.ColorCyan
	//p.Border = false
	p.BorderTop = false
	p.SetRect(0, h-3, w, h)

	ui.Render(p)
}

func updateTitle(p *widgets.Paragraph, text string) {
	w, _ := ui.TerminalDimensions()
	p.Title = "Title"
	p.Text = text
	p.TextStyle.Fg = ui.ColorYellow
	p.PaddingLeft = (w - len(text)) / 2
	p.BorderStyle.Fg = ui.ColorRed
	p.SetRect(0, 0, w, 3)

	ui.Render(p)
}

func updateTitleSub(p *widgets.Paragraph, text string) {
	w, h := ui.TerminalDimensions()
	p.Text = fmt.Sprintf("%-*s", w-2, text)
	//p.Text = fmt.Sprintf("%*s", w-2, text)
	p.TextStyle.Fg = ui.ColorBlack
	p.TextStyle.Bg = ui.ColorCyan
	//p.Border = false
	p.BorderBottom = false
	p.SetRect(0, h-5, w, h-2)

	ui.Render(p)
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ptitle := widgets.NewParagraph()
	updateTitle(ptitle, "Hello World!")

	ptitle_log := widgets.NewParagraph()
	updateTitleSub(ptitle_log, "Sub Title")

	plog := widgets.NewParagraph()
	updateLog(plog, "log...")

	// // test : make unique id
	// // 2147483648 (2^31)
	// // 20210414xx
	// // 113176xxxx (2021 * 04 * 14)
	// tmNow := time.Now()
	// year, month, day := tmNow.Date()
	// time.Now().Hour()
	// id := (year * int(month) * day) * 10000
	// fmt.Println(id, year, month, day)

	// //var cs sync.Mutex
	// var data uint16
	// go func() {
	// 	for {
	// 		//cs.Lock()
	// 		data++
	// 		//cs.Unlock()
	// 		time.Sleep(time.Nanosecond)
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		//cs.Lock()
	// 		if data == 0 {
	// 			time.Sleep(time.Nanosecond)
	// 			updateLog(plog, fmt.Sprintf("the value is %v.", data))
	// 			//time.Sleep(time.Nanosecond)
	// 		}
	// 		//cs.Unlock()
	// 	}
	// }()

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
