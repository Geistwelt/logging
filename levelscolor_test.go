package logging

import (
	"bytes"
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	var tests = []struct {
		desc    string
		level   LogLevel
		message string
	}{
		{
			desc:    "debug",
			level:   DebugLevel,
			message: "output debug level log",
		},
		{
			desc:    "info",
			level:   InfoLevel,
			message: "output info level log",
		},
		{
			desc:    "warn",
			level:   WarnLevel,
			message: "output warn level log",
		},
		{
			desc:    "error",
			level:   ErrorLevel,
			message: "output error level log",
		},
		{
			desc:    "panic",
			level:   PanicLevel,
			message: "output panic level log",
		},
	}

	buf := new(bytes.Buffer)
	buf.Reset()
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fmt.Fprint(buf, test.level.SpecifiedColor().Color(), test.message, ResetColor()+"\n")
			t.Log("")
		})
	}
	fmt.Print(buf.String())
}
