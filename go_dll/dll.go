package main


import (
	"fmt"
	"syscall"
	"unsafe"
)

//export Add
func Add(a, b int) int {
	return a + b
}

var (
	goDLL  = "go.dll"
	cppDLL = "cpp_dll.dll"
)

// 调用C++函数参数需要转换
func IntPtr(n int) uintptr {
	return uintptr(n)
}

func IsFinishError(err error) bool {
	if err.Error() == "The operation completed successfully." {
		return true
	}
	return false
}

func GoCallDll1(a, b int) uintptr {
	dllFile := syscall.NewLazyDLL(cppDLL)
	fmt.Println("dll:", dllFile.Name)
	add := dllFile.NewProc("add")
	fmt.Println("+++++++NewProc:", add, "+++++++")

	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil && IsFinishError(err) {
		fmt.Println(cppDLL+fmt.Sprintf(":%d+%d", a, b)+"运算结果为:", ret)
	} else {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
	return ret
}

func GoCallDll2(a, b int) uintptr {
	dllFile, _ := syscall.LoadLibrary(cppDLL)
	fmt.Println("+++++++syscall.LoadLibrary:", dllFile, "+++++++")
	defer syscall.FreeLibrary(dllFile)
	add, err := syscall.GetProcAddress(dllFile, "add")
	fmt.Println("GetProcAddress", add)

	ret, _, err := syscall.Syscall(add,
		2,
		IntPtr(a),
		IntPtr(b),
		0)
	if err != nil && IsFinishError(err) {
		fmt.Println(cppDLL+fmt.Sprintf(":%d+%d", a, b)+"运算结果为:", ret)
	} else {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
	return ret
}

func GoCallDll3(a, b int) uintptr {
	DllTestDef := syscall.MustLoadDLL(cppDLL)
	add := DllTestDef.MustFindProc("add")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil && IsFinishError(err) {
		fmt.Println(cppDLL+fmt.Sprintf(":%d+%d", a, b)+"结果为:", ret)
	} else {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
	return ret
}

type CPoint struct {
	Name string
}

// 指针测试
func GoCallDllPoint(c *CPoint) uintptr {
	DllTestDef := syscall.MustLoadDLL(cppDLL)
	point := DllTestDef.MustFindProc("point")

	fmt.Println("+++++++MustFindProc：", point, "+++++++")
	ret, _, err := point.Call(uintptr(unsafe.Pointer(c)))
	if err != nil && IsFinishError(err) {
		p := (*CPoint)(unsafe.Pointer(ret))
		fmt.Println(cppDLL+fmt.Sprintf("原始:%p", c)+",运算结果为:", fmt.Sprintf("%p", p))
	} else {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
	return ret
}

type Common interface{}

// 测试golang 调用 dll方法
func GoCallDllTest() {
	// 三种调用方式
	res1 := GoCallDll1(4, 5)
	fmt.Println("r1:", (int)(res1))

	res2 := GoCallDll2(3, 6)
	fmt.Println("r2:", Common(res2))

	res3 := GoCallDll3(2, 7)
	fmt.Println("r3:", (int)(res3))

	// 特殊传输指针: 但是c++ 不做调用(回调使用)
	GoCallDllPoint(&CPoint{Name: "dong"})
}

// 等待命令
func Pause() {
	var str string
	fmt.Println("")
	fmt.Print("请按任意键继续...")
	fmt.Scanln(&str)
	fmt.Print("程序退出...")
}

func GoCallDll(a, b int) uintptr {
	fmt.Println(111111)
	DllTestDef := syscall.MustLoadDLL(goDLL)
	add := DllTestDef.MustFindProc("Add")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil && IsFinishError(err) {
		fmt.Println(goDLL+fmt.Sprintf(":%d+%d", a, b)+"结果为:", ret)
	} else {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
	return ret
}
