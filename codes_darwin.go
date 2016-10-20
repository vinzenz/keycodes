package keycodes

// #cgo LDFLAGS: -framework Cocoa -framework Carbon
// #include "darwin_impl.c"
import "C"
import "fmt"
import "unicode"

type TranslateInfo struct {
	Code   KeyCode
	String string
}

func TranslateKey(code uint32) TranslateInfo {
	buffer := C.translate(C.uint32(code))
	if buffer.Length > 0 {
		value := buffer.buffer[0]
		r := rune(value)
		if(value > 32) {
			if unicode.IsDigit(r) || unicode.IsLetter(r) || unicode.IsPunct(r) || unicode.IsSpace(r) {
				return TranslateInfo{
					KeyCode(value),
					string([]rune{rune(value)}),
				}
			}
		}
		info := KeyCode(value).Info()
		if info.Code != UnknownKey {
			return TranslateInfo{
				info.Code,
				info.Name,
			}
		}
	}
	return TranslateInfo{
		0,
		"",
	}
}

const (
	Key_Home      = C.kHomeCharCode
	Key_Enter     = C.kEnterCharCode
	Key_End       = C.kEndCharCode
	Key_Backspace = C.kBackspaceCharCode
	Key_Tab       = C.kTabCharCode
	Key_PageUp    = C.kPageUpCharCode
	Key_PageDown  = C.kPageDownCharCode
	Key_Return    = C.kReturnCharCode
	Key_Escape    = C.kEscapeCharCode
	Key_Left      = C.kLeftArrowCharCode
	Key_Right     = C.kRightArrowCharCode
	Key_Up        = C.kUpArrowCharCode
	Key_Down      = C.kDownArrowCharCode
	Key_Help      = C.kHelpCharCode
	Key_Delete    = C.kDeleteCharCode

	Key_F1  = 122
	Key_F2  = 120
	Key_F3  = 99
	Key_F4  = 118
	Key_F5  = 96
	Key_F6  = 97
	Key_F7  = 98
	Key_F8  = 100
	Key_F9  = 101
	Key_F10 = 109
	Key_F11 = 103
	Key_F12 = 111
	Key_F13 = 105
	Key_F14 = 107
	Key_F15 = 113
	Key_F16 = 106

	//	Key_Up = 0xF700
	//	Key_Down = 0xF701
	//	Key_Left = 0xF702
	//	Key_Right = 0xF703
	Key_Insert = 0xF727
	//	Key_Delete = 0xF728
	//	Key_Home = 0xF729
	Key_Begin = 0xF72A
	//	Key_End = 0xF72B
	//	Key_PageUp = 0xF72C
	//	Key_PageDown = 0xF72D
	Key_PrintScreen  = 0xF72E
	Key_ScrollLock   = 0xF72F
	Key_Pause        = 0xF730
	Key_SysReq       = 0xF731
	Key_Break        = 0xF732
	Key_Reset        = 0xF733
	Key_Stop         = 0xF734
	Key_Menu         = 0xF735
	Key_User         = 0xF736
	Key_System       = 0xF737
	Key_Print        = 0xF738
	Key_ClearLine    = 0xF739
	Key_ClearDisplay = 0xF73A
	Key_InsertLine   = 0xF73B
	Key_DeleteLine   = 0xF73C
	Key_InsertChar   = 0xF73D
	Key_DeleteChar   = 0xF73E
	Key_Prev         = 0xF73F
	Key_Next         = 0xF740
	Key_Select       = 0xF741
	Key_Execute      = 0xF742
	Key_Undo         = 0xF743
	Key_Redo         = 0xF744
	Key_Find         = 0xF745
	//	Key_Help = 0xF746
	Key_ModeSwitch = 0xF747
)
