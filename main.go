package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"iter"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Hello, VS Code!")

	// 练习1：
	//  编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	var i int = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "Hello!Go语言"
	fmt.Printf("变量i的值为：%v，类型为：%T\n", i, i)
	fmt.Printf("变量f的值为：%v，类型为：%T\n", f, f)
	fmt.Printf("变量b的值为：%v，类型为：%T\n", b, b)
	fmt.Printf("变量s的值为：%v，类型为：%T\n", s, s)
	//  编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	str1 := "hello沙河小王子"
	count := 0
	len1 := utf8.RuneCountInString(str1)
	for _, ch := range str1 {
		if ch >= 0x4E00 && ch <= 0x9FA5 { // 汉字的Unicode编码范围,因为汉字类型是rune,rune代表unicode码点（每个字符都有对应的码点）
			count++
		}
	}
	fmt.Printf("字符串\"%s\"总字符数为：%d\n", str1, len1)
	fmt.Printf("字符串\"%s\"中汉字的数量为：%d\n", str1, count)

	//  计算字符串每个字符在UTF-8编码下占用的字节数
	for index, char := range str1 {
		// index:字符的起始字节索引
		// char:字符本身，类型是rune(int32)

		fmt.Printf("索引: %d, 字符: %c, 码点 (rune): %d, 字节数: %d\n",
			index,
			char,
			char,
			getByteSize(char))
	}

	// 练习2：
	// 从一个只有一个数字出现一次，其他数字都出现两次的数字字符串里面找出出现一次的数字
	str2 := "a52425b"
	singleNum, found := findSingleNumberInString_XOR(str2)
	if found {
		fmt.Printf("字符串\"%s\"中只出现一次的数字是：%d\n", str2, singleNum)
	} else {
		fmt.Printf("字符串\"%s\"中没有数字\n", str2)
	}

	str2_2 := "j6767898"
	singleNUm2, found2 := findSingleNumberInString_HashMap(str2_2)
	if found2 {
		fmt.Printf("字符串\"%s\"中只出现一次的数字是：%d\n", str2_2, singleNUm2)
	} else {
		fmt.Printf("字符串\"%s\"中没有只出现一次的数字\n", str2_2)
	}

	// 练习3：九九乘法表
	fmt.Println("--- 九九乘法表 (跳过含6的项) ---")
	// 调用函数来执行打印任务
	printMultiplicationTable_Skip6()
	fmt.Println("---------------------------------")

	// 练习4：找出数组中和为指定值的两个元素的下标
	nums4 := []int{1, 8, 3, 5, 7, 0, 8}
	target4 := 8
	fmt.Println("原始数组：", nums4)
	fmt.Println("目标和:", target4)
	pairs2 := findSumPairsHash(nums4, target4)
	if len(pairs2) > 0 {
		fmt.Println(pairs2) // 输出可能为 [[0 6] [1 5] [3 2]] 顺序不固定
	} else {
		fmt.Println("未找到任何组合。")
	}

	// 练习5：输出中英混合字符串中每个单词和汉字出现的次数（重点：单词累加器）
	text := "hello 世界， Hello world! 我爱我家，爱世界 I love my home"
	fmt.Println("原始字符串:")
	fmt.Println(text)
	wordCountsMap := countWordsAndChars(text)

	// 规范化输出结果
	fmt.Println("统计结果:")
	for item, count := range wordCountsMap {
		fmt.Printf("'%s':%d\n", item, count)
	}

	// 练习6：闭包
	//创建一个说Hello的问候函数
	greetHello := makeGreeter("Hello")
	greetNihao := makeGreeter("你好")
	fmt.Println(greetHello("Alice"))
	fmt.Println(greetNihao("小名"))
	// 创建一个说“你好”的问候函数

	// 练习7：金币分配
	fmt.Println("开始进行金币分配...")
	left, err := dispatchCoin(users, coins, distribution)
	// go中最常见的代码块: if err != nil 如果err不是nil，说明发生了错误
	if err != nil {
		log.Fatalf("!!!分配失败:%v", err)
		// 使用 log.Fatalf 可以打印错误信息并以非0状态码退出程序
	}
	fmt.Println("分配成功！")
	fmt.Println("剩下：", left)

	// 练习8：面向对象学生数据库
	// 创建一个班级花名册
	classOne := NewRoster(101)

	// 添加学生
	classOne.AddStudent("Alice", 18, []string{"数学", "物理"})
	classOne.AddStudent("Bob", 19, []string{"化学", "生物"})
	classOne.AddStudent("Charlie", 18, []string{"历史", "地理"})

	// 展示学生列表
	classOne.ShowAllStudents()

	// 编辑学生信息
	// 1. 先获取要编辑的学生信息（在真实应用中，这可能是从前端传来的）
	studentToUpdate := Student{
		ID:       2, // 我们要更新 Bob
		Name:     "Robert (Bob)",
		Age:      20,
		Subjects: []string{"化学", "计算机科学"},
	}
	// 2. 调用更新方法
	err8_1 := classOne.UpdateStudent(studentToUpdate)
	if err8_1 != nil {
		fmt.Println(err8_1)
	}

	// 再次展示，查看更新结果
	classOne.ShowAllStudents()

	// 删除学生
	err8_1 = classOne.DeleteStudent(1) // 删除 Alice
	if err != nil {
		fmt.Println(err8_1)
	}

	// 再次添加学生，验证ID生成是否安全
	classOne.AddStudent("David", 17, []string{"美术", "音乐"})

	defaultsubjects := []string{"中文"}
	classOne.AddDefaultSubjectsToAll(defaultsubjects)

	// 最终展示
	classOne.ShowAllStudents()

	// 练习9：接口，依赖注入
	fmt.Println("--- 场景一:使用控制台日志记录器---")
	consoleLogger := NewConsoleLogger()                 //创建ConsoleLogger实例
	consoleUserService := NewUserService(consoleLogger) //依赖注入：将consoleLogger注入到UserService中
	consoleUserService.CreateUser("Alice")

	fmt.Println("--- 场景二：使用文件日志记录器---")
	fileLogger, err := NewFileLogger("exercise9_application.log")
	if err != nil {
		//如果日志文件都无法创建，程序无法正常运行，直接退出
		log.Fatalf("无法创建文件日志记录器: %v", err)
	}

	defer fileLogger.close()                      // 确保在main函数退出前，文件句柄一定会被关闭
	fileUserService := NewUserService(fileLogger) // 依赖注入:将fileLogger注入到另一个UserService实例中
	// 调用业务方法
	fileUserService.CreateUser("Bob")
	fileUserService.DeleteUser("Alice")

	fmt.Println("\n操作完成。请检查与本程序同目录下的 'application.log' 文件。")

	// 练习十：goroutine和channel并发，生成100个随机数，让24个并发程序去执行它
	const numJobs = 100
	const numWorkers = 24

	jobChan := make(chan int64, numJobs)     // 创建存储任务的jobchannel，带缓冲,意味着producer可以一口气将所有100个任务全部发送到jobchan而不会被阻塞，即便没有worker立即开始处理
	resultChan := make(chan string, numJobs) // 创建存储结果的resultchannel，带缓冲

	// 1.创建一个带有1秒超时的context
	// context.WithTimeout 返回一个新的context和一个cancel函数
	// defer cancel()是必须的，能确保在main函数退出时，所有与此context相关的资源都被释放
	// context本质上是一个接口，我们通常从一个空的context.Background()开始，像套娃一样，用context.WithCancel, context.WithTimeout或context.WithDeadline派生出新的，带有特定功能的context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 2.使用errgroup.WithContext创建一个与我们的context绑定的group
	// g会和waitGroup一样，管理goroutine的生命周期
	// errgroup包的withContext函数接收一个父context，返回一个新的Group对象（类似waitgroup）和一个子context对象
	g, gCtx := errgroup.WithContext(ctx)

	// 3.启动生产者，作为group的一部分
	// errgroup.Gropu提供了Go和Wait两个方法，这两个方法需提供一个返回err的函数
	// Go函数会在新的goroutine中调用传入的函数f，第一个返回非零错误的调用将取消该Group， 下面的Wait方法会返回该错误
	// Wait会阻塞直至上述Go方法调用的所有函数都返回，然后从它们返回第一个非nil的错误
	g.Go(func() error {
		return producer(gCtx, jobChan, numJobs)
	})

	// 4.启动所有工人，作为group的一部分
	for i := 1; i <= numWorkers; i++ {
		workerID := i // 闭包问题：必须在循环内创建局部变量
		g.Go(func() error {
			return worker(gCtx, workerID, jobChan, resultChan)
		})
	}

	// 5.启动一个专门的goroutine来等待所有生产者和工人完成
	// 同样需要一个goroutine在所有worker完成后关闭resultChan
	go func() {
		g.Wait()          // 阻塞当前这个匿名goroutine,直到WaitGroup的计数器变为0
		close(resultChan) // 当wg.wait()返回时，意味着所有的worker都已调用Done()并退出，此刻不会再有任何数据被写入ResultChan，关闭它
	}()

	fmt.Println("主程序：开始接收结果，设置1秒超时...")
	timeout := time.After(1 * time.Second) // 一次性定时器
	// time.After(duration）在调用时，会立即返回一个channel(chan time.Time),即Go 的运行时会在后台为你启动一个计时器
	// 当duration（这里是1秒）的时间过去后，Go运行时会自动地向这个channel里发送一个值（当前的时间）
	// timeout变量的本质是： 一个在未来某个时间点（1s后）才会受到数据的特殊channel

	var resultsCollected int = 0 // 用于在超时发生时报告已完成了多少个工作

	// select多路复用：它会阻塞，直到其下的某一个case的channel操作准备就绪
	// ResultsLoop标签用于从外层循环中跳出，在并发程序中监听多个通道事件，并在特定条件下退出循环而不返回
