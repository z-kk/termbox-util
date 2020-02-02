package termboxUtil

import (
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

func (win *Window) Init() {
	win.cursor.SetPoint(0, 0)
	win.strings = []string{""}
	win.PosInit()
}

func (win *Window) GetPos() (int, int) {
	return win.posx, win.posy
}

func (win *Window) SetPos(x, y int) {
	win.posx = x
	win.posy = y
}

func (win *Window) PosInit() {
	win.SizeInit()
	w, h := termbox.Size()
	x := (w - win.width) / 2
	y := (h - win.height) / 2
	win.SetPos(x, y)
}

func (win *Window) GetSize() (int, int) {
	return win.width, win.height
}

func (win *Window) ResetSize() {
	win.PosInit()
	win.ResetString()
}

func (win *Window) SetSize(w, h int) {
	win.width = w
	win.height = h
	win.ResetString()
}

func (win *Window) SizeInit() {
	w, h := termbox.Size()
	w = w * 4 / 5
	h = h * 2 / 3
	win.SetSize(w, h)
}

func (win *Window) MoveCursor(x, y int) {
	if x < 0 {
		x = 0
	} else if win.Row() > len(win.showStrings) - 1 {
		x = 0
	} else if x > runewidth.StringWidth(win.showStrings[win.Row()]) {
		x = runewidth.StringWidth(win.showStrings[win.Row()])
	}
	if y < 0 {
		win.Scroll(y)
		y = 0
	} else if y > win.height - 1 {
		win.Scroll(y - win.height + 1)
		y = win.height - 1
	}
	if y > win.headLine + len(win.showStrings) - 1 {
		y = win.headLine + len(win.showStrings) - 1
	}

	win.cursor.SetPoint(x, y)
	MoveCursor(win.posx + x, win.posy + y)
}

func (win *Window) MoveCursorDiff(x, y int) {
	if y == 0 {
		i := 0;
		if x < 0 {
			for i = -1; x < 0; i-- {
				if win.Row() + i < 0 {
					break
				}
				x += runewidth.StringWidth(win.showStrings[win.Row() + i])
			}
		} else if win.Row() > len(win.showStrings) - 1 {
			i = -1
		} else if x > runewidth.StringWidth(win.showStrings[win.Row()]) {
			for i = 1; x > runewidth.StringWidth(win.showStrings[win.Row() + i]); i++ {
				x -= runewidth.StringWidth(win.showStrings[win.Row() + i])
			}
		}
		y += i;
	}

	win.MoveCursor(win.cursor.x + x, win.cursor.y + y)
}

func (win *Window) Row() (int) {
	return win.headLine + win.cursor.y
}

func (win *Window) SetString(s, h, f []string) {
	win.strings = s
	win.header = h
	win.footer = f
	win.ResetString()
	win.ShowStringLine(0)
}

func (win *Window) ResetString() {
	win.showStrings = []string{}
	win.showStringsIndex = []int{}
	var i int
	for j, s := range win.strings {
		runes := []rune(s)
		win.showStrings = append(win.showStrings, "")
		win.showStringsIndex = append(win.showStringsIndex, j)
		for _, r := range runes {
			if runewidth.RuneWidth(r) + runewidth.StringWidth(win.showStrings[i]) > win.width {
				i++
				win.showStrings = append(win.showStrings, "  ")
				win.showStringsIndex = append(win.showStringsIndex, j)
			}
			win.showStrings[i] += string(r)
		}
		i++
	}
	if len(win.showStrings) > 0 {
		win.ShowString()
		win.MoveCursorDiff(0, 0)
	} else {
		win.MoveCursor(0, 0)
	}
}

func (win *Window) ShowStringLine(lineNo int) {
	win.headLine = len(win.showStrings)
	for i, j := range win.showStringsIndex {
		if j == lineNo {
			win.headLine = i
			break
		}
	}
	win.ShowString()
}

func (win *Window) Scroll(scrollCount int) {
	win.headLine += scrollCount
	win.ShowString()
	win.MoveCursorDiff(0, 0)
}

func (win *Window) ShowString() {
	termbox.Clear(colDef, colDef)
	_, h := termbox.Size()
	// ヘッダ
	if win.posy <= len(win.header) && h > len(win.header) + 2 {
		btm := win.posy + win.height
		win.posy = len(win.header) + 1
		if win.posy < btm && win.posy + win.height > h {
			win.height = h - win.posy
		}
	}
	for i, s := range win.header {
		MoveWriteStr(win.posx, i, s)
	}
	MoveWriteStr(win.posx, win.posy - 1, "----")
	//フッタ
	if win.posy + win.height + len(win.footer) > h - 1 {
		if win.posy + len(win.footer) < h - 1 {
			win.height = h - win.posy - len(win.footer) - 1
		}
	}
	for i, s := range win.footer {
		MoveWriteStr(win.posx, i + win.posy + win.height + 1, s)
	}
	MoveWriteStr(win.posx, win.posy + win.height, "----")

	if win.headLine + win.height > len(win.showStrings) {
		win.headLine = len(win.showStrings) - win.height
	}
	if win.headLine < 0 {
		win.headLine = 0
	}
	for i, s := range win.showStrings {
		if i < win.headLine {
			continue
		} else if i >= win.headLine + win.height {
			break
		}
		MoveWritelnStr(win.posx, win.posy + i - win.headLine, s)
	}
}
