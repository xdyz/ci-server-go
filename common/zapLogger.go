package common

import (
	"go-basic-web/global"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志切割定义 传入日志文件名 所在位置
func getWriter(fileName string) *rotatelogs.RotateLogs {
	log, err := rotatelogs.New(
		// 分割后的文件名称 aaa.log.YYYY-MM-DD.HH
		fileName+"/%Y-%m-%d/"+"info.log",
		// rotatelogs.WithLinkName(fileName), // 生成软链，指向最新日志文件
		// 文件最大保存时间
		rotatelogs.WithMaxAge(time.Hour*24*30), // 文件最大保存时间 30天
		// 文件切割时间间隔
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔 24小时

		// 最大保存文件的个数，这里为10个超过10个 自动删除第一个
		// rotatelogs.WithRotationCount(10), // 日志切割时间间隔 24小时
	)

	if err != nil {
		panic(err)
	}

	return log
}

func initLogger() *zap.Logger {

	// 配置日志的格式  字段
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",                       // 结构化（json）日志的字段名 msg
		LevelKey:     "level",                     // 结构化（json）日志的字段名 level
		TimeKey:      "time",                      // 结构化（json）日志的字段名 time
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 将日志级别转换为大写 INFO WARN ERROR
		EncodeCaller: zapcore.ShortCallerEncoder,  // 将日志调用者转换为短文件名和行号 (common/zapLogger.go:23)
		// 时间的输出格式 类似于YYYY-MM-DD HH:MM:SS
		// 2006-01-02 15:04:05 就是GO的 YYYY-MM-DD HH:MM:SS 要求必须这么写
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, // 将时间转换为秒
	}

	// 当前日志级别小于 错误级别时 都为info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel // 只记录错误级别以下的日志 归为info类别
	})

	// 当前日志级别大于等于错误级别时 都为error级别
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel // 只记录错误级别以上的日志 归为error类别
	})

	// 实现切割日志文件 info级别和error级别都切割
	workDir, _ := os.Getwd() // 获取当前文件所在目录

	infoPath := path.Join(workDir, "logs/info") // 拼接info日志文件路径
	infoRotateLogs := getWriter(infoPath)

	errorPath := path.Join(workDir, "logs/error") // 拼接error日志文件路径
	errorRotateLogs := getWriter(errorPath)

	// 组合上述配置， NewTee 同时 表示将日志写入两个文件
	// 这里将两个日志都用json格式写入 NewJSONEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(infoRotateLogs), infoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(errorRotateLogs), errorLevel),
		// 输出到控制台
		// zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.DebugLevel)),

		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zap.NewAtomicLevelAt(zap.DebugLevel)),
	)

	// 生成最终的日志调用方法
	// addCaller 是否记录文件名和行号
	// addStacktrace 是否记录堆栈信息 传入等级 只有大于等于等级的日志才会记录堆栈信息
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger
}

func init() {

	global.Logger = initLogger() // 初始化日志 到全局变量
}
