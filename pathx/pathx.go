package pathx

import (
	"fmt"
	"os"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeOutputDirectory(path string) (string, error) {
	// get the current data + time using Go's stupid formatting
	t := time.Now()
	dirname := t.Format("2006-01-01 15:04:05")
	final := path + "/" + dirname

	// check if path is valid and directory
	pathExists, err := PathExists(final)
	if err != nil {
		return final, fmt.Errorf("could not check for existing path: %v", err)
	}
	if pathExists {
		// make sure it is directory with 0o644 permissions
		return final, fmt.Errorf("found existing path: %v", final)
	}

	// create directory with data + time
	err = os.MkdirAll(final, 0766)
	if err != nil {
		return final, fmt.Errorf("could not make directory: %v", err)
	}
	return final, nil
}
