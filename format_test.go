package logging

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/go-stack/stack"
	"github.com/stretchr/testify/require"
)

func TestRegexp(t *testing.T) {
	var formatRegexp = regexp.MustCompile(`%{(color|level|time|module|location|message)(?::(.*?))?}`)
	spec := "%{color}%{level:in}"

	matches := formatRegexp.FindAllStringSubmatchIndex(spec, -1)
	for _, match := range matches {
		t.Log(match)
	}
}

func TestParseFormat(t *testing.T) {
	mf := NewMultiFormatter()

	var tests = []struct {
		spec string
		e    Entry
	}{
		{
			spec: "%{color}%{level}[%{time}]%{color:reset}%{message}",
			e: Entry{
				Level:   DebugLevel,
				Time:    time.Now(),
				Module:  "consensus",
				Call:    stack.Caller(0),
				Message: "prepare consensus",
			},
		},
		{
			spec: "%{color}[%{time}]%{color:reset} => %{message}",
			e: Entry{
				Level:   InfoLevel,
				Time:    time.Now(),
				Module:  "consensus",
				Call:    stack.Caller(0),
				Message: "prepare consensus",
			},
		},
		{
			spec: "%{color}%{level}[%{time}][%{module}]%{color:reset}%{message}",
			e: Entry{
				Level:   WarnLevel,
				Time:    time.Now(),
				Module:  "consensus",
				Call:    stack.Caller(0),
				Message: "prepare consensus",
			},
		},
		{
			spec: "%{color}%{level}[%{time}][%{location}]%{color:reset}%{message}",
			e: Entry{
				Level:   ErrorLevel,
				Time:    time.Now(),
				Module:  "consensus",
				Call:    stack.Caller(0),
				Message: "prepare consensus",
			},
		},
		{
			spec: "%{color}[%{time}][%{location}]%{message}%{color:reset}",
			e: Entry{
				Level:   PanicLevel,
				Time:    time.Now(),
				Module:  "consensus",
				Call:    stack.Caller(0),
				Message: "prepare consensus",
			},
		},
	}

	buf := new(bytes.Buffer)
	for _, test := range tests {
		formatters, err := ParseFormat(test.spec)
		require.NoError(t, err)
		mf.SetFormatters(formatters)
		mf.Format(buf, test.e)
		buf.Write([]byte("\n"))
	}
	fmt.Println(buf.String())
}