ResultsLoop:
	for {
		select {
		case result, ok := <-resultChan:

			if !ok {
				fmt.Println("所有结果已成功接收")
				break ResultsLoop
			}

			fmt.Println(result)
			resultsCollected++
		// 如果channel未关闭且有数据，程序打印结果并增加计数器。如果channel已被关闭且无数据（正常完成的信号），程序打印成功信息，并通过break跳出循环

		// 主循环超时与context（监听所有的goroutine）超时分开
		// timeout让主程序知道1秒到了，context让所有worker和producer构成的goroutine检测到一秒到了并停止
		case <-timeout:
			cancel() // 主动取消所有goroutine
			fmt.Printf("\n!!! 处理超时，1秒内收集到%d/%d个结果\n", resultsCollected, numJobs)
			break ResultsLoop
			// timeout channel有数据了，就代表超时了

		case <-gCtx.Done(): //新增：监听errgroup的取消信号
			fmt.Printf("\n程序因错误而终止,收集到%d/%d个结果\n", resultsCollected, numJobs)
			break ResultsLoop

		}
	}
	fmt.Println("\n--- 继续执行后续练习 ---")

	// 练习十一：订单系统筛选迭代器
	// 准备一些模拟数据
	mgr := OrderManager{
		orders: []Order{
			{ID: 1, Amount: 50.0, Status: "Pending"},  // 金额太小，应该被跳过
			{ID: 2, Amount: 200.0, Status: "Paid"},    // 状态不对，应该被跳过
			{ID: 3, Amount: 150.0, Status: "Pending"}, // 符合条件！
			{ID: 4, Amount: 300.0, Status: "Pending"}, // 符合条件！
			{ID: 5, Amount: 20.0, Status: "Pending"},  // 金额太小
			{ID: 6, Amount: 500.0, Status: "Pending"}, // 符合条件，但我们可能不需要处理这么多
		},
	}
	fmt.Println("---开始处理大额待支付订单---")

	// 使用迭代器:调用之前定义的迭代器方法
	// 注意这里代码非常干净，主函数完全不知道筛选逻辑是怎样的
	var orderCount int = 0
	for order := range mgr.BigPendingOrders(100) {
		orderCount++
		fmt.Printf("处理订单ID:%d, 金额:%.2f, 状态:%s\n", order.ID, order.Amount, order.Status)
		// 模拟：我们只处理前两个就够了
		if orderCount >= 2 {
			fmt.Println(">>任务已达标，停止处理")
			break
		}
	}
	fmt.Println("---流程结束---")

	// 练习十二:日志文件处理器
	// 使用flag包定义命令行参数
	sourceFile := flag.String("src", "", "Source Log file path(required)")
	destFile := flag.String("dst", "", "Destination Log file path(required)")
	var keyword string
	flag.StringVar(&keyword, "key", "ERROR", "Keyword to filter by") //两种不同的命令行参数方式
	flag.Parse()                                                     // 解析命令行参数 必须要

	// 参数校验
	if *sourceFile == "" || *destFile == "" {
		fmt.Println("错误:源文件路径和目标文件路径都是必需的")
		flag.Usage() // 打印帮助信息
		os.Exit(1)   // 以非零状态码退出，表示错误
	}

	err12 := processLogFile(*sourceFile, *destFile, keyword)
	if err12 != nil {
		// 使用log.Fatalf打印错误并退出程序
		log.Fatalf("处理文件时发生知名错误：%v", err12)
	}
	fmt.Println("日志文件处理完成")

}

