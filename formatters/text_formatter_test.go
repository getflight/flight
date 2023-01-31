package formatters

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestTextFormatter(t *testing.T) {
	t.Run("Format returns log in with message", func(t *testing.T) {
		// given
		formatter := TextFormatter{}
		entry := &log.Entry{
			Logger:  nil,
			Data:    nil,
			Time:    time.Now(),
			Level:   0,
			Caller:  nil,
			Message: "test",
			Buffer:  nil,
			Context: nil,
		}

		// when
		result, err := formatter.Format(entry)

		// then
		assert.Nil(t, err)
		assert.True(t, strings.Contains(string(result), entry.Message))
	})
}
