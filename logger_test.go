package logging

import (
	"os"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	opt := Option{
		Module:         "blockchain",
		FilterLevel:    DebugLevel,
		Spec:           "%{time} %{module} %{message}",
		FormatSelector: "",
		Writer:         os.Stdout,
	}

	logger, err := NewLogger(opt)
	require.NoError(t, err)

	logger.Debug("debug", "name", "wuxiangyu")
	logger.Info("info", "name", "wuxiangyu")
	logger.Warn("warn", "name", "wuxiangyu")
	logger.Error("error", "name", "wuxiangyu")
	// logger.Panic("panic", "name", "wuxiangyu")
}

func TestMain(t *testing.T) {
	t.Log("hello")
}

func TestDeriveChild(t *testing.T) {
	opt := Option{
		Module:         "blockchain",
		FilterLevel:    DebugLevel,
		Spec:           "%{time} %{module} %{message}",
		FormatSelector: "",
		Writer:         os.Stdout,
	}

	logger, err := NewLogger(opt)
	require.NoError(t, err)

	logger.SetModule("consensus", WarnLevel)

	child := logger.DeriveChildLogger("consensus")

	child.Debug("debug", "test", 12)
	child.Info("info", "test", 11)
	child.Warn("warn", "test", 10)
	child.Error("error", "test", 9)

	child.Update(Option{
		Module:         "",
		FilterLevel:    DebugLevel,
		Spec:           "",
		FormatSelector: "json",
		Writer:         nil,
	})

	child.Debug("debug", "test", 12)
	child.Info("info", "test", 11)
	child.Warn("warn", "test", 10)
	child.Error("error", "test", 9)

	child.Update(Option{
		Module:         "p2p",
		FilterLevel:    InfoLevel,
		Spec:           "%{color}%{level} => %{message}%{color:reset}",
		FormatSelector: "terminal",
		Writer:         nil,
	})
	child.Debug("debug", "test", 12)
	child.Info("info", "test", 11)
	child.Warn("warn", "test", 10)
	child.Error("error", "test", 9)
}

func TestLogIntoFile(t *testing.T) {
	file, _ := os.OpenFile("log.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	opt := Option{
		Module:         "blockchain",
		FilterLevel:    DebugLevel,
		Spec:           "%{time} %{module} => %{message}",
		FormatSelector: "json",
		Writer:         file,
	}

	logger, err := NewLogger(opt)
	require.NoError(t, err)

	logger.Debug("debug", "name", "wuxiangyu", "age", 18)
	logger.Info("info", "name", "wuxiangyu", "age", 18)
	logger.Warn("warn", "name", "wuxiangyu", "age", 18)
	logger.Error("error", "name", "wuxiangyu", "age", 18)

	file.Close()
}

func TestDefaultSpec(t *testing.T) {
	opt := Option{
		Module:         "blockchain",
		FilterLevel:    DebugLevel,
		Spec:           "%{color}%{level}[%{time}] [%{module}]%{color:reset}: %{message}",
		FormatSelector: "json",
		Writer:         os.Stdout,
	}
	logger, err := NewLogger(opt)
	if err != nil {
		panic(err)
	}

	logger.SetModule("consensus", InfoLevel)
	consensusLogger := logger.DeriveChildLogger("consensus")
	consensusLogger.Update(Option{
		Spec: "%{level}[%{time}] [%{module}]: %{message}",
	})
	consensusLogger.Info("info message", "key", "value")
}

func TestLogFormat(t *testing.T) {
	opt := Option{
		Module:         "blockchain",
		FilterLevel:    DebugLevel,
		Spec:           "%{color}%{level}[%{time}] [%{module}]%{color:reset}: %{message}",
		FormatSelector: "json",
		Writer:         os.Stdout,
	}
	logger, err := NewLogger(opt)
	require.NoError(t, err)

	logger.Debugf("在 [%s] 处创建 KeyStore...", "/root/exp/")
	logger.Infof("Creating KeyStore at [%s]...", "/root/exp/")
	logger.Warnf("Creating KeyStore at [%s]...", "/root/exp/")
	logger.Errorf("Creating KeyStore at [%s]...", "/root/exp/")
}