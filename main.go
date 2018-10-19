package main

import (
	"github.com/lxn/win"
	"syscall"
	"time"
	"unsafe"
)

const (
	KEYEVENTF_KEYDOWN     = 0 //key UP
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002 //key UP
	KEYEVENTF_UNICODE     = 0x0004
	KEYEVENTF_SCANCODE    = 0x0008 // scancode
)

type HWND uintptr

func auto() {
	className, _ := syscall.UTF16PtrFromString("Notepad")
	wndName, _ := syscall.UTF16PtrFromString("無題 - メモ帳")
	hwnd := win.FindWindow(className, wndName)

	win.SetForegroundWindow(hwnd)

	time.Sleep(time.Second * 2)

	sendkeys("お熱い時間が、マジできた！！！")
	sendkeyVk(win.VK_RETURN)
	sendkeyVk(win.VK_RETURN)
	sendkeys("by ルパン三世")
}

func main() {
	auto()
}

func sendkeyVk(vk uint16) {
	var inputs []win.KEYBD_INPUT
	inputs = append(inputs, win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki: win.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYDOWN,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki: win.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYUP,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	cbSize := int32(unsafe.Sizeof(win.KEYBD_INPUT{}))

	for _, inp := range inputs {
		win.SendInput(1, unsafe.Pointer(&inp), cbSize)
	}
}

func sendkey(s uint16) {
	var inputs []win.KEYBD_INPUT
	inputs = append(inputs, win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki: win.KEYBDINPUT{
			WVk:         0,
			WScan:       s,
			DwFlags:     KEYEVENTF_KEYDOWN | KEYEVENTF_UNICODE,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki: win.KEYBDINPUT{
			WVk:         0,
			WScan:       s,
			DwFlags:     KEYEVENTF_KEYUP | KEYEVENTF_UNICODE,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	cbSize := int32(unsafe.Sizeof(win.KEYBD_INPUT{}))

	for _, inp := range inputs {
		win.SendInput(1, unsafe.Pointer(&inp), cbSize)
	}
}

func sendkeys(str string) {
	for _, s := range str {
		sendkey(uint16(s))
		time.Sleep(time.Millisecond * 300)
	}
}
