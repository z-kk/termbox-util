package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

func ShowItems(s []string) {
	ShowItemsHeadFoot(s, []string{}, []string{})
}

func ShowItemsHead(s, h []string) {
	ShowItemsHeadFoot(s, h, []string{})
}

func ShowItemsFoot(s, f []string) {
	ShowItemsHeadFoot(s, []string{}, f)
}

func ShowItemsHeadFoot(s, h, f []string) {
	Init()
	defer Close()
	HideCursor()
	var win Window
	win.Init()
	win.SetString(s, h, f)
	maxLine := len(win.strings)

mainloop:
	for {
		termbox.Flush()
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter, termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowUp:
				win.Scroll(-1)
			case termbox.KeyArrowDown:
				win.Scroll(1)
			case termbox.KeyArrowLeft, termbox.KeyPgup:
				win.Scroll(-win.height + 1)
			case termbox.KeyArrowRight, termbox.KeyPgdn:
				win.Scroll(win.height - 1)
			case termbox.KeyHome:
				win.ShowStringLine(0)
			case termbox.KeyEnd:
				win.ShowStringLine(maxLine)
			default:
				switch ev.Ch {
				case 'h', 'b':
					win.Scroll(-win.height + 1)
				case 'j':
					win.Scroll(1)
				case 'k':
					win.Scroll(-1)
				case 'l', 'f':
					win.Scroll(win.height - 1)
				case 'g':
					win.ShowStringLine(0)
				case 'G':
					win.ShowStringLine(maxLine)
				case 'q':
					break mainloop
				}
			}
		case termbox.EventResize:
			for i := 0; i < 3; i++ {
				win.ResetSize()
			}
		}
	}
}

func SelectItem(s []string) (int) {
	return SelectItemHeadFoot(s, []string{}, []string{})
}

func SelectItemHead(s, h []string) (int) {
	return SelectItemHeadFoot(s, h, []string{})
}

func SelectItemFoot(s, f []string) (int) {
	return SelectItemHeadFoot(s, []string{}, f)
}

func SelectItemHeadFoot(s, h, f []string) (int) {
	for i := 0; i < len(s); i++  {
		s[i] = "  " + s[i];
	}

	Init()
	defer Close()
	HideCursor()
	var win Window
	win.Init()
	win.SetString(s, h, f)
	win.MoveCursor(0, 0)
	maxLine := len(win.strings)
	selVal := 0

mainloop:
	for {
		showLen := len(win.showStrings)
		for i := 0; i < maxLine; i++ {
			if i == selVal {
				win.strings[i] = "> " + win.strings[i][2:]
			} else {
				win.strings[i] = "  " + win.strings[i][2:]
			}
		}
		win.ResetString()
		termbox.Flush()
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter, termbox.KeySpace:
				break mainloop
			case termbox.KeyEsc:
				return -1
			case termbox.KeyArrowUp:
				for win.Row() > 0 && win.showStringsIndex[win.Row()] == selVal {
					win.MoveCursorDiff(0, -1)
				}
			case termbox.KeyArrowDown:
				for win.Row() < showLen - 1 && win.showStringsIndex[win.Row()] == selVal {
					win.MoveCursorDiff(0, 1)
				}
			case termbox.KeyArrowLeft, termbox.KeyPgup:
				win.Scroll(-win.height + 1)
				if win.showStringsIndex[win.Row()] == selVal {
					win.MoveCursor(0, 0)
				}
			case termbox.KeyArrowRight, termbox.KeyPgdn:
				win.Scroll(win.height - 1)
				if win.showStringsIndex[win.Row()] == selVal {
					win.MoveCursor(0, maxLine)
				}
			case termbox.KeyHome:
				win.ShowStringLine(0)
				win.MoveCursor(0, 0)
			case termbox.KeyEnd:
				win.ShowStringLine(maxLine)
				win.MoveCursor(0, maxLine)
			default:
				switch ev.Ch {
				case 'h', 'b':
					win.Scroll(-win.height + 1)
					if win.showStringsIndex[win.Row()] == selVal {
						win.MoveCursor(0, 0)
					}
				case 'j':
					for win.Row() < showLen - 1 && win.showStringsIndex[win.Row()] == selVal {
						win.MoveCursorDiff(0, 1)
					}
				case 'k':
					for win.Row() > 0 && win.showStringsIndex[win.Row()] == selVal {
					win.MoveCursorDiff(0, -1)
				}
				case 'l', 'f':
					win.Scroll(win.height - 1)
					if win.showStringsIndex[win.Row()] == selVal {
						win.MoveCursor(0, maxLine)
					}
				case 'g':
					win.ShowStringLine(0)
					win.MoveCursor(0, 0)
				case 'G':
					win.ShowStringLine(maxLine)
					win.MoveCursor(0, maxLine)
				case 'H':
					win.MoveCursor(0, 0)
				case 'L':
					win.MoveCursor(0, win.height - 1)
				case 'q':
					return -1
				}
			}
		case termbox.EventResize:
			for i := 0; i < 3; i++ {
				win.ResetSize()
			}
		}
		// 現在の選択値を保存
		selVal = win.showStringsIndex[win.Row()]

		// カーソル行を選択文字列の最後の行に移動する
		for win.Row() < showLen - 1 {
			if win.showStringsIndex[win.Row() + 1] > selVal {
				break
			}
			win.MoveCursorDiff(0, 1)
		}
		// カーソル行を選択文字列の最初の行に移動する
		for win.Row() > 0 {
			if win.showStringsIndex[win.Row() - 1] < selVal {
				break
			}
			win.MoveCursorDiff(0, -1)
		}
	}
	return selVal
}

