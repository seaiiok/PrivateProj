package print
 
import (
	"fmt"
    "syscall"
)

const(
    Ok=2*(iota+1)
    Err
    Warn
    Info
)
   
func Line(i int,s string) (int,error){
    kernel32 := syscall.NewLazyDLL("kernel32.dll")
    proc := kernel32.NewProc("SetConsoleTextAttribute")
    handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
    n,err:=fmt.Println(s)
    handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
    CloseHandle := kernel32.NewProc("CloseHandle")
    CloseHandle.Call(handle)
    return n,err
}

func Linef(i int,format string, a ...interface{}) (int,error){
    kernel32 := syscall.NewLazyDLL("kernel32.dll")
    proc := kernel32.NewProc("SetConsoleTextAttribute")
    handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
    n,err:=fmt.Printf(format,a)
    fmt.Println("")
    handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
    CloseHandle := kernel32.NewProc("CloseHandle")
    CloseHandle.Call(handle)
    return n,err
}
