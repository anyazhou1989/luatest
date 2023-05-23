/**
 * Title: luaManager_test
 * Author: Anyazhou
 * Date: 5/23/23 11:40 AM
 * Description: This is a luaManager_test file.
 */
package main

import (
	"github.com/aarzilli/golua/lua"
	"testing"
	"time"
)

func TestLuaPool(t *testing.T) {
	L := LuaPool.Get().(*lua.State)
	L.SetTop(0)

	LuaPool.Put(L)

	L = LuaPool.Get().(*lua.State)
	defer LuaPool.Put(L)
	L.SetTop(0)

	// 在go中调用lua函数
	L.GetField(lua.LUA_GLOBALSINDEX, "add")
	L.PushInteger(1)
	L.PushInteger(2)
	L.Call(2, 2)
	result1 := L.ToInteger(1)
	result2 := L.ToInteger(2)

	if result1 != 3 {
		t.Errorf("Add(1, 2) = %d; want %d", result1, 3)
	}
	if result2 != -1 {
		t.Errorf("Dec(1, 2) = %d; want %d", result2, -1)
	}
	L.Pop(2)

	L.GetField(lua.LUA_GLOBALSINDEX, "GetStr")
	L.Call(0, 1)
	resultStr := L.ToString(1)
	if resultStr != "hello world lua" {
		t.Errorf("GetStr = %s; want %s", resultStr, "hello world lua")
	}

	// 在lua中回调go接口
	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoLua")
	L.Call(0, 0)

	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoWithParamLua")
	L.Call(0, 0)

	L.GetField(lua.LUA_GLOBALSINDEX, "CallBackGoWithParamResultLua")
	L.Call(0, 0)
	time.Sleep(time.Second)
}
