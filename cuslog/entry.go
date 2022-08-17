package cuslog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)
//首先定义了一个 Entry 结构体类型，该类型用来保存所有的日志信息，即日志配置和日志内容。
//写入逻辑都是围绕 Entry 类型的实例来完成的。用 Entry 的 write 方法来完成日志的写入，在 write 方法中，会首先判断日志的输出级别和开关级别，如果输出级别小于开关级别，则直接返回，不做任何记录。
//在 write 中，还会判断是否需要记录文件名和行号，如果需要则调用 runtime.Caller() 来获取文件名和行号，
//调用 runtime.Caller() 时，要注意传入正确的栈深度。write 函数中调用 e.format 来格式化日志，调用 e.writer 来写入日志，
//在创建 Logger 传入的日志配置中，指定了输出位置 output io.Writer ，output 类型为 io.Writer
type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map map[string]interface{}
	Level Level
	Time time.Time
	File string
	Line int
	Func string
	Format string
	Args []interface{}
}


func entry (logger *logger)*Entry{
	return &Entry{logger: logger,Buffer: new(bytes.Buffer),Map: make(map[string]interface{},5)}
}


func(e *Entry)write(level Level,format string,args ...interface{}){
	if e.logger.opt.level > level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}


func (e *Entry)format(){
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}