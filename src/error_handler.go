package error_handler

import "log"

func ErrorHandler(errorString string, err error) {
	if err == nil {
		return
	}

	log.Fatal(errorString)
	panic(err)
}
