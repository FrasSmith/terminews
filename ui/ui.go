package ui

import (
	"fmt"
	gc "github.com/rthornton128/goncurses"
)

type Displayer interface {
	Display() string
}

type TMenuItem struct {
	*gc.MenuItem
	inc  int
	data Displayer
}

func (tmi *TMenuItem) Create(inc int, data Displayer) error {
	value := fmt.Sprintf("%3d. %v", inc, data.Display())
	item, err := gc.NewItem(value, "")
	if err != nil {
		return err
	}
	tmi.inc = inc
	tmi.data = data
	tmi.MenuItem = item

	return nil
}

type TMenu struct {
	*gc.Menu
	items []*TMenuItem
}

func (tm *TMenu) Create(tmis []*TMenuItem) error {
	tm.items = tmis
	gcMenuItems := make([]*gc.MenuItem, len(tmis))
	for i, tmi := range tmis {
		gcMenuItems[i] = tmi.MenuItem
	}
	menu, err := gc.NewMenu(gcMenuItems)
	if err != nil {
		return err
	}
	tm.Menu = menu

	return nil
}

type TWindow struct {
	*gc.Window
	title      string
	h, w, y, x int
}

func (tw *TWindow) Create(title string, h, w, y, x int) error {
	win, err := gc.NewWindow(h, w, y, x)
	if err != nil {
		return err
	}
	tw.Window = win
	tw.h = h
	tw.w = w
	tw.y = y
	tw.x = x

	tw.Keypad(true)
	tw.SetContour()
	tw.SetTitle(title)

	return nil
}

func (tw *TWindow) SetTitle(title string) {
	tw.title = title
	tw.MovePrint(1, (tw.w/2)-(len(title)/2), title)
}

func (tw *TWindow) SetContour() {
	tw.Box(0, 0)
	tw.MoveAddChar(2, 0, gc.ACS_LTEE)
	tw.HLine(2, 1, gc.ACS_HLINE, tw.w-2)
	tw.MoveAddChar(2, tw.w-1, gc.ACS_RTEE)
}

func (tw *TWindow) Focus(cc int16) {
	tw.ColorOn(cc)
	tw.SetContour()
	tw.ColorOff(cc)
}

func (tw *TWindow) Unfocus(cc int16) {
	tw.ColorOn(cc)
	tw.SetContour()
	tw.ColorOff(cc)
}

func (tw *TWindow) AttachMenu(tm *TMenu) {
	tm.Menu.SetWindow(tw.Window)
	tm.Menu.SubWindow(tw.Derived(tw.h-6, tw.w-4, 4, 2))
	tm.Menu.Format(tw.h-6, 1)
	tm.Menu.Mark("")
}
