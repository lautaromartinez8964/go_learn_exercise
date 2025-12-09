好的，我们继续深入 Go 的标准库，这次是 `log` 包。它是一个看似简单但至关重要的包，用于在程序中生成日志消息。

### 一、`log` 包的核心思想与目的

**核心目的**：`log` 包提供了一种简单、标准化的方式来**输出格式化的日志消息**。

日志是程序运行的“航海日志”，它记录了关键的事件、错误和状态信息，对于调试、监控和问题排查至关重要。`log` 包的设计哲学是**简单、直接、够用**。

**设计思想**：
*   **开箱即用**: 你无需任何配置，只需 `import "log"`，就可以立即使用 `log.Println("hello")` 来打印日志。
*   **默认安全**: 默认情况下，所有日志都**原子地**写入**标准错误输出 (Stderr)**。这意味着即使在多个 goroutine 中同时调用日志函数，日志消息也不会出现交错或损坏。写入 Stderr 是一个Unix的惯例，它允许用户将程序的正常输出（Stdout）和错误/诊断信息（Stderr）分开重定向。
*   **可配置性**: 虽然默认设置很方便，但你可以轻松地配置日志的**输出位置**（如文件）、每条日志的**前缀**以及包含的**元数据**（如日期、时间、文件名）。
*   **提供了基础的日志级别**: 它内置了三种行为级别：常规打印（Print）、打印后退出（Fatal）和打印后恐慌（Panic）。

---

### 二、核心组件：标准日志记录器 (Standard Logger)

`log` 包的核心是一个名为**“标准日志记录器”**的全局单例。我们平时直接调用的包级别函数，如 `log.Printf`, `log.Fatal`, `log.Panicf` 等，实际上都是在使用这个全局的、共享的记录器实例。

这使得在程序的任何地方记录日志都非常方便，但也意味着对它的任何配置（如更改输出文件）都会影响到整个程序。

---

### 三. 基本用法：三种行为级别的函数

`log` 包的函数可以分为三大家族，它们的命名和行为都非常有规律。

#### 1. Print 家族：只打印日志

这些函数只是简单地输出日志消息，然后程序继续正常执行。

*   `log.Print(v ...interface{})`:  以默认格式打印参数，参数间有空格。
*   `log.Println(v ...interface{})`: 类似于 `Print`，但在参数间加空格，并在末尾添加换行符。
*   `log.Printf(format string, v ...interface{})`:  使用格式化字符串打印，类似于 `fmt.Printf`。

**示例**：
```go
log.Println("服务启动成功。")
user := "admin"
port := 8080
log.Printf("用户 '%s' 正在监听端口 %d", user, port)```

#### 2. Fatal 家族：打印日志后程序退出

这些函数在打印完日志消息后，会**立即调用 `os.Exit(1)` 来终止整个程序**。这通常用于发生了无法恢复的严重错误，程序没有继续运行的意义了。

*   `log.Fatal(v ...interface{})`
*   `log.Fataln(v ...interface{})`
*   `log.Fatalf(format string, v ...interface{})`
```go

**示例**:
```go
db, err := sql.Open("mysql", "user:password@/dbname")
if err != nil {
    // 数据库连接失败是致命错误，程序无法继续
    log.Fatalf("无法连接到数据库: %v", err)
}
// 这行代码永远不会被执行，因为上面已经退出了
fmt.Println("这不会被打印")
```

#### 3. Panic 家族：打印日志后引发恐慌

这些函数在打印完日志消息后，会**引发一个 `panic`**。`panic` 是一种 Go 语言的异常处理机制。如果这个 `panic` 没有被 `recover` 捕获，程序的默认行为是打印堆栈跟踪信息然后退出。

*   `log.Panic(v ...interface{})`
*   `log.Panicln(v ...interface{})`
*   `log.Panicf(format string, v ...interface{})`

**示例**:
```go
// 假设有一个不应该发生的内部逻辑错误
if calculatedValue < 0 {
    log.Panicf("计算结果出现不一致！值为: %d，这不应该发生。", calculatedValue)
}
```

---

### 四、配置标准日志记录器

你可以通过 `log` 包提供的 `Set` 系列函数来修改标准日志记录器的行为。

#### 1. `log.SetOutput(w io.Writer)` - 配置输出位置

默认输出到 `os.Stderr`。你可以将其改为任何实现了 `io.Writer` 接口的东西，最常见的就是文件。

**示例：将日志写入文件**
```go
// os.O_CREATE: 如果文件不存在则创建
// os.O_WRONLY: 只写模式
// os.O_APPEND: 追加模式
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    log.Fatal("无法打开日志文件: ", err)
}
defer file.Close()

// 将标准日志记录器的输出设置为文件
log.SetOutput(file)

