package main

import "github.com/aarzilli/golua/lua"

//
////type I1 interface {
////	I1M1()
////	I1M2()
////	I1M3()
////}
////
////type I1Impl struct {
////}
////
////func (name *I1Impl) I1M1() {
////	println("I1M1")
////}
////
////func (name *I1Impl) I1M2() {
////	println("I1M2")
////}
////
////func (name *I1Impl) I1M3() {
////	println("I1M3")
////}
//
////type I2 interface {
////	I2M1()
////	I2M2()
////	I2M3()
////}
////
////type O1 struct{}
////
////func (o *O1) O1M1() {}
////func (o *O1) O1M2() {}
////func (o *O1) O1M3() {}
////
////type O2 struct{}
////
////func (o *O2) O2M1() {}
////func (o *O2) O2M2() {}
////func (o *O2) O2M3() {}
////
////type O3 struct{}
////
////func (o *O3) O3M1() {}
////func (o *O3) O3M2() {}
////func (o *O3) O3M3() {}
////
////type O4 struct{}
////
////func (o *O4) O4M1() {}
////func (o *O4) O4M2() {}
////func (o *O4) O4M3() {}
////
////type O99 struct {
////	//I1
////	I1 I1
////	//I2 I2
////	//
////	//O1
////	//O2 O2
////	//*O3
////	//O4 *O4
////}
////
////func (o *O99) O99M1() {}
////func (o *O99) O99M2() {}
////func (o *O99) O99M3() {}
////
////func deferName() {
////	var name string
////	name = "anyazhou"
////	println(name)
////}
////
////var stringPool = sync.Pool{
////	New: func() interface{} {
////		return ""
////	},
////}
////
////func getGoroutineID() string {
////	var buf [64]byte
////	n := runtime.Stack(buf[:], false)
////	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
////	return idField
////}
//
//// go build -gcflags=-m main.go
//func generate8191() {
//	nums := make([]int, 8191) // < 64KB
//	for i := 0; i < 8191; i++ {
//		nums[i] = rand.Int()
//	}
//}
//
//func generate8193() {
//	nums := make([]int, 8193) // = 64KB
//	for i := 0; i < 8193; i++ {
//		nums[i] = rand.Int()
//	}
//}
//
//func generate(n int) {
//	nums := make([]int, n) // 不确定大小
//	for i := 0; i < n; i++ {
//		nums[i] = rand.Int()
//	}
//}

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

	//generate8191()
	//generate8193()
	//generate(1)
	//// 声明源数组
	//sourceArray := [5]int{1, 2, 3, 4, 5}
	//// 声明目标数组
	//targetArray := [5]int{}
	//copy(targetArray[:], sourceArray[:])
	//
	//fmt.Println("开始执行程序")
	//println(getGoroutineID())
	//
	//time.AfterFunc(time.Second*5, func() {
	//	fmt.Println("5秒后执行该函数")
	//	println(getGoroutineID())
	//})
	//time.Sleep(time.Second * 10)
	//fmt.Println("程序结束")
	//println(getGoroutineID())
	//var flag atomic.Bool
	//var num atomic.Int32
	//num.Inc()
	//flag.Toggle()
	//var name string
	//name = "anyazhou"
	//namePtr := &name
	//println(*namePtr)
	//
	//name = "anyazhou123"
	//println(*namePtr)
	/*
		var name string
		fmt.Printf("str1: %s, address: %p\n", name, &name)

		stringPool.Put(name)

		str1 := stringPool.Get().(string)
		fmt.Printf("str1: %s, address: %p\n", str1, &str1)
		stringPool.Put(str1)

		str1 = stringPool.Get().(string)
		fmt.Printf("str1: %s, address: %p\n", str1, &str1)


	*/
	/*
		{
			f, err := os.Create("test.txt")
			_ = err
			f.WriteString("222222")
			defer f.Close()
			defer println("3")
			defer println("4")
			println("1")
		}
		println("2")
		defer deferName()


	*/
	/*
		fSet := token.NewFileSet()

		curAst, err := parser.ParseFile(fSet, "test.go", nil, 0)
		if err != nil {
			println("success")
		}
		for _, decl := range curAst.Decls {
			//对函数声明进行分析，获取函数名
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				line := fSet.Position(funcDecl.Pos()).Line
				_ = line
			}
		}
		_ = curAst
		o := &O99{
			I1: &I1Impl{},
		}
		o.I1.I1M1()
	*/
	//t := reflect.TypeOf(o)
	//fmt.Println("O99's NumMethod()=", t.NumMethod())
	//for i := 0; i < t.NumMethod(); i++ {
	//	fmt.Printf("O99's Method(%d).Name=%s\n", i, t.Method(i).Name)
	//}
}