// 练习1 辅助函数，计算一个rune在UTF-8编码下占用的字节数
func getByteSize(r rune) int {
	if r <= 0x7F {
		return 1
	} else if r <= 0x7FF {
		return 2
	} else if r <= 0xFFFF {
		return 3
	} else {
		return 4
	}
}

// 练习2 辅助函数，返回数字字符串里只出现一次的数字
// 首先是字典哈希表算法
func findSingleNumberInString_HashMap(s string) (int, bool) {
	countsmap := make(map[rune]int) //创建一个map来存储每个数字字符的出现次数，key是rune类型，value是int类型
	//遍历字符串中每一个rune
	for _, r := range s {
		if unicode.IsDigit(r) { //判断r是否是数字字符
			countsmap[r]++ //如果是数字字符，则在map中对应的key的value加1
		}
	}

	//遍历map，找到只出现一次的数字字符
	for r, count := range countsmap {
		r_int := int(r - '0') //将rune类型的数字字符转换为整数数字
		if count == 1 {
			return r_int, true //返回只出现一次的数字和true表示找到了
		}

	}
	return 0, false //如果没有找到只出现一次的数字，返回0和false
}

// 相比字典哈希表算法，使用异或运算符(XOR)空间复杂度更低
// XOR核心原理：相同的数字异或结果为0，0与任何数字异或结果为该数字本身，不同的数字异或结果为1
//
//	XOR还有一个重要性质：运算满足交换律和结合律
//	即：按下面的代码，（52542）循环连续异或，可以看作(5^5)^(2^2)^4 = 0^0^4 = 4
//
// 但下面这个函数在“不保证一定有解”的情况下有歧义，错误时无法返回“没找到只出现一次的数字”这一信息
func findSingleNumberInString_XOR(s string) (int, bool) {
	result := 0
	foundAnyDigit := false //一个标志，用于判断字符串中是否有数字
	for _, r := range s {
		if unicode.IsDigit(r) {
			foundAnyDigit = true
			// 将字符数字转换为整数数字
			// '4‘的ASCII码减去’0‘的ASCII码正好等于4
			num := int(r - '0')
			result = result ^ num
		}
	}
	if foundAnyDigit {
		return result, true
	}
	return 0, false
}

// 练习3 辅助函数：按照每行增加的样式，打印一个跳过包含数字6的九九乘法表
func printMultiplicationTable_Skip6() {
	// 外层循环控制行（第一个乘数i）
	for i := 1; i <= 9; i++ {
		// 内层循环控制列（第二个乘数j）
		for j := 1; j <= i; j++ {
			product := i * j

			//将两个乘数转换为str
			iStr := strconv.Itoa(i)
			jStr := strconv.Itoa(j)

			// 检查i,j的字符串形式中是否包含"6"
			if strings.Contains(iStr, "6") || strings.Contains(jStr, "6") {
				continue // break会终止整个内层循环，如果这儿用break,那么直接跳到下一个i了
			}

			// 使用Printf进行格式化输出
			// %-2d用于乘法结果的占位，负号表示左对齐，占位2个字符宽度，可以让结果更整齐
			// \t表示制表符，用于在式子之间产生较大的固定间距
			fmt.Printf("%dx%d=%-2d\t", j, i, product)

		}
		// 每当一行（内层循环结束），就打印一个换行符
		if i != 6 {
			fmt.Println()
		}
	}
}