log.Println("这条日志将会被写入 app.log 文件。")
```

#### 2. `log.SetPrefix(prefix string)` - 配置日志前缀

为每一条日志消息添加一个固定的前缀字符串。

**示例**:
```go
log.SetPrefix("[WebApp] ")
log.Println("用户请求处理中...")
// 输出: [WebApp] 2023/10/27 10:30:00 用户请求处理中...
```

#### 3. `log.SetFlags(flag int)` - 配置元数据

这是最有用的配置之一。你可以控制每条日志自动包含哪些元数据，比如日期、时间、文件名和行号。这些标志是整数常量，可以通过 `|` (位或运算符) 组合使用。

**常用标志**:
*   `log.Ldate`: 日期，格式 `2009/01/23`
*   `log.Ltime`: 时间，格式 `01:23:23`
*   `log.Lmicroseconds`: 微秒，格式 `01:23:23.123123`
*   `log.Llongfile`: 完整文件路径和行号, e.g., `/a/b/c/main.go:23`
*   `log.Lshortfile`: 文件名和行号, e.g., `main.go:23`
*   `log.LUTC`: 使用 UTC 时间而非本地时间
*   `log.LstdFlags`: 标准标志，是 `log.Ldate | log.Ltime` 的组合，**这是默认值**。

**示例**:
```go
// 设置日志标志为：日期 + 毫秒级时间 + 短文件名和行号
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
log.Println("这是一个带有详细元数据的日志。")
// 输出: 2023/10/27 10:30:00.123456 main.go:42 这是一个带有详细元数据的日志。
```

---

### 五、更进一步：创建自定义的 `Logger` 实例

直接使用全局的标准日志记录器虽然方便，但在大型应用或库中，这通常不是一个好主意，因为它会产生“副作用”（一个模块的配置会影响另一个）。

更好的做法是创建自己的 `log.Logger` 实例。

**`log.New(out io.Writer, prefix string, flag int)`**: 这是创建自定义 Logger 的构造函数。它接收的三个参数正好对应我们上面配置标准 Logger 的三个方面。

**示例：创建多个独立的 Logger**
```go
var (
    InfoLogger    *log.Logger
    WarningLogger *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("无法打开日志文件: ", err)
    }

    // 创建 INFO 级别的 Logger，输出到标准输出
    InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

    // 创建 WARNING 级别的 Logger，也输出到标准输出
    WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)

    // 创建 ERROR 级别的 Logger，同时输出到标准错误和文件
    // io.MultiWriter 可以将写入操作复制到多个 Writer
    multiWriter := io.MultiWriter(os.Stderr, file)
    ErrorLogger = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
    InfoLogger.Println("服务正在启动...")
    WarningLogger.Println("配置文件未找到，使用默认配置。")
    ErrorLogger.Println("一个模拟的严重错误发生了。")
}
```
在这个例子中，我们通过创建自定义 Logger 实现了：
*   **伪日志级别**：通过不同的前缀 (`INFO:`, `WARNING:`) 模拟了日志级别。
*   **独立的配置**：每个 Logger 都有自己的格式 (`Lshortfile` vs `Llongfile`)。
*   **灵活的输出**: `ErrorLogger` 可以同时向屏幕和文件输出。

---

### 六、局限性与现代实践

标准的 `log` 包非常适合中小型项目、命令行工具或学习目的。但对于大型、高并发的生产级应用，它有几个明显的局限性：

1.  **没有真正的日志级别**: 我们只能通过前缀来“假装”有 INFO, DEBUG, ERROR 等级别，无法做到“只开启 DEBUG 以上级别的日志”这样的动态过滤。
2.  **非结构化日志**: 它输出的是纯文本字符串。现代日志系统（如 ELK Stack, Datadog）更喜欢**结构化日志**（如 JSON 格式），因为它们易于机器解析、索引和查询。
    *   **文本日志**: `[ERROR] User login failed for user 'admin'`
    *   **结构化日志**: `{"level": "error", "message": "User login failed", "user_id": "admin"}`
3.  **扩展性有限**: 很难添加钩子（hook，例如在每次打错误日志时都发一封邮件）或自定义日志格式。

**现代实践**:
对于需要更强大日志功能的应用，社区和官方提供了更好的选择：
*   **`slog` (Go 1.21+ 新增)**: Go 官方推出的**新的结构化日志包**，现在是标准库的一部分。它解决了上述所有痛点，支持日志级别、结构化输出（JSON），并且性能极高。**对于新项目，`slog` 是首选**。
*   **`logrus`**: 非常流行和成熟的第三方库，支持钩子、多种格式化器（JSON, text）和日志级别。
*   **`zap`**: 由 Uber 开发，以其**极致的性能**而闻名，同样是结构化的日志库。

**结论**：`log` 包是每个 Go 开发者都应该熟练掌握的基础工具。它简单可靠，足以应付许多场景。但当你开始需要日志级别过滤、结构化查询和更高性能时，就应该考虑转向 `slog` 或其他优秀的第三方库。