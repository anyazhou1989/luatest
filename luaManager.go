/**
 * Title: luaManager
 * Author: Anyazhou
 * Date: 5/23/23 11:06 AM
 * Description: This is a luaManager file.
 */
package main

import (
	"github.com/aarzilli/golua/lua"
	"sync"
)

func CallBackGo(L *lua.State) int {
	println("anyazhou go")
	return 0
}

func CallBackGoWithParam(L *lua.State) int {
	arg1 := L.ToString(1)
	arg2 := L.ToString(2)
	println(arg1 + arg2)
	return 0
}

func CallBackGoWithParamResult(L *lua.State) int {
	arg1 := L.ToString(1)
	arg2 := L.ToString(2)
	result := arg1 + arg2
	L.PushString(result)
	L.PushString("success")

	return 2
}

// LuaPool 在这里可以设定几个不同的luapool，比如AI，副本等
var LuaPool = sync.Pool{
	New: func() interface{} {
		L := lua.NewState()
		L.OpenLibs()
		L.DoFile("./main.lua")

		L.Register("CallBackGo", CallBackGo)
		L.Register("CallBackGoWithParam", CallBackGoWithParam)
		L.Register("CallBackGoWithParamResult", CallBackGoWithParamResult)

		return L
	},
}