// 练习4 辅助函数：找出数组和中为指定值的两个元素的下标
// 一开始的想法：暴力查找法，两层遍历所有的组合，外层0到n，内层i+1到n，如果和为指定值就把(i,j)存入索引切片中
// 但这种方法时间复杂度是O(n^2)，效率较低
// 更优的做法是使用哈希表，时间复杂度和空间复杂度都为O(n)
func findSumPairsHash(arr []int, target int) [][]int {
	// args: arr是输入的整数数组，target是目标和
	// return: 返回一个二维切片，包含所有符合条件的索引对
	var result [][]int

	// 1.创建一个map，key是数字，value是索引
	seen := make(map[int]int)

	// 2.单层循环遍历数组
	for i, num := range arr {
		// 计算需要的“另一半”
		complement := target - num

		// 在map中查找complement是否存在
		if j, found := seen[complement]; found {
			// 如果找到了，说明nums[j] + nums[i] == target
			// 我们就找到了一个解(j,i)
			result = append(result, []int{j, i})
		}

		//无论找没找到，都把当前数字和它的索引存入map，供后面元素进行查找
		seen[num] = i
	}
	return result
}

// 练习5 辅助函数 统计一个中英文混合字符串中，每个英文单词和中文汉字的出现次数
// 代码实现策略
//   1.统一的计数器：创建一个map[string]int 来存储所有结果，因为无论是单词还是汉字都可以用string表示，无需分开创建字典
//   2.单词累积器（🔛）：使用strings.Builder来高效地拼接英文字符，形成单词
//   3.逐Rune遍历：使用for _,r := range str 遍历字符串
//   4.分类处理：
//      - 如果r是汉字（unicode.IsHan):
//        - 首先，检查单词累积器(strings.Builder)中是否有内容，如果有，说明一个英文档次刚结束，需要将其存入map并清空累积器
// 	  - 将当前汉字作为string直接存入map中
// 	- 如果r是字母：
// 	  - 将这个字母rune添加到单词累积器中
// 	- 如果r是其他字符（空格，标点等）：
// 	  - 把它看作是一个“单词边界”，检查单词累积器中是否有内容，如果有，处理它并清空累积器
//  5.收尾工作：循环结束后，单词累积器中可能还留着最后一个英文单词（如果字符串以英文单词结尾），需要做一次最终的检查与处理

func countWordsAndChars(s string) map[string]int {
	// 1.使用一个map统一存储结果
	counts := make(map[string]int)

	// 2.使用strings.Builder 相比加号（s = s +"a" + "b")，更能高效累积英文单词
	var wordBuilder strings.Builder

	// 定义一个内部函数/闭包，用于处理累积器的单词，避免代码重复
	processWord := func() {
		word := wordBuilder.String()
		if word != "" { //如果单词累积器里不为空
			counts[strings.ToLower(word)]++ // 将单词转为小写，实现不区分大小写记数
			wordBuilder.Reset()             // 清空累积器

		}
	}

	// 3.逐个Rune遍历字符串并分类处理
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			processWord()       // 遇到汉字代表单词累积结束，需要处理一下累积器里的单词了
			counts[string(r)]++ // 强制类型转换并充当字典的key

		} else if unicode.IsLetter(r) {
			wordBuilder.WriteRune(r) // 如果是字母，加入单词累积器
		} else {
			processWord() //如果是空格，标点等，视为单词边界，处理累积的英文单词
		}
	}

	processWord() // 收尾工作

	return counts
}

// 练习六：闭包函数
// 闭包就是一个函数和它能访问的环境变量（在它外面定义的变量）的组合体。即使外部函数已经执行完了，这个内部函数仍然能记住并操作那些外部变量
// 闭包可以用来创建一系列功能相似但是配置不同的函数
// makeGreeter是一个函数工厂，它接受一个前缀（汉语名或英语名），然后生产出汉语/英语的问候函数
func makeGreeter(prefix string) func(string) string {
	//返回的函数是一个闭包，它捕获了外部变量prefix
	return func(name string) string {
		return prefix + "," + name
	}
}

// 练习七：分配金币
// 你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
// 分配规则如下：
// a. 名字中每包含1个'e'或'E'分1枚金币
// b. 名字中每包含1个'i'或'I'分2枚金币
// c. 名字中每包含1个'o'或'O'分3枚金币
// d: 名字中每包含1个'u'或'U'分4枚金币
// 写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

//思路：
// 1.在dispatchCoin函数开始时，定义一个分配金币总数的变量
// 2.在函数开头使用defer安排一个匿名函数，用于在最后打印出完整的distribution map，作为一份详细的分配报告
// 3,开始遍历users切片，对于每一个name:
//    .定义一个变量personCoins用于计算当前这个人的金币数，初始化为0
//    .遍历name字符串中的每一个字符(rune) 这两个遍历都用range遍历
//    .对于每一个字符，使用switch语句来判断，是什么字母，加多少金币
//    .内层循环（字符遍历）结束后，personCoins就是这个人应得的金币总数，存入字典，并更新已分配的金币总数
// 4.外层循环（用户遍历）结束后，计算剩余金币，返回剩余值

