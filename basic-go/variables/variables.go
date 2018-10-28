package main

/**
	变量初始化
 */
//func main()  {
//	var value string
//	value = "adadadad"
//	fmt.Println("this input " + value)
//}

/*
	变量初始化赋值
 */
//func main()  {
//	var value string = "hello world"
//	fmt.Println("this input " + value)
//}

/*
	类型推断
 */
//func main()  {
//	var value = 12345
//	fmt.Println("value is ", value)
//}

/*
	多变量同类型申明
 */
//func main()  {
//	var value1, value2 int = 123, 456
//	fmt.Println("value1 ", value1, "value2 ", value2)
//}

/*
	多变量不同类型申明
 */
//func main()  {
//	var (
//		value1 = "string 123"
//		value2 = 456
//		value3 = "string 789"
//	)
//
//	fmt.Println("value1", value1, "value2", value2, "value3", value3)
//}

/*
	简短申明
	简短申明左侧需要存在未被申明过的变量，一般初始化时候使用
 */
//func main() {
//	value1, value2 := "value1 is 123", 456
//	fmt.Println("value", value1, "value2", value2)
//}

//func main() {
//
//	var num1, num2 int = 123, 234
//
//	num3 := math.Min(float64(num1), float64(num2))
//	fmt.Println(num3)
//}