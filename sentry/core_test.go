package zapsentry

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestSentryCore_CheckZapCoreInterface(t *testing.T) {
	var _ zapcore.Core = &Core{}
}
