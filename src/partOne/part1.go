package main
import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

var customVal1,customVal2,customVal3 bool= true,false,true;//批量声明并赋值
//成组声明并赋值
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
//default值测试
var (
	defaultint int;
	defualtstr string;
	defaultfloat float32;
	defaultbool bool;
)

const mychangliang int = 144

const (
	Big = 1<<100;//常亮中声明的值默认为高精度float型
	Small = Big>>99;
) 
func main(){
	fmt.Println("我的number是",rand.Intn(200));
	var i = 7;
	var str = fmt.Sprintf("%d的平方根值为%g",i,math.Sqrt(float64(i)));
	fmt.Println(str);
	//fmt.Println(math.pi);//报错，因为math中不存在pi
	fmt.Println(math.Pi);
	var a = 1;
	var b = 3;
	fmt.Println(fmt.Sprintf("a + b 的值为%d",gmlAdd(a,b)));

	var str1,str2 = makeStrs();
	fmt.Println(str1,str2);

	arg1,arg2 := makeArgs();//注意 := 只可以在函数func内使用，用来代替var关键字，而在func外使用会编译错误
	fmt.Println(fmt.Sprintf("arg1=%d,arg2=%d",arg1,arg2));
	
	fmt.Println(customVal1,customVal2,customVal3);

	fmt.Printf("Type: %T Value: %v\n", gint32, gint32);//输出某个变量的类型和值

	fmt.Printf("%T的默认值为%v,\n%T的默认值为%v,\n%T的默认值为%v,\n%T的默认值为%v\n",defaultint,defaultint,defualtstr,defualtstr,defaultfloat,defaultfloat,defaultbool,defaultbool)

	fmt.Printf("常量定义%d\n",mychangliang);

	fmt.Println(Small * 10 + 1)
	fmt.Println(Small * 0.1)
	fmt.Println(Big * 0.1)
}

func gmlAdd(arg1,arg2 int)int{
	return arg1 + arg2;
}

func makeStrs()(string,string){
	return "我是str1","我是str2";
}

func makeArgs()(x,y int){
	// 在返回值中定义变量，也可以在函数内部直接使用，并且return不需要明确表示要返回谁,编译器自动识别返回值
	x = 14;
	y = 255;
	return ;
}