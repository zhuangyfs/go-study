package main

//类型转换
//func main()  {
//	var num1 int = 1;
//	var num2 float64 = 2.3;
//
//	fmt.Println(num1 + int(num2));
//}

//常量

//func main()  {
//	const a int = 1
//	fmt.Println(a)
//}

func main() {
	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam" // 允许
	customName = defaultName // 不允许

}