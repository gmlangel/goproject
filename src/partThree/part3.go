package main

import(
	"fmt"
)
/**
结构体声明
*/
type GMLStruct struct{
	x int;
	y int;
}

var(
	arg1 = GMLStruct{10,13};
	arg2 = GMLStruct{x:14};//输定属性赋值，未指定的属性依旧为默认值
	arg3 = GMLStruct{};
	parg = &GMLStruct{60,66};
)

func main(){
	i:= 42;
	var ptr = &i;
	fmt.Printf("ptr的类型%T,ptr的值为%v\n",ptr,ptr);
	fmt.Println("ptr指针对应的值为",*ptr);

	//结构体测试

	var mystruct GMLStruct = GMLStruct{13,14};
	fmt.Printf("mystruct的类型为%T,值为%v\n",mystruct,mystruct);
	
	fmt.Printf("mystruct.x=%d,mystruct.y=%d\n",mystruct.x,mystruct.y);

	structPtrTest();

	fmt.Println(arg1,arg2,arg3,*parg);

	//不可变长数组
	var arr1 [6]int = [6]int{1,2,3,4,5,6}
	arr:= [6]int{1,2,3,4,5,6};
	fmt.Println("arr值为",arr,"arr1值为\n",arr1);
	//变长数组
	var arr2 []int = arr[1:4];//数组切面，其实索引为1，终止索引为4，截取1-4之前不含4的数据
	fmt.Println("arr2值为",arr2);

	arr2[0] = 14//切片只是对原数组的引用，这一点一定要和其它语言分清（其它语言是潜复制），一但改变切片的某个值，会直接影响原数组中索引对应的值
	fmt.Println("arr值为",arr);

	//创建struct数组
	var structArr []GMLStruct = []GMLStruct{GMLStruct{1,2},GMLStruct{3,4}};
	fmt.Println("struct数组测试,值为",structArr);
}

//结构体的指针用法
func structPtrTest(){
	gmlstruc:= GMLStruct{22,23};
	gmlstructPtr:= &gmlstruc;
	fmt.Println("正常情况下调用gmlstructPtr或其下的属性，需要用*gmlstructPtr，比如：");
	(*gmlstructPtr).x = 40;
	(*gmlstructPtr).y = 50;
	fmt.Printf("输出gmlstruc=%v\n",*gmlstructPtr);
	fmt.Println("使用go语言的语法糖可以省去*取值之后调用属性，比如：");
	gmlstructPtr.x = 44;
	gmlstructPtr.y = 55;
	fmt.Printf("出书gmlstruc=%v\n",*gmlstructPtr);
}