func dispatchCoin(users []string, coins int, distribution map[string]int) (int, error) {
	//定义一个累计已分配金币的变量
	totalDistributed := 0

	//使用defer在函数退出前打印最终的分配详情 可以把defer当作“稍后处理”的便签
	defer func() {
		fmt.Println("===== 分配详情 =====")
		for name, amount := range distribution {
			fmt.Printf("%s:%d\n", name, amount)
		}
		fmt.Println("=========")
	}() // 最后这个括号是函数调用操作符，意思是：执行这个函数 defer函数的完整定义：defer func() { ... }()

	// 1.遍历所有用户
	for _, name := range users {
		personalCoins := 0 // 定义个人的分配数量
		// 2.遍历当前用户名的每一个字符
		for _, rune := range name {
			switch rune {
			case 'e', 'E':
				personalCoins += 1
			case 'i', 'I':
				personalCoins += 2
			case 'o', 'O':
				personalCoins += 3
			case 'u', 'U':
				personalCoins += 4
			}
		}
		// 3.将计算出的个人金币数存入字典
		distribution[name] = personalCoins
		// 4.累加到总分配金币数中
		totalDistributed += personalCoins
	}

	// 以下是错误处理：如果分配的金币超过了金币总数
	// 这里可以用panic和recover，但是会被认为是滥用
	//  panic被设计用来处理真正灾难性的，程序无法继续正常运行的错误，典型例如数组越界，空指针引用
	// 这里可以用error返回值，error被设计用来处理可预期的，业务逻辑范围内的失败情况，典型例如格式错误，打开不存在文件，网络超时等
	//   error是程序正常运行中可能遇到的“非成功”状态
	if totalDistributed > coins {
		// 创建一个error对象
		err := fmt.Errorf("金币不足！需要%d,但只有%d", totalDistributed, coins)
		return 0, err //这个大函数需要定义一个error返回类型，出错返回error对象，不出错返回nil
	}
	return coins - totalDistributed, nil
}

// 练习八：面向对象的学生数据管理系统具体逻辑详解
// 1. 数据结构 (Student 和 Roster)
//    Student 结构体:
//      逻辑: 这是一个纯粹的“数据容器”或“模型”，定义了一个学生应该具备哪些属性（ID, Name, Age, Subjects）。它简单、直接，不包含任何业务逻辑。
//      代码: type Student struct { ... }
//    Roster 结构体:
//      逻辑: 这是整个系统的“大脑”和“管理者”。它不仅仅是一个学生列表，还包含了管理这个列表所必需的“元数据”（metadata）。
//      ClassID int: 标识这个花名册属于哪个班级，是业务属性。
//      nextID int: 这是内部管理状态，负责生成唯一的学生ID。它是私有的（首字母小写），意味着只有 Roster 自己的方法才能访问和修改它，外部代码无法干预，保证了ID生成的安全性。
//      Students map[int]*Student: 这是核心数据存储。我们没有用切片 []Student，而是用了映射（map）。
//        key 是 int: 直接使用学生的 ID 作为键。
//        value 是 *Student: 存储的是指向 Student 对象的指针。这样做的好处是，当我们修改 map 中的学生对象时，我们是在修改原始的那个对象，而不是它的副本。这在Go中是处理复杂结构体集合时的常见做法，既高效又符合直觉。
// 2. 构造函数 (NewRoster)
//     逻辑: 任何复杂的对象都应该有一个“标准”的创建流程，以确保它被创建出来时处于一个可用、有效的状态。这就是构造函数的作用。
//     代码: func NewRoster(classID int) *Roster { ... }
//     执行细节:
//         接收一个 classID 作为参数。
//         创建一个 Roster 实例。
//         将 ClassID 赋值。
//         关键: 初始化 nextID 为 1，确保ID从一个有意义的数字开始。
//         关键: 使用 make(map[int]*Student) 来初始化 Students map。如果忘记这一步，map会是 nil，后续任何向map中添加数据的操作都会导致程序 panic（运行时恐慌）。
//         返回这个初始化好的 Roster 对象的指针。
// 3. 核心方法 (CRUD: Create, Read, Update, Delete)
//     AddStudent (Create):
//      逻辑: 负责“制造”一个新学生并将其登记在册。
//      执行细节:
// 	   接收创建学生所需的所有原始信息（name, age, subjects）。
//        使用 Roster 内部维护的 nextID 作为新学生的唯一ID。
//        创建一个 Student 实例。
//        将新创建的 Student 实例的指针存入 Students map中，键就是它的ID。
//        关键: 将 r.nextID++，为下一个要添加的学生做好准备。
//     ShowAllStudents (Read):
//      逻辑: 遍历并展示所有已登记的学生信息。
//      执行细节:
//      这是一个只读操作，所以它使用了值接收者 (r Roster)，虽然用指针接收者也没错，但值接收者更能从语法上表明“我不会修改你”。
//       通过 range 遍历 r.Students map，打印每个学生的信息。
//      UpdateStudent (Update):
//       逻辑: 用新的信息替换掉一个已存在的学生信息。
//       执行细节:
//         接收一个完整的 Student 对象作为参数，这个对象的 ID 字段指明了要更新哪个学生。
//         安全检查: 首先通过 ID 检查 Students map中是否存在这个学生。这是非常重要的防御性编程，防止对不存在的数据进行操作。如果不存在，返回一个 error。
// 		如果存在，直接用新的 Student 对象指针替换掉 map 中旧的对象指针。操作简单、高效。
//      DeleteStudent (Delete):
//       逻辑: 将一个学生从名册中除名。
//       执行细节:
//         接收一个 id 作为参数。
//         安全检查: 同样，先检查这个 id 对应的学生是否存在，不存在则返回错误。
//         如果存在，使用Go的内置函数 delete(r.Students, id) 从 map 中安全地移除这个键值对。

// ---1.数据结构定义---
type Student struct {
	ID       int
	Name     string
	Age      int
	Subjects []string
}

// Roster(花名册)，整个班级的学生,包含了学生列表和一个用于安全生成ID的计数器
type Roster struct {
	ClassId  int
	nextID   int              // 小写私有字段，用于内部自增ID，更安全
	Students map[int]*Student // 使用map来存储学生，相比切片，通过ID查找会非常快 每个int对应的是地址变量，可以修改原始变量
}

// ---2.构造函数---
// Roster构造函数，初始化一个空的班级花名册
func NewRoster(classID int) *Roster {
	return &Roster{
		ClassId:  classID,
		nextID:   1,                      //学生ID从1开始
		Students: make(map[int]*Student), //初始化map
	}
}

// --- 3.Roster的方法（增删查改）---
// AddStudent向花名册中添加一个新学生，使用指针接收者，因为要修改Roster的内部状态(nextID和Students.map)
func (r *Roster) AddStudent(name string, age int, subjects []string) *Student {
	// 创建一个新学生实例
	newStudent := &Student{
		ID:       r.nextID,
		Name:     name,
		Age:      age,
		Subjects: subjects, //在这个简单场景下，直接赋值即可
	}

	//将新学生存入Map
	r.Students[newStudent.ID] = newStudent
	r.nextID++ //ID自增，为下一个学生做准备

	fmt.Printf("成功添加学生: %s(ID:%d)\n", name, newStudent.ID)
	return newStudent
}

// ShowAllStudents 展示所有学生信息 使用值接收者Roster即可，因为它不需要修改任何状态，只读取
func (r Roster) ShowAllStudents() {
	fmt.Printf("\n---班级%d 学生列表 ---\n", r.ClassId)
	if len(r.Students) == 0 {
		fmt.Println("班级里没有学生")
	}
	for _, s := range r.Students {
		fmt.Printf("ID: %d, 姓名：%s, 年龄: %d, 科目: %v\n,", s.ID, s.Name, s.Age, s.Subjects)
	}
	fmt.Println("--------")
}

// UpdateStudent 更新一个已经存在的学生信息
// 我们传递一个包含更新后数据的Student对象，更灵活且类型安全,返回值是error类型（类似fastapi里面的jsonexception），成功的话就返回无错误响应
func (r *Roster) UpdateStudent(updatestudent Student) error {
	// 检查学生是否存在u
	_, exists := r.Students[updatestudent.ID]
	if !exists {
		return errors.New("更新失败，找不到该ID的学生")
	}
	// 用更新后的学生数据直接替换map中的旧数据，因为Map中的value类型是指针，这里需要取地址
	r.Students[updatestudent.ID] = &updatestudent
	fmt.Printf("成功更新学生信息 (ID: %d)\n", updatestudent.ID)
	return nil

}

// deleteStudent，从花名册中删除一个学生
func (r *Roster) DeleteStudent(id int) error {
	_, exists := r.Students[id]
	if !exists {
		return errors.New("删除失败：找不到该ID的学生")
	}
	delete(r.Students, id) // 内置函数，传入字典和要删除的key值
	fmt.Printf("成功删除学生 (ID: %d)\n", id)
	return nil
}

// 一个需要深拷贝的例子：为所有学生批量添加一个默认科目，但又不希望改变传入的那个“默认科目列表”
func (r *Roster) AddDefaultSubjectsToAll(defaultSubjects []string) {
	for _, student := range r.Students {
		// 如果这里直接 student.Subjects = append(student.Subjects, defaultSubjects...)
		// 那么所有学生都会共享同一个defaultSubjects切片的底层数组，非常危险
		// 正确的做法是为每个学生都创建一个副本
		newSubjects := make([]string, len(student.Subjects), len(student.Subjects)+len(defaultSubjects)) //创建了一个全新的切片，元素数量为原来科目的数量，容量为原来科目+新默认科目，拥有自己独立的一块新内存
		copy(newSubjects, student.Subjects)                                                              // 将原学生科目中的内容，逐个元素复制到这个全新的newSubjects切片中
		newSubjects = append(newSubjects, defaultSubjects...)                                            // 将默认科目添加到这个全新的切片中
		student.Subjects = newSubjects                                                                   // 学生的科目指向了这个全新的，独立的切片
	}
}

// 练习九：使用接口及依赖注入的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库

// 1.契约层(Abstraction): Logger 接口
//
//	LOgger定义了所有日志记录器都必须遵守的规范
type Logger interface {
	Log(message string)
	Error(message string)
}

// 2.实现层(Implementations):具体的日志记录器
// ---控制台日志实现--
type ConsoleLogger struct{} // ConsoleLogger将日志打印到标准输出(控制台)

func NewConsoleLogger() ConsoleLogger {
	// ConsoleLogger的构造函数
	return ConsoleLogger{}
}

func (c ConsoleLogger) Log(message string) {
	log.Printf("CONSOLE LOG:%s\n", message)
}

func (c ConsoleLogger) Error(message string) {
	log.Printf("CONSOLE ERROR:%s\n", message)
}

// ---文件日志实现---
type FileLogger struct {
	file *os.File // file变量存储的是一个指向os.File的指针 文件句柄
}

func NewFileLogger(filename string) (*FileLogger, error) {
	// FileLogger的构造函数，会打开或创建指定的日志文件，如果失败则返回错误
	// os.O_CREATE: 文件不存在则创建 os.O_WRONLY: 只写模式 os.O_APPEND: 追加内容到文件末尾
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{file: file}, err
	// 任何可能失败的初始化操作（如打开文件，连接数据库）都应当在构造函数中处理，并通过返回一个error来明确告知调用者操作是否成功

}

