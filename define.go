package logger

const layout = "2006-01-02 15:04:05.000"

const (
	//0:重置所有属性
	//1:设置加粗
	//2:置灰
	//3:斜体
	//4:下划线
	//30-37:前景颜色(字体颜色)
	//40-47:背景颜色
	//90-97:前景高亮颜色
	//100-107:背景高亮颜色
	traceColor = "\u001B[96m%s\u001B[0m"
	debugColor = "\u001B[95m%s\u001B[0m"
	infoColor  = "\u001B[92m%s\u001B[0m"
	warnColor  = "\u001B[93;1m%s\u001B[0m"
	errorColor = "\u001B[91;1m%s\u001B[0m"
	fatalColor = "\u001B[101;93m%s\u001B[0m"
)

const (
	traceFormat = "[%s TRAC] %v\n"
	debugFormat = "[%s DEBU] %v\n"
	infoFormat  = "[%s INFO] %v\n"
	warnFormat  = "[%s WARN] %v\n"
	errorFormat = "[%s ERRO] %v\n"
	fatalFormat = "[%s FATA] %v\n"
)

const (
	logExt = ".log"
)

const (
	KB = 1024 // 定义1KB为1024字节
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

var ColorMap = map[Level]string{
	LevelTrace: traceColor,
	LevelDebug: debugColor,
	LevelInfo:  infoColor,
	LevelWarn:  warnColor,
	LevelError: errorColor,
	LevelFatal: fatalColor,
}
