package oomg

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	env := os.Getenv("OOMG_ENV")

	if env == "development" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.FatalLevel)
	}

	log.SetReportCaller(true)

}

//Well I suposed to be internal use  but at least need some good logging to develop and
// may just maybe someone else want to have some verbose of the library
