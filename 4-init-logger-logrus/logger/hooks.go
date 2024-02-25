package logger

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// hidePrivateInfoHook
// hook for hiding in logger private info such as password, secrets, etc
type hidePrivateInfoHook struct {
	privateWords []string
}

func (h *hidePrivateInfoHook) Levels() []logrus.Level {
	// define on which levels hook should set up
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (h *hidePrivateInfoHook) Fire(e *logrus.Entry) error {
	// example of hiding private info in message
	e.Message = hidePrivateInfo(e.Message, h.privateWords)
	// example of hiding private info in "URI" field of message
	if _, ok := e.Data["URI"]; ok {
		e.Data["URI"] = hidePrivateInfo(e.Data["URI"].(string), h.privateWords)
	}
	return nil
}

// hidePrivateInfo replaces remainder of message after private word with "*****"
func hidePrivateInfo(message string, privateWords []string) string {
	for _, word := range privateWords {
		r, _ := regexp.Compile(strings.ToLower(word) + `.*`)
		if strings.Contains(strings.ToLower(message), strings.ToLower(word)) {
			message = r.ReplaceAllString(strings.ToLower(message), strings.ToLower(word)+"*****")
		}
	}
	return message
}
