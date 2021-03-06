package posposet

import (
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.StandardLogger()
	log.SetLevel(logrus.DebugLevel)
}

// SetLogger sets logger for whole package.
func SetLogger(custom *logrus.Logger) {
	if custom == nil {
		panic("Nil-logger set")
	}
	log = custom
}
