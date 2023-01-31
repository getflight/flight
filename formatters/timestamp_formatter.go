package formatters

import (
	"bytes"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type TimestampFormatter struct {
	log.TextFormatter
}

func (f *TimestampFormatter) Format(entry *log.Entry) ([]byte, error) {
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

func (f *TimestampFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}

	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	stringVal = fmt.Sprintf("%s %s", time.Now().Format("15:04:05"), stringVal)

	b.WriteString(stringVal)
}
