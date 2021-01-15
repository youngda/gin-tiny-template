package logger

import (
	"go.uber.org/zap"
	"testing"
)

func Test_PrintLog(t *testing.T) {
	Init("./logs")
	Log.Info("这是msg",zap.String("自定义名字一","自定义值一"))
}
