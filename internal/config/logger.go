package config

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"math/rand"
	"os"
	"path"
	"runtime"
	"time"
)

type logger struct {
	Filename      string
	MaxSize       int
	MaxAge        int
	Compress      bool
	FullLevelName bool
	Color         bool
}

var (
	format *zt_formatter.ZtFormatter
	jack   *lumberjack.Logger
	sid    uint32
	Log    *log.Entry
)

func (l *logger) newLogger() {
	rand.Seed(time.Now().UnixNano())
	sid = rand.Uint32()

	format = &zt_formatter.ZtFormatter{
		Formatter: nested.Formatter{
			NoColors:      !l.Color,
			ShowFullLevel: l.FullLevelName,
		},
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	l.localInit()
}

func (l *logger) localInit() {
	jack = &lumberjack.Logger{
		Filename: l.Filename,
		MaxSize:  l.MaxSize,
		MaxAge:   l.MaxAge,
		Compress: l.Compress,
	}
	writer := io.MultiWriter(os.Stdout, jack)
	log.SetOutput(writer)
	log.SetReportCaller(true)
	log.SetFormatter(format)
	Log = log.WithField("sid", sid)
}
