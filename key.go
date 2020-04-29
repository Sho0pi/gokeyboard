package gokeyboard

import (
	"strings"
	"syscall"
)

var keyCodeToString = map[uint16]string{
	1:   "[ESC]",
	2:   "1",
	3:   "2",
	4:   "3",
	5:   "4",
	6:   "5",
	7:   "6",
	8:   "7",
	9:   "8",
	10:  "9",
	11:  "0",
	12:  "-",
	13:  "=",
	14:  "[BS]",
	15:  "[TAB]",
	16:  "q",
	17:  "w",
	18:  "e",
	19:  "r",
	20:  "t",
	21:  "y",
	22:  "u",
	23:  "i",
	24:  "o",
	25:  "p",
	26:  "[",
	27:  "]",
	28:  "[ENTER]",
	29:  "[L_CTRL]",
	30:  "a",
	31:  "s",
	32:  "d",
	33:  "f",
	34:  "g",
	35:  "h",
	36:  "j",
	37:  "k",
	38:  "l",
	39:  ";",
	40:  "'",
	41:  "`",
	42:  "[L_SHIFT]",
	43:  "\\",
	44:  "z",
	45:  "x",
	46:  "c",
	47:  "v",
	48:  "b",
	49:  "n",
	50:  "m",
	51:  ",",
	52:  ".",
	53:  "/",
	54:  "[R_SHIFT]",
	55:  "*",
	56:  "[L_ALT]",
	57:  "[SPACE]",
	58:  "[CAPS_LOCK]",
	59:  "[F1]",
	60:  "[F2]",
	61:  "[F3]",
	62:  "[F4]",
	63:  "[F5]",
	64:  "[F6]",
	65:  "[F7]",
	66:  "[F8]",
	67:  "[F9]",
	68:  "[F10]",
	69:  "[NUM_LOCK]",
	70:  "[SCROLL_LOCK]",
	71:  "[HOME]",
	72:  "[UP_8]",
	73:  "[PGUP_9]",
	74:  "-",
	75:  "[LEFT_4]",
	76:  "5",
	77:  "[RT_ARROW_6]",
	78:  "+",
	79:  "[END_1]",
	80:  "[DOWN]",
	81:  "[PGDN_3]",
	82:  "[INS]",
	83:  "[DEL]",
	84:  "",
	85:  "",
	86:  "",
	87:  "[F11]",
	88:  "[F12]",
	89:  "",
	90:  "",
	91:  "",
	92:  "",
	93:  "",
	94:  "",
	95:  "",
	96:  "[R_ENTER]",
	97:  "[R_CTRL]",
	98:  "/",
	99:  "[PRT_SCR]",
	100: "[R_ALT]",
	101: "",
	102: "[Home]",
	103: "[Up]",
	104: "[PgUp]",
	105: "[Left]",
	106: "[Right]",
	107: "[End]",
	108: "[Down]",
	109: "[PgDn]",
	110: "[Insert]",
	111: "[Del]",
	112: "",
	113: "",
	114: "",
	115: "",
	116: "",
	117: "",
	118: "",
	119: "[Pause]",
}

// Will represent a key, will be able to get the code and the string value from it.
// By default 2 keys are written, one when the key pressed, and one when released, therefore make sure to check the status
// with IsPressed or IsReleased
type Key struct {
	_      syscall.Timeval
	Type   uint16
	Code   uint16
	Status int32
}

// Will return the key string value, if does not exists will return empty string.
func (key *Key) ToString() string {
	return keyCodeToString[key.Code]
}

func (key *Key) IsPressed() bool {
	return key.Status == 1
}

func (key *Key) IsReleased() bool {
	return key.Status == 0
}

// Private method to check if its a key event.
func (key *Key) isKeyEvent() bool {
	return key.Type == 0x01
}

// Will check if the key is shift, for easy case handling
func (key *Key) IsShift() bool {
	return strings.Contains(keyCodeToString[key.Code], "SHIFT")
}
