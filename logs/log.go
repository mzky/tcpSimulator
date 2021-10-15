package logs

import (
	"github.com/bingoohuang/golog/pkg/logfmt"
	"github.com/bingoohuang/golog/pkg/spec"
	"github.com/bingoohuang/golog/pkg/timex"
)

func InitLog(logPath string) {
	var size spec.Size
	size.Parse("100M")
	layout := `%t{yyyy-MM-dd HH:mm:ss.SSS} [%-5l{length=5}] %caller %fields %msg%n`
	maxAge, _ := timex.ParseDuration("1095d") // 最大保留3年
	gzipAge, _ := timex.ParseDuration("3d")   // 3天后归档压缩

	logrusOption := logfmt.LogrusOption{
		Level:       "debug",
		LogPath:     logPath,
		Rotate:      ".yyyy-MM-dd",
		MaxAge:      maxAge,
		GzipAge:     gzipAge,
		MaxSize:     int64(size),
		PrintColor:  true,
		PrintCaller: true,
		Stdout:      true,
		Layout:      layout,
	}
	logrusOption.Setup(nil)
}
