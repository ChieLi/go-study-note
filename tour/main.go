package main

import (
    "fmt"
    "math/rand"
    "math/cmplx"
    "math"
    "time"
)

func main() {
    fmt.Println("hello world")    
    // 随机数
    rand.Seed(time.Now().Unix())
    for i := 0; i < 5; i++ {
        fmt.Println(randN(100))
    }

    fmt.Println("rand add =", add(randN(100), randN(100)))
    fmt.Println("rand addV2 =", addV2(randN(100), randN(100)))
 
    a, b := swap("hello", "world")
    fmt.Println(a, b)

    result1, result2 := calculate(randN(100), randN(100))
    fmt.Println("result1=", result1, ",result2=", result2)

    printVar()

    initVar()

    shortInitVar()

    fundamentalType()

    typeTransfer()

    typeInfer()

    constValue()

    constValue2()



}
// 函数 参数 返回值
func randN(n int) int {
    return rand.Intn(n)
}

// 函数多参数
func add(x int, y int) int {
    return x + y
}

// 函数多参数 省略
func addV2(x, y int) int {
    return x + y
}

// 函数 多值返回
func swap(a, b string) (string, string) {
    return b, a
}

// 命名返回值
func calculate(x, y int) (multi, divide int) {
    multi = x * y
    divide = x / y
    return
}

// 变量
var defaultBool bool
func printVar() {
    var defaultInt int
    fmt.Println(defaultBool, defaultInt)
}

// 变量初始化
var i, j int = 3, 4
func initVar() {

    var a, b, c = 1, true, "hello"

    fmt.Println(a, b, c, i, j)
}

// 短变量声明
// * 只允许在函数内使用
func shortInitVar() {
    a, b, c := "hello", "world", "!"

    fmt.Println(a, b, c)
}


// 基本类型
// 没被初始化会被赋予零值
// 包名一样也可以分组
var (

    BoolValue bool = true
    StringValue string = "string"
    // int,uint, uintptr 在64位系统中为64位，32位系统中32位
    // int8 int16 int32 int64
    // rune, int32别名 表示一个 Unicode 码点
    IntValue int = 64
    // uint uint8 uint16 uint32 uint64 uintptr
    // unit8别名
    MaxIntValue uint64 = 1 << 64 - 1
    // float32 float64
    FloatValue float64 = 1.23
    // complex64 complex128 
    ComplexValue complex128 = cmplx.Sqrt(-5 + 12i)

)
func fundamentalType() {

    fmt.Printf("%T = %v\n", BoolValue, BoolValue)
    fmt.Printf("%T = %v\n", StringValue, StringValue)
    fmt.Printf("%T = %v\n", IntValue, IntValue)
    fmt.Printf("%T = %v\n", MaxIntValue, MaxIntValue)
    fmt.Printf("%T = %v\n", FloatValue, FloatValue)
    fmt.Printf("%T = %v\n", ComplexValue, ComplexValue)

}

// 类型转换
func typeTransfer() {
    x, y := 3, 4
    // 必须显示转换类型
    z := float64(math.Sqrt(float64(x*x + y*y)))
    f := uint64(z)
    fmt.Printf("x type=%T, value=%v\n", x, x)
    fmt.Printf("y type=%T, value=%v\n", y, y)
    fmt.Printf("z type=%T, value=%v\n", z, z)
    fmt.Printf("f type=%T, value=%v\n", f, f)

}

// 类型推导
func typeInfer() {
    x := 1
    y := 1.2345
    z := 1 + 10i
    fmt.Printf("x type = %T\n", x)
    fmt.Printf("y type = %T\n", y)
    fmt.Printf("z type = %T\n", z)
}

// 常量
// 常量不能使用:=
func constValue() {
    const Hello = "Hello"
    fmt.Println(Hello, "world")
}

// 数值常量
// 一个未指定类型的常量由上下文来决定其类型。
const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(n int) int {
    return n
}

func needFloat(n float64) float64 {
    return n
}

func constValue2() {
    fmt.Printf("%T , %v\n", needInt(Small), needInt(Small))
    fmt.Printf("%T , %v\n", needFloat(Small), needFloat(Small))
    fmt.Printf("%T , %v\n", needFloat(Big), needFloat(Big))
    // 编译无法通过  constant 1267650600228229401496703205376 overflows int
    // fmt.Printf("%T , %v\n", needInt(Big), needInt(Big))
    // 编译无法通过  constant 1267650600228229401496703205376 overflows int
    // x := 1 << 100
    // fmt.Printf("%T , %v\n", needInt(Small), needInt(Small))
}