// 将日志写入文件的函数，使用指针接收者，因为它需要操作结构体内部的file字段
func (f *FileLogger) Log(message string) {
	// 使用fmt.Fprintf将格式化后的字符串写入到f.file中
	fmt.Fprintf(f.file, "FILE LOG:%s\n", message)
}

func (f *FileLogger) Error(message string) {
	fmt.Fprintf(f.file, "FILE Error:%s\n", message)
}

func (f *FileLogger) close() {
	f.file.Close() // Close是一个重要的办法，用于在程序结束时关闭文件句柄，防止资源泄漏
}

// 3.消费层 依赖于接口的服务
// UserService 负责用户相关的业务逻辑，它依赖于Logger接口 (业务服务，依赖抽象接口)
type UserService struct {
	logger Logger
}

func NewUserService(logger Logger) *UserService {
	// 业务服务的构造函数，通过依赖注入来接受一个Logger
	return &UserService{logger: logger}
}

func (us *UserService) CreateUser(username string) {
	//
	us.logger.Log(fmt.Sprintf("starting to create user'%s'...", username))
	// 这里是创建用户的复杂逻辑
	fmt.Printf("... (业务逻辑) User '%s' created in database.\n", username)
	us.logger.Log(fmt.Sprintf("User '%s' created successfully.", username))
}

func (us *UserService) DeleteUser(username string) {
	us.logger.Error(fmt.Sprintf("Starting to delete user '%s'...", username))
	// ... 假设这里是危险的删除用户逻辑 ...
	fmt.Printf("... (业务逻辑) User '%s' deleted from database.\n", username)
	us.logger.Error(fmt.Sprintf("User '%s' deleted successfully.", username))

}

// 练习十：生成一百个随机数
// 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// 开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 主 goroutine 从resultChan取出结果并打印到终端输出
// select多路复用条件：
// 必须满足以下两个条件之一：全部完成：所有 100 个任务的结果都被成功接收。
// 全局超时：如果在 1秒钟 内没有完成所有任务，主程序将不再等待，立即报告超时并退出
// 目标: 在规定时间内（1秒），并发地处理完一个固定数量（100个）的任务。如果超时，则放弃任务并报告
// 设计思想:
// 职责分离: 程序被清晰地划分为三个
// 生产者 (producer): 负责创建任务
// 消费者/工人 (worker): 负责执行任务。
// 调度者 (main): 负责启动和编排所有部分，并根据不同事件（任务完成或超时）来控制程序的最终流程。
// 通信取代共享内存: 角色之间不共享任何需要加锁的变量。它们通过 channel 这一管道来安全地传递数据（任务和结果）。

// producer函数-任务的创造者
// 接受一个context优雅退出
func producer(ctx context.Context, jobs chan<- int64, numjobs int) error {
	// jobchan 用于发送随机数任务 chan<- int64 表示这是一个只写channel，<-在chan右边表示这个函数只能向该channel发送数据，不能从中接收
	// numjobs表示需要生成的任务总数

	defer close(jobs)
	// 确保在producer函数执行完毕退出前，一定关闭jobs channel
	// 如果没有这句，worker goroutine中的for range循环在处理完所有任务后会永远阻塞，等待新任务，导致整个程序死锁

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 初始化一个高质量的随机数生成器，种子为当前时间的纳秒数
	for i := 0; i < numjobs; i++ {
		// 在发送任务前，检查context是否已被取消
		select {
		case <-ctx.Done():
			// 如果context被取消（例如因为超时或另一个goroutine出错),生产者就没必要再继续发送任务了
			fmt.Printf("生产者:收到取消信号，停止生产。 错误: %v\n", ctx.Err())
			return ctx.Err()
		default:
			// context正常，继续发送任务
			jobs <- r.Int63()
		}

	}
	// 循环生成指定数量的随机数，并将其发送到jobs channel

	fmt.Println("生产者：所有任务已经发送完毕")
	return nil

}

// worker函数-任务的执行者
// 接受context并返回error，不需要waitgroup参数了，现在使用context和errormap
func worker(ctx context.Context, id int, jobs <-chan int64, results chan<- string) error {
	// wg必须是指针
	// jobs <-chan int64: 一个只读channel, <-在chan左边，表示这个函数只能从该channel接收数据
	// results chan<- string:一个只写channel， 用于发送处理结果

	// 循环地从jobs channel中接受任务，直到channel被关闭
	// for...range用于、channel时，会自动处理以下逻辑:
	// 1.阻塞并等待jobs channel中有新数据 2.当有数据时，将其赋值给num并执行循环体 3.当jobs channel被关闭并且channel中所有已被发送的数据都被接收完毕后，循环自动结束
	for num := range jobs {
		//在处理每个任务前，先检查context是否已取消
		select {
		case <-ctx.Done():
			fmt.Printf("工人 %d：收到取消信号，停止工作，错误:%v\n", id, ctx.Err())
			return ctx.Err()
		default:
			// 模拟一个可能出错的场景
			if num%11 == 0 {
				// 假设遇到11的倍数就是一个无法处理的错误
				err := fmt.Errorf("工人 %d: 遇到一个无法处理的数字: %d", id, num)
				fmt.Println(err.Error())
				return err // 返回错误，这将触发整个errgroup的取消
			}

			time.Sleep(50 * time.Millisecond)
			// 每个任务耗时50ms

			originalNum := num
			var sum int64 = 0
			for num > 0 {
				sum += num % 10
				num /= 10
			}
			// 计算每位数字和

			result := fmt.Sprintf("Worker %d | 随机数: %d | 个位数之和: %d", id, originalNum, sum)
			// 打印至控制台

			// 在发送结果前，再次检查context
			select {
			case results <- result:
			case <-ctx.Done():
				fmt.Printf("工人 %d:准备发送结果时受到信号，错误:%v\n", id, ctx.Err())
				return ctx.Err()
			}
		}

	}
	return nil
	// 	当工人i返回错误后:
	// 1.errgroup 检测到错误
	// g.Go() 启动的某个函数返回了非 nil 错误

	// 2.errgroup 自动取消 gCtx
	// errgroup 内部会调用与 gCtx 关联的 cancel 函数
	// gCtx.Done() channel 被关闭

	// 3.所有其他 goroutine 收到取消信号
	// 其他正在工作的 worker 会在下次检查时退出: worker()中的case <- ctx.Done(): return ctx.Err()

	// 4.等待goroutine检测到gCtx完成
	// g.Wait(), close(resultChan)

	// 5.主程序的select命中取消分支
	// case <- gCtx.Done():
}

