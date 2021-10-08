// cpp_dll.cpp : 定义 DLL 应用程序的导出函数。
//

#include "stdafx.h"
#include <stdio.h>

extern "C" __declspec(dllexport) int add(int a, int b)
{
	return (a + b);
}

extern "C" __declspec(dllexport) int sub(int a, int b)
{
	return (a - b);
}

// 测试golang指针传输
extern "C" __declspec(dllexport) void * point(void *ctx){
	printf("ctx:%p\n", ctx);
	return ctx;
}

