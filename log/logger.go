package log

import (
	"errors"
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var Log *Logger

type Logger struct {
	//entry       *log.Entry
	localLogger *log.Logger
	esLogger    *log.Logger
}

func (l *Logger) init() {
	l.localLogger = log.New()
	l.esLogger = log.New()
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.localLogger.WithFields(log.Fields{"trace_id": ""}).Infof(format, args...)
	l.esLogger.WithFields(log.Fields{"trace_id": ""}).Infof(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.localLogger.Errorf(format, args...)
	l.esLogger.Errorf(format, args...)
}

func (l *Logger) InfoWithID(traceID string, args ...interface{}) {
	l.localLogger.WithFields(log.Fields{"trace_id": traceID}).Info(args...)
	l.esLogger.WithFields(log.Fields{"trace_id": traceID}).Info(args...)
}

func (l *Logger) ErrorWithID(traceID string, args ...interface{}) {
	l.localLogger.WithFields(log.Fields{"trace_id": traceID}).Error(args...)
	l.esLogger.WithFields(log.Fields{"trace_id": traceID}).Error(args...)
}

func (l *Logger) Infof(traceID string, format string, args ...interface{}) {
	l.localLogger.WithFields(log.Fields{"trace_id": traceID}).Infof(format, args...)
	l.esLogger.WithFields(log.Fields{"trace_id": traceID}).Infof(format, args...)
}

func (l *Logger) Errorf(traceID string, format string, args ...interface{}) {
	l.localLogger.WithFields(log.Fields{"trace_id": traceID}).Errorf(format, args...)
	l.esLogger.WithFields(log.Fields{"trace_id": traceID}).Errorf(format, args...)
}


func InitLogger(logPath string) error {
	logger := &Logger{}
	logger.init()
	logger.localLogger.Hooks.Add(NewContextHook())
	logger.esLogger.Hooks.Add(NewContextHook())
	writer, err := rotateLogs.New(
		logPath+"local/root.%Y-%m-%d",
		rotateLogs.WithLinkName(logPath),
		rotateLogs.WithRotationCount(180) ,    //number 默认7份 大于7份 或到了清理时间 开始清理
		//rotateLogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err !=nil {
		return err
	}
	writerEs, err := rotateLogs.New(
		logPath+"elasticsearch/es.%Y-%m-%d",
		rotateLogs.WithLinkName(logPath),
		rotateLogs.WithRotationCount(10),        //number 默认7份 大于7份 或到了清理时间 开始清理
		//rotateLogs.WithRotationTime(time.Duration(24)*time.),
	)
	if err !=nil {
		return err
	}
	logger.localLogger.SetOutput(writer)
	logger.esLogger.SetOutput(writerEs)

	logger.localLogger.SetFormatter(new(MyFormatter))
	logger.esLogger.SetFormatter(&log.JSONFormatter{})
	Log = logger
	return nil
}

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s %s] [%s] %s %s\n", timestamp, entry.Data["location"], entry.Data["fn_name"],
		strings.ToUpper(entry.Level.String()), entry.Data["trace_id"], entry.Message)
	return []byte(msg), nil
}

// ContextHook for log the call context
// location 文件位置，fn_name 函数名称
type contextHook struct {
	Location, FnName, TraceID string
	Skip                      int
	levels                    []log.Level
}

// NewContextHook use to make an hook
// 根据上面的推断，我们递归深度可以设置到9
func NewContextHook(levels ...log.Level) log.Hook {
	hook := contextHook{
		Location: "location",
		FnName:   "fn_name",
		Skip:     10,
		levels:   levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = log.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook contextHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire implement fire
func (hook contextHook) Fire(entry *log.Entry) error {
	_, entry.Data[hook.Location] = findCaller(hook.Skip)
	entry.Data[hook.FnName], _ = findCaller(hook.Skip)
	return nil
}

/*
	对caller进行递归查询，直到找到非logrus包产生的第一个调用.
 	因为filename获取到了上层目录名，因此所有logrus包的调用的文件名都是 logrus/...
 	因此通过排除logrus开头的文件名，就可以排除所有logrus包的自己的函数调用
*/
func findCaller(skip int) (string, string) {
	var file, fnName string
	var line int
	for i := 0; i < 10; i++ {
		fnName, file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fnName, fmt.Sprintf("%s:%d", file, line)
}

/*
	获取函数名称 fnName := runtime.FuncForPC(pc).Name()
 	文件的全路径往往很长, 而文件名在多个包中往往有重复，因此这里选择多取一层, 取到文件所在的上层目录那层.
*/
func getCaller(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	var fnName string
	fnPath := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if len(fnPath) > 0 {
		fnName = fnPath[1]
	}
	if !ok {
		return "", "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return fnName, file, line
}

func InitLogDir(logPath string, appName string) (*logFileWriter, error) {
	fileDate := time.Now().Format("2006-01-02")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s", logPath), os.ModePerm)
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("%s/%s.log", logPath, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		return nil, err
	}
	fileWriter := &logFileWriter{file, logPath, fileDate, appName}
	return fileWriter, nil
}

type logFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}

	// 当天日志为输出为 root.log，并将昨天的备份
	if p.fileDate != time.Now().Format("2006-01-02") { // 需要将昨天的进行备份
		p.fileDate = time.Now().Format("2006-01-02")
		srcName := fmt.Sprintf("%s/%s.log", p.logPath, p.appName)
		dstName := fmt.Sprintf("%s/%s.log.%s", p.logPath, p.appName, time.Now().Add(
			-24*time.Hour).Format("2006-01-02"))
		err = CopyFile(srcName, dstName, 32)
		if err != nil {
			return 0, err
		}
		err := os.Truncate(srcName, 0)
		if err != nil {
			return 0, err
		}
		file, err := os.OpenFile(srcName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		if err != nil {
			return 0, err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return 0, err
		}
		p.file = file
	}
	n, e := p.file.Write(data)
	return n, e
}

func CopyFile(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
