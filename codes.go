package keycodes

type Modifier uint8

const (
	Key_Control Modifier = 1 << iota
	Key_Shift
	Key_Alt
	Key_AltGr
	Key_Meta

	Key_Option  = Key_Alt
	Key_Command = Key_Meta
	Key_Windows = Key_Meta
)

type KeyCode uint32

type KeyInfo struct {
	Code KeyCode
	Name string
}

const UnknownKey KeyCode = 0xFFFF

func (k KeyCode) Info() KeyInfo {
	if info, ok := infoMap[k]; ok {
		return info
	}
	return KeyInfo{UnknownKey, "Unknown key"}
}

var infoMap = map[KeyCode]KeyInfo{
	Key_Home:         KeyInfo{Key_Home, "Home"},
	Key_Enter:        KeyInfo{Key_Enter, "Enter"},
	Key_End:          KeyInfo{Key_End, "End"},
	Key_Backspace:    KeyInfo{Key_Backspace, "Backspace"},
	Key_Tab:          KeyInfo{Key_Tab, "Tab"},
	Key_PageUp:       KeyInfo{Key_PageUp, "PageUp"},
	Key_PageDown:     KeyInfo{Key_PageDown, "PageDown"},
	Key_Return:       KeyInfo{Key_Return, "Return"},
	Key_Escape:       KeyInfo{Key_Escape, "Escape"},
	Key_Left:         KeyInfo{Key_Left, "Left"},
	Key_Right:        KeyInfo{Key_Right, "Right"},
	Key_Up:           KeyInfo{Key_Up, "Up"},
	Key_Down:         KeyInfo{Key_Down, "Down"},
	Key_Help:         KeyInfo{Key_Help, "Help"},
	Key_Delete:       KeyInfo{Key_Delete, "Delete"},
	Key_F1:           KeyInfo{Key_F1, "F1"},
	Key_F2:           KeyInfo{Key_F2, "F2"},
	Key_F3:           KeyInfo{Key_F3, "F3"},
	Key_F4:           KeyInfo{Key_F4, "F4"},
	Key_F5:           KeyInfo{Key_F5, "F5"},
	Key_F6:           KeyInfo{Key_F6, "F6"},
	Key_F7:           KeyInfo{Key_F7, "F7"},
	Key_F8:           KeyInfo{Key_F8, "F8"},
	Key_F9:           KeyInfo{Key_F9, "F9"},
	Key_F10:          KeyInfo{Key_F10, "F10"},
	Key_F11:          KeyInfo{Key_F11, "F11"},
	Key_F12:          KeyInfo{Key_F12, "F12"},
	Key_F13:          KeyInfo{Key_F13, "F13"},
	Key_F14:          KeyInfo{Key_F14, "F14"},
	Key_F15:          KeyInfo{Key_F15, "F15"},
	Key_F16:          KeyInfo{Key_F16, "F16"},
	Key_Insert:       KeyInfo{Key_Insert, "Insert"},
	Key_Begin:        KeyInfo{Key_Begin, "Begin"},
	Key_PrintScreen:  KeyInfo{Key_PrintScreen, "PrintScreen"},
	Key_ScrollLock:   KeyInfo{Key_ScrollLock, "ScrollLock"},
	Key_Pause:        KeyInfo{Key_Pause, "Pause"},
	Key_SysReq:       KeyInfo{Key_SysReq, "SysReq"},
	Key_Break:        KeyInfo{Key_Break, "Break"},
	Key_Reset:        KeyInfo{Key_Reset, "Reset"},
	Key_Stop:         KeyInfo{Key_Stop, "Stop"},
	Key_Menu:         KeyInfo{Key_Menu, "Menu"},
	Key_User:         KeyInfo{Key_User, "User"},
	Key_System:       KeyInfo{Key_System, "System"},
	Key_Print:        KeyInfo{Key_Print, "Print"},
	Key_ClearLine:    KeyInfo{Key_ClearLine, "ClearLine"},
	Key_ClearDisplay: KeyInfo{Key_ClearDisplay, "ClearDisplay"},
	Key_InsertLine:   KeyInfo{Key_InsertLine, "InsertLine"},
	Key_DeleteLine:   KeyInfo{Key_DeleteLine, "DeleteLine"},
	Key_InsertChar:   KeyInfo{Key_InsertChar, "InsertChar"},
	Key_DeleteChar:   KeyInfo{Key_DeleteChar, "DeleteChar"},
	Key_Prev:         KeyInfo{Key_Prev, "Prev"},
	Key_Next:         KeyInfo{Key_Next, "Next"},
	Key_Select:       KeyInfo{Key_Select, "Select"},
	Key_Execute:      KeyInfo{Key_Execute, "Execute"},
	Key_Undo:         KeyInfo{Key_Undo, "Undo"},
	Key_Redo:         KeyInfo{Key_Redo, "Redo"},
	Key_Find:         KeyInfo{Key_Find, "Find"},
	Key_ModeSwitch:   KeyInfo{Key_ModeSwitch, "ModeSwitch"},
}
