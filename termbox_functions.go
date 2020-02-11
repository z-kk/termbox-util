package termboxUtil

import (
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

const (
	colDef = termbox.ColorDefault
)

var (
	cursor Cursor
)

func Init() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	SetCursor(0, 0)
	return nil
}

func Close() {
	termbox.Close()
}

func GetCursorPoint() (int, int) {
	return cursor.GetPoint()
}

func HideCursor() {
	termbox.HideCursor()
	cursor.isShow = false
}

func MoveCursor(x, y int) {
	cursor.SetPoint(x, y)
	if cursor.isShow {
		ShowCursor()
	}
}

func SetCursor(x, y int) {
	cursor.isShow = true
	MoveCursor(x, y)
}

func ShowCursor() {
	cursor.isShow = true
	termbox.SetCursor(cursor.GetPoint())
}

func WritelnStr(s string) {
	MoveWritelnStr(cursor.x, cursor.y, s)
}

func MoveWritelnStr(x, y int, s string) {
	MoveWritelnStrCol(x, y, s, colDef, colDef)
}

func MoveWritelnStrCol(x, y int, s string, fgCol, bgCol termbox.Attribute) {
	WriteString(x, y, s, fgCol, bgCol, false)
}

func WriteStr(s string) {
	MoveWriteStr(cursor.x, cursor.y, s)
}

func MoveWriteStr(x, y int, s string) {
	MoveWriteStrCol(x, y, s, colDef, colDef)
}

func WriteStrCol(s string, fgCol, bgCol termbox.Attribute) {
	MoveWriteStrCol(cursor.x, cursor.y, s, fgCol, bgCol)
}

func MoveWriteStrCol(x, y int, s string, fgCol, bgCol termbox.Attribute) {
	WriteString(x, y, s, fgCol, bgCol, true)
}

func WriteString(x, y int, s string, fgCol, bgCol termbox.Attribute, isWrap bool) {
	runes := []rune(s)

	for _, r := range runes {
		if isWrap && runewidth.RuneWidth(r) > 1 {
			w, _ := termbox.Size()
			if x == w - 1 {
				x = 0
				y++
			}
		}
		termbox.SetCell(x, y, r, fgCol, bgCol)
		if r == 9 {
			x += 4 - x % 4  // TAB
		} else {
			x += runewidth.RuneWidth(r)
		}
		if isWrap {
			w, _ := termbox.Size()
			if x >= w {
				x = 0
				y++
			}
		}
	}
	MoveCursor(x, y)
}