func SelectItems(s []string) ([]int) {
	return SelectItemsHeadFoot(s, []string{}, []string{})
}

func SelectItemsHead(s, h []string) ([]int) {
	return SelectItemsHeadFoot(s, h, []string{})
}

func SelectItemsFoot(s, f []string) ([]int) {
	return SelectItemsHeadFoot(s, []string{}, f)
}

func SelectItemsHeadFoot(s, h, f []string) ([]int) {
	res := []int{}
	selected := []bool{}
	for i := 0; i < len(s); i++  {
		s[i] = "[ ] " + s[i];
		selected = append(selected, false)
	}
	exit := false
	tabExit := false
	if len(f) == 0 {
		tabExit = true
	} else if len(f) > 2 {
		tabExit = len(f[0]) == 0 && len(f[1]) == 0
	}

	Init()
	defer Close()
	ShowCursor()
	var win Window
	win.Init()
	win.SetString(s, h, f)
	win.MoveCursor(1, 0)
	maxLine := len(win.strings)
	selRow := 0

mainloop:
	for {
		showLen := len(win.showStrings)
		for i := 0; i < maxLine; i++ {
			if selected[i] {
				win.strings[i] = "[X] " + win.strings[i][4:]
			} else {
				win.strings[i] = "[ ] " + win.strings[i][4:]
			}
		}
		win.ResetString()
		if tabExit {
			if exit {
				MoveWritelnStrCol(win.posx + win.width / 2 - 1, win.posy + win.height + 2, "OK", termbox.ColorBlack, termbox.ColorWhite)
				HideCursor()
			} else {
				MoveWritelnStr(win.posx + win.width / 2 - 1, win.posy + win.height + 2, "OK")
				ShowCursor()
			}
		}
		win.MoveCursorDiff(0, 0)
		termbox.Flush()
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter, termbox.KeySpace:
				if exit {
					break mainloop
				} else {
					selected[win.showStringsIndex[win.Row()]] = !selected[win.showStringsIndex[win.Row()]]
				}
			case termbox.KeyEsc, termbox.KeyBackspace2:
				break mainloop
			case termbox.KeyTab:
				if tabExit {
					exit = !exit
				}
			case termbox.KeyArrowUp:
				for win.Row() > 0 && win.showStringsIndex[win.Row()] == selRow {
					win.MoveCursorDiff(0, -1)
				}
			case termbox.KeyArrowDown:
				for win.Row() < showLen - 1 && win.showStringsIndex[win.Row()] == selRow {
					win.MoveCursorDiff(0, 1)
				}
			case termbox.KeyArrowLeft, termbox.KeyPgup:
				win.Scroll(-win.height + 1)
				if win.showStringsIndex[win.Row()] == selRow {
					win.MoveCursor(1, 0)
				}
			case termbox.KeyArrowRight, termbox.KeyPgdn:
				win.Scroll(win.height - 1)
				if win.showStringsIndex[win.Row()] == selRow {
					win.MoveCursor(1, maxLine)
				}
			case termbox.KeyHome:
				win.ShowStringLine(0)
				win.MoveCursor(1, 0)
			case termbox.KeyEnd:
				win.ShowStringLine(maxLine)
				win.MoveCursor(1, maxLine)
			default:
				switch ev.Ch {
				case 'h', 'b':
					win.Scroll(-win.height + 1)
					if win.showStringsIndex[win.Row()] == selRow {
						win.MoveCursor(1, 0)
					}
				case 'j':
					for win.Row() < showLen - 1 && win.showStringsIndex[win.Row()] == selRow {
						win.MoveCursorDiff(0, 1)
					}
				case 'k':
					for win.Row() > 0 && win.showStringsIndex[win.Row()] == selRow {
						win.MoveCursorDiff(0, -1)
					}
				case 'l', 'f':
					win.Scroll(win.height - 1)
					if win.showStringsIndex[win.Row()] == selRow {
						win.MoveCursor(1, maxLine)
					}
				case 'g':
					win.ShowStringLine(0)
					win.MoveCursor(1, 0)
				case 'G':
					win.ShowStringLine(maxLine)
					win.MoveCursor(1, maxLine)
				case 'H':
					win.MoveCursor(1, 0)
				case 'L':
					win.MoveCursor(1, win.height - 1)
				}
			}
		case termbox.EventResize:
			for i := 0; i < 3; i++ {
				win.ResetSize()
			}
		}
		// 現在の選択行を保存
		selRow = win.showStringsIndex[win.Row()]

		// カーソル行を選択文字列の最後の行に移動する
		for win.Row() < showLen - 1 {
			if win.showStringsIndex[win.Row() + 1] > selRow {
				break
			}
			win.MoveCursorDiff(0, 1)
		}
		// カーソル行を選択文字列の最初の行に移動する
		for win.Row() > 0 {
			if win.showStringsIndex[win.Row() - 1] < selRow {
				break
			}
			win.MoveCursorDiff(0, -1)
		}
	}
	for i := 0; i < len(selected); i++ {
		if selected[i] {
			res = append(res, i)
		}
	}
	return res
}
