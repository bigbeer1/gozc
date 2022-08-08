package http

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateFileHttp(modelList map[string]*CodeTuple, srcPath string) error {

	dirAbs := filepath.Dir(srcPath)

	dirAbs = filepath.Join(dirAbs, "http")

	is, _ := IsPathExist(dirAbs)
	if is == false {
		err := os.Mkdir(dirAbs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for tableName, codes := range modelList {

		name := fmt.Sprintf("%v.json", SafeString(tableName))
		filename := filepath.Join(dirAbs, name)
		err := ioutil.WriteFile(filename, []byte(codes.Api), os.ModePerm)
		if err != nil {
			return err
		}

	}
	fmt.Println("HttpDone.")
	return nil
}

// SafeString converts the input string into a safe naming style in golang
func SafeString(in string) string {
	if len(in) == 0 {
		return in
	}
	if strings.Contains(in, "_") {
		in = strings.Replace(in, "_", "", -1)
	}
	return in
}

func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
