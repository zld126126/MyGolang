package main

import (
	"fmt"
	"syscall"
)

const dllFileName = "cpp_dll.dll"

// 调用C++函数参数需要转换
func IntPtr(n int) uintptr {
	return uintptr(n)
}

func GoCallDll1(a, b int) {
	dllFile := syscall.NewLazyDLL(dllFileName)
	fmt.Println("dll:", dllFile.Name)
	add := dllFile.NewProc("add")
	fmt.Println("+++++++NewProc:", add, "+++++++")

	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil {
		fmt.Println(dllFileName+fmt.Sprintf(":%d+%d", a, b)+"运算结果为:", ret)
	}
}

func GoCallDll2(a, b int) {
	dllFile, _ := syscall.LoadLibrary(dllFileName)
	fmt.Println("+++++++syscall.LoadLibrary:", dllFile, "+++++++")
	defer syscall.FreeLibrary(dllFile)
	add, err := syscall.GetProcAddress(dllFile, "add")
	fmt.Println("GetProcAddress", add)

	ret, _, err := syscall.Syscall(add,
		2,
		IntPtr(a),
		IntPtr(b),
		0)
	if err != nil {
		fmt.Println(dllFileName+fmt.Sprintf(":%d+%d", a, b)+"运算结果为:", ret)
	}
}

func GoCallDll3(a, b int) {
	DllTestDef := syscall.MustLoadDLL(dllFileName)
	add := DllTestDef.MustFindProc("add")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil {
		fmt.Println(dllFileName+fmt.Sprintf(":%d+%d", a, b)+"的运算结果为:", ret)
	}
}

// 测试golang 调用 dll方法
func GoCallDllTest(){
	// 三种调用方式
	GoCallDll1(4, 5)
	GoCallDll2(3, 6)
	GoCallDll3(2, 7)
}

// 等待命令
func Pause() {
	var str string
	fmt.Println("")
	fmt.Print("请按任意键继续...")
	fmt.Scanln(&str)
	fmt.Print("程序退出...")
}

func main() {
	GoCallDllTest()
	Pause()
}
