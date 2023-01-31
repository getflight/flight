package formatters

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TextFormatter struct {
	log.TextFormatter
}

func (f *TextFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.appendValue(b, entry.Message)

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *TextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}

	stringVal, ok := value.(string)

	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}
