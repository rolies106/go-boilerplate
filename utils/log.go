package utils

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gorilla/schema"
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
)

func GetLogger() loggo.Logger {

	var logger = loggo.GetLogger("mortred")

	// Configuring logger
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))

	return logger
}

func Log(etype string, err error) {

	switch r := reflect.ValueOf(err); r.Kind() {
	case reflect.String:
	case reflect.Ptr:
		logging(etype, err)
	case reflect.Map:
		for _, v := range err.(schema.MultiError) {
			logging(etype, v)
		}
	default:
		fmt.Printf("Unhandled error kind %s\n", r.Kind())
	}
}

func logging(etype string, err error) {

	switch etype {
	case "error":
		GetLogger().Errorf(err.Error() + "\n")
	case "info":
		GetLogger().Infof(err.Error() + "\n")
	}
}
