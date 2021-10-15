package common

import (
	"github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		logrus.Fatalf("Fatal error:%s", err.Error())
	}
}
