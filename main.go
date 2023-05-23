package main

import "github.com/aarzilli/golua/lua"

func main() {
	L := LuaPool.Get().(*lua.State)
	L.SetTop(0)

	LuaPool.Put(L)

	L = LuaPool.Get().(*lua.State)
	L.SetTop(0)

	// 在go中调用lua函数
	L.GetField(lua.LUA_GLOBALSINDEX, "add")
	L.PushInteger(1)
	L.PushInteger(2)
	L.Call(2, 2)
	result1 := L.ToInteger(1)
	result2 := L.ToInteger(2)
	_ = result1
	_ = result2
	println(result1, result2)
	L.Pop(2)

	L.GetField(lua.LUA_GLOBALSINDEX, "GetStr")
	L.Call(0, 1)
	resultStr := L.ToString(1)
	println(resultStr)

	// 在lua中回调go接口
	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoLua")
	L.Call(0, 0)

	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoWithParamLua")
	L.Call(0, 0)

	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoWithParamResultLua")
	L.Call(0, 0)
}
