---
--- Generated by Luanalysis
--- Created by admin.
--- DateTime: 5/19/23 4:22 PM
---

function add(a, b)
    return a + b, a-b
end

function GetStr()
    return "hello world lua"
end

function CallBackGoLua()
    CallBackGo()
end

function CallBackGoWithParamLua()
    CallBackGoWithParam("Hello ", "World")
end

function CallBackGoWithParamResultLua()
    local result, num = CallBackGoWithParamResult("Hello ", "Anyazhou")
    print(result, num)
end