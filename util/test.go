package util

import (
	"fmt"
	"runtime"
	"sync"
)

func DeferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func PaseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//fmt.Println(stus)
	for _, stu := range stus {
		//fmt.Printf("%v",stu)
		m[stu.Name] = &stu
		//range返回值拷贝的每次指向最后一个地址，所以都是一样的
	}

	for k, v := range m {
		fmt.Printf("%s ,%v", k, v)
	}

	//答：输出的均是相同的值：&{wang 22}
	//解析 因为for遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝
	//故m[stu.Name]=&stu实际上一致指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝。
	//形同如下代码：

	/*
		var stu student
		for _, stu = range stus {
			m[stu.Name] = &stu
		}

		//修正方案，取数组中原始值的指针：
		for i, _ := range stus {
			stu:=stus[i]
			m[stu.Name] = &stu
		}
	*/
}

func RangeCloure() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i---: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// 将随机输出数字，但前面一个循环中并不会输出所有值。
	/*
		2017年7月25日：将GOMAXPROCS设置为1，将影响goroutine的并发，后续代码中的go func()相当于串行执行。

		两个for循环内部go func 调用参数i的方式是不同的，导致结果完全不同。这也是新手容易遇到的坑。

		第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。

		第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址
	*/

	/*
		i---:  9
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i:  10
		i---:  0
		i---:  1
		i---:  2
		i---:  3
		i---:  4
		i---:  5
		i---:  6
		i---:  7
		i---:  8
	*/
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Extends() {
	t := Teacher{}
	t.ShowA()
	/*
		Go中没有继承！ 没有继承！没有继承！是叫组合！组合！组合！
		这里People是匿名组合People。
		被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法，
		但他们的方法(ShowA())调用时接受者并没有发生变化。
		这里仍然是People。毕竟这个People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
		因此这里执行t.ShowA()时，在执行ShowB()时该函数的接受者是People，而非Teacher。具体参见官方文档
	*/
}

func SelectPanic() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

	/**
	有可能触发异常，是随机事件。
	单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
	select 中只要有一个case能return，则立刻执行。
	当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
	如果没有一个case能return则可以执行”default”块。
	此考题中的两个case中的两个chan均能return，则会随机执行某个case块。故在执行程序时，有可能执行第二个case，触发异常。
	*/
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func DeferInt() {
	a := 1                               //line 1
	b := 2                               //2
	defer calc("1", a, calc("10", a, b)) //3
	a = 0                                //4
	defer calc("2", a, calc("20", a, b)) //5
	b = 1                                //6

	/*
		在解题前需要明确两个概念：

		defer是在函数末尾的return前执行，先进后执行，具体见问题1。
		函数调用时 int 参数发生值拷贝。

		不管代码顺序如何，defer calc func中参数b必须先计算，故会在运行到第三行时，执行calc("10",a,b)输出：10 1 2 3得到值3，将cal("1",1,3)存放到延后执执行函数队列中。
		执行到第五行时，现行计算calc("20", a, b)即calc("20", 0, 2)输出：20 0 2 2得到值2,将cal("2",0,2)存放到延后执行函数队列中
		执行到末尾行，按队列先进后出原则依次执行：cal("2",0,2)、cal("1",1,3) ，依次输出：2 0 2 2、1 1 3 4 。
	*/
}

func MakeSlice() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	/*
		make可用于初始化数组，第二个可选参数表示数组的长度。数组是不可变的。

		当执行make([]int,5)时返回的是一个含义默认值(int的默认值为0)的数组:[0,0,0,0,0]。而append函数是便是在一个数组或slice后面追加新的元素，并返回一个新的数组或slice。

		这里append(s,1,2,3)是在数组s的继承上追加三个新元素:1、2、3，故返回的新数组为[0 0 0 0 0 1 2 3]
	*/
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

/**
答： 在执行 Get方法时可能被panic

解析

虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。会报错误信息:“fatal error: concurrent map read and map write”。

可以在在线运行中执行，复现该问题。那么如何改善呢? 当然Go1.9新版本中将提供并发安全的map。首先需要了解两种锁的不同：

sync.Mutex互斥锁
sync.RWMutex读写锁，基于互斥锁的实现，可以加多个读锁或者一个写锁。
利用读写锁可实现对map的安全访问，在线运行改进版 。利用RWutex进行读锁。
*/

/*
type RWMutex
func (rw *RWMutex) Lock()
func (rw *RWMutex) RLock()
func (rw *RWMutex) RLocker() Locker
func (rw *RWMutex) RUnlock()
func (rw *RWMutex) Unlock()
*/
/*
type threadSafeSet struct {
	//sync.Mutex
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
*/

/**
答： 内部迭代出现阻塞。默认初始化时无缓冲区，需要等待接收者读取后才能继续写入。
解析
chan在使用make初始化时可附带一个可选参数来设置缓冲区。默认无缓冲，题目中便初始化的是无缓冲区的chan，这样只有写入的元素直到被读取后才能继续写入，不然就一直阻塞。

设置缓冲区大小后，写入数据时可连续写入到缓冲区中，直到缓冲区被占满。从chan中接收一次便可从缓冲区中释放一次。可以理解为chan是可以设置吞吐量的处理池。

来自社区fiisio的说明
ch := make(chan interface{}) 和 ch := make(chan interface{},1)是不一样的

无缓冲的 不仅仅是只能向 ch 通道放 一个值 而是一直要有人接收，那么ch <- elem才会继续下去，要不然就一直阻塞着，也就是说有接收者才去放，没有接收者就阻塞。

而缓冲为1则即使没有接收者也不会阻塞，因为缓冲大小是1只有当 放第二个值的时候 第一个还没被人拿走，这时候才会阻塞
*/

type Tpeople interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func Tstruct() {
	var peo Tpeople = &Stduent{}
	//var peo Tpeople = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

/**
 编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型。

解析

考题中的 func (stu *Stduent) Speak(think string) (talk string) 是表示结构类型*Student的指针有提供该方法，但该方法并不属于结构类型Student的方法。因为struct是值类型。

修改方法：

定义为指针 go var peo People = &Stduent{}
方法定义在值类型上,指针类型本身是包含值类型的方法。 go func (stu Stduent) Speak(think string) (talk string) { //... }
*/

/*
type Fpeople interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func Ttype() {
	if live() == nil { //cannot convert nil to type people
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}


	# golang-learning/util
	util\test.go:331: cannot use stu (type *Student) as type People in return argument
	util\test.go:335: cannot convert nil to type People

}
*/
