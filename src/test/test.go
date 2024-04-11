package main

import (
   "fmt"
   "./myMath"
   "unsafe"
)

//ddd := 6  //这种不带声明格式的只能在函数体中出现
var eee = 6

var (
   a1 int
   b1 bool
)

func numbers()(int,int,string){
   a, b, c := 1, 2, "str"
   return a,b,c
}

//常量枚举
const (
   Unkown = 0
   Female = 1
   Male = 2
)

const (
   a4 = "abc"
   b4 = len(a4)
   c4 = unsafe.Sizeof(a4)
)

func main() {
   //go run test.go  //直接执行
   //go build test.go 生成二进制文件，再执行 ./test

   fmt.Println("Hello, World!")  // 等同于 fmt.Print("hello, world\n") 
   fmt.Println([3]int{1, 2, 3})


   fmt.Println(mathClass.Add(1,1))
   fmt.Println(mathClass.Sub(1,1))

   fmt.Println("【变量声明】")
   var a string = "hello" //1.指定变量类型，如果没有初始化，则变量默认为零值。
   var aa = "world"       //2.根据值自行判定变量类型。
   aaa := "hi"            //3.直接用 := 声明变量 
                          //  这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
   fmt.Println(a, aa, aaa) // hello world hi

   var b, c int = 1, 2
   var bb, cc = 3, 4
   bbb := 5
   ccc := 6
   fmt.Println(b, c, bb, cc, bbb, ccc, eee) // 1 2 3 4 5 6 6
   //fmt.Println(ddd)

   b, c = c, b //交换两个变量的值
   fmt.Println(b, c) // 2 1

   //如果你声明了一个局部变量却没有在相同的代码块中使用它，会得到编译错误，但是全局变量是允许声明但不使用的。

   //空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
   //_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
   var b2 int;
   _, b2 = 5, 7
   fmt.Println(b2) // 7

   _,numb,strs := numbers() //只获取函数返回值的后两个
   fmt.Println(numb,strs)   // 2 str

   fmt.Println("【常量声明】")
   const LENGTH int = 10
   const WIDTH int = 5
   var area int
   const a3, b3, c3 = 1, false, "str"

   area = LENGTH + WIDTH
   fmt.Printf("面积为：%d", area)  //面积为：15
   println()
   println(a3, b3, c3)  //1 false str

   println(Female, Male) // 1 2
   println(a4, b4, c4)   //abc 3 16

   //iota 特殊常量 todo

   //注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

   var a5 int = 10
   if a5 < 20 {
      fmt.Printf("a 小于 20\n")
   } else {
      fmt.Printf("a 不小于 20\n")
   }

   //switch 和 select todo

   fmt.Println("【for循环】")
   sum := 0
   for i := 0; i <= 10; i++ {
      sum += i
   }
   fmt.Println(sum) //55
   println(sum)   //55

   sum1 := 1
   //for ; sum1 <= 10; {
   for sum1 <= 10 {
      sum1 += sum1
   }
   fmt.Println(sum1) //16

   strings := []string{"google", "runoob"}
   for i, s := range strings {
      fmt.Println(i, s)  //0 google
                         //1 runoob
   }

   numbers := [6]int{1, 2, 3, 5}
   for i, x := range numbers {
      fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
         //第 0 位 x 的值 = 1
         //第 1 位 x 的值 = 2
         //第 2 位 x 的值 = 3
         //第 3 位 x 的值 = 5
         //第 4 位 x 的值 = 0
         //第 5 位 x 的值 = 0
   }

   //如果只想读取 key，格式如下：
   //for key := range oldMap
   //或者这样：for key, _ := range oldMap

   //如果只想读取 value，格式如下：
   //for _, value := range oldMap

   map1 := make(map[int]float32)
   map1[1] = 1.0
   map1[2] = 2.0

   //读取 key 和 value
   for key, value := range map1 {
      fmt.Printf("key is: %d - value is: %f\n", key, value)
         //key is: 2 - value is: 2.000000
         //key is: 1 - value is: 1.000000  //顺序随机？
   }

   //读取 key
   for key := range map1 {
      fmt.Printf("key is: %d\n", key)  //顺序随机？ 
         //key is: 2
         //key is: 1
   }

   //读取 value
   for _, value := range map1 {
      fmt.Printf("value is: %f\n", value)
   }

   //九九乘法表
   // for m := 1; m < 10; m++ {
   //    /*    fmt.Printf("第%d次：\n",m) */
   //    for n := 1; n <= m; n++ {
   //       fmt.Printf("%dx%d=%d ",n,m,m*n)
   //    }
   //    fmt.Println("")
   // }
   mathClass.printMultiTable()
   // mathClass.Add(1, 2)

   //break todo
}