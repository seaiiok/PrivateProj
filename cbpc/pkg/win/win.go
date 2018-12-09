package win

import (
	"syscall"
	"unsafe"
)

// 调用系统API MessageBox
func messageBoxApi(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}



// MessageBox
func MessageBox(msg string,flags uint) int {
	const (
		NULL  = 0
		caption="iFixTools 附加数采"
	)
	return messageBoxApi(NULL, msg, caption,flags)
}

