package hook

import (
	"io"

	"github.com/sirupsen/logrus"
)

// logrus output support multiple io.Writer by using io.MultiWriter,
// but it can't control the output of each io.Writer, especially when
// you need to print colorful logs to the terminal but not to the file.
// So, we use a io.Writer hook to solve this problem.

// writerHook is a hook to write logs to the given io.Writer.
type writerHook struct {
	output io.Writer
	levels []logrus.Level
}

// NewWriterHook returns a new writerHook,
// which will write logs to the given io.Writer.
func NewWriterHook(output io.Writer) *writerHook {
	return &writerHook{
		output: output,
		levels: []logrus.Level{
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	}
}

// SetLevels sets the levels which this hook will trigger.
func (w *writerHook) SetLevels(levels []logrus.Level) {
	w.levels = levels
}

// Levels implements logrus.Hook interface.
func (w *writerHook) Levels() []logrus.Level {
	return w.levels
}

// Fire implements logrus.Hook interface.
func (w *writerHook) Fire(entry *logrus.Entry) error {
	_, err := w.output.Write(w.format(entry))
	return err
}

func (w *writerHook) format(entry *logrus.Entry) []byte {
	format := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	serialized, _ := format.Format(entry)
	return serialized
}