// 练习十一：迭代器的使用
// 假设有一个电商系统订单列表，需要处理所有“待支付”且金额>100元的订单
// 1.数据结构定义
type Order struct {
	ID     int
	Amount float64
	Status string // "Pending", "Paid", "Shipped"
}

// OrderManager 管理一组订单
type OrderManager struct {
	orders []Order
}

// 2.迭代器逻辑
// BigPendingOrders是一个专门的迭代器方法
// 它的任务是：遍历所有订单，剔除不符合条件的，只把符合条件的推给用户
func (om *OrderManager) BigPendingOrders(minAmount float64) iter.Seq[Order] {
	// 返回一个匿名函数，这就是标准的迭代器定义
	return func(yield func(Order) bool) {
		// 内部循环：负责具体逻辑（遍历切片，判断逻辑，筛选订单）
		for _, o := range om.orders {
			// 逻辑A ：如果不是待支付，直接跳过
			if o.Status != "Pending" {
				continue
			}
			// 逻辑B：如果金额不够大，直接跳过
			if o.Amount <= minAmount {
				continue
			}

			// 逻辑C（核心）：
			// 找到符合条件的订单，调用yield把它推给主函数
			// 这里的 !yield(o)是在问主函数：你还要继续吗？
			// 如果主函数里break了， yield会返回false，我们也必须return停止干活
			if !yield(o) {
				return
			}

		}
	}
}

// 练习十二:日志文件处理器 -flag,os,io,buffio等库的综合运用
// 让我们来构建一个稍微复杂但非常实用的例子：一个命令行工具，它可以完成以下任务：
// 读取一个指定的源日志文件。
// 过滤出包含特定关键字（如 "ERROR"）的行。
// 将这些过滤后的行写入一个新的目标文件。
// 在写入前，自动创建目标文件所在的目录（如果不存在）。
// 使用 flag 包来接收命令行参数。

// processLogFile 是核心处理函数
// 它逐行读取inputFIle, 检查是否包含keyword， 如果包含， 则写入outputFile
func processLogFile(inputFIle, outputFile, keyword string) error {
	// ---步骤1：打开源文件进行读取 ---
	// os.Open 只用于读取文件
	srcFile, err := os.Open(inputFIle)
	if err != nil {
		return fmt.Errorf("无法打开源文件%s:%w", inputFIle, err)
	}

	// 使用defer确保文件句柄在推出时一定会被关闭
	defer srcFile.Close()

	// --- 步骤2:确保目标路径存在 ---
	// filepath.Dir 获取文件路径中的目录部分 os.MkdirAll创建所有必需的父目录，如果存在也不会报错
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("无法创建目录%s:%w", outputDir, err)
	}

	// --- 步骤3：创建或清空目标文件进行写入 ---
	// os.Create是一个方便的函数，相当于os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// 它会创建文件，如果文件已存在，则会清空其内容
	dstFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("无法创建目标文件%s:%w", outputFile, err)
	}
	defer dstFile.Close()

	// --- 步骤4：逐行读取源文件并写入目标文件 ---
	// 使用bufio.Scanner来高效逐行读取大文件
	scanner := bufio.NewScanner(srcFile)
	// 使用bufio.Writer来合并多次小写入为一次大写入
	writer := bufio.NewWriter(dstFile)

	// defer writer.Flush()确保所有缓存都被写入,很重要，否则可能丢失最后一部分数据
	defer writer.Flush()

	// ---步骤5：逐行处理文件---
	fmt.Println("开始处理文件...")
	linesWritten := 0
	for scanner.Scan() {
		line := scanner.Text()
		// strings.Contains检查行中是否包含关键字
		if strings.Contains(line, keyword) {
			// writer.WriteString:写入字符串到缓冲区
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("写入目标文件时出错:%w", err)
			}
			linesWritten++
		}
	}

	// scanner.Scan()循环结束后，必须检查scanner.Err()以确定是正常结束还是因为读取错误
	if err := scanner.Err(); err != nil {
		// io.EOF(文件结束)不是一个错误，scanner会自动处理
		// 这里捕获的是真正的IO错误
		return fmt.Errorf("读取源文件时出错:%w", err)
	}
	fmt.Printf("成功处理文件，写入%d行包含%s的日志\n", linesWritten, keyword)
	return nil
}
