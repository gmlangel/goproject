package main

import(
	"fmt"
	"math/cmplx"
	"math"
	"runtime"
)
var (
	gint int= 2;//在32为系统中长度为int32，在64位系统中长度为int64
	gint8 int8 = 255 / 2;
	gint16 int16 = 65535 / 2;
	gint32 int32 = 0xffffffff / 2;
	gint64 int64 = (1<<64) / 2 - 1;
	guint uint = 2;//在32为系统中长度为uint32，在64位系统中长度为uint64
	guint8 uint8 = 255;
	guint16 uint16 = 0xffff;
	guint32 uint32 = 0xffffffff;
	guint64 uint64 = 1<<64 - 1;
	gbyte byte = 255;//等于一个uint8
	grune rune = 0xffffffff / 2;//等于1个int32
	gfloat32 float32 = 0xffffffff;
	gfloat64 float64 = 1<<64 - 1;
	guintptr uintptr = 10;//不知道做什么用,在32为系统中长度为uint32，在64位系统中长度为uint64
	gcomplex128 complex128 = cmplx.Sqrt(-5 + 12i)
)
func main(){
	var count = 0;
	for i:=0;i<10;i++{
		count +=i;
	}
	fmt.Println("for循环计算结果为",count);
	// //手写一个无限循环,非常简洁
	// for{

	// }
	var sum = 1;
	for sum < 1000{
		sum += sum;
	}
	fmt.Println("go语言中while关键字被for代替，sum的值为",sum);

	// var targ1 = 1.010
	// var targ2 = 2;
	// var targ3 = "i"
	//fmt.Println("两种不同的值类型相加结果为",targ1 + targ2);//报错 mismatched types float64 and int
	//fmt.Println("两种不同的值类型相加结果为",targ1 + targ3);//报错 mismatched types float64 and string

	fmt.Println(sqrt(2), sqrt(-4));

	//在if条件判断中添加 变量声明，则在if表达式，代码块以及else代码块中均可使用该变量
	if tc:=43;tc < 24{
		fmt.Println("进入了判断,tc=",tc)
	}else{
		fmt.Println("进入了else,tc=",tc)
	}

	switch os:=runtime.GOOS;os{
	case "darwin":
		fmt.Println("当前系统为Mac OS");
	case "linux":
		fmt.Println("当前系统为 Linux");
	default:
		fmt.Println("当前系统为 ",os);
	}

	testDefer();

	testDeferFor();
}

func sqrt(x float64)string{
	if(x < 0){
		return sqrt(-x) + "i";
	}
	return fmt.Sprint(math.Sqrt(x));
}

func testDefer(){
	fmt.Println("defer测试:1")
	defer fmt.Println("defer测试:2")
	fmt.Println("defer测试:3")
	defer fmt.Println("defer测试:4")
	fmt.Println("defer测试:5")
}

func testDeferFor(){
	fmt.Println("defer的循环测试开始")
	for i:= 0;i<10;i++{
		defer fmt.Println("defer循环",i);
	}
}