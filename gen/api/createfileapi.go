package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateFileApi(modelList map[string]*CodeTuple, srcPath string) error {

	dirAbs := filepath.Dir(srcPath)

	dirAbs = filepath.Join(dirAbs, "api")

	is, _ := IsPathExist(dirAbs)
	if is == false {
		err := os.Mkdir(dirAbs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for tableName, codes := range modelList {

		name := fmt.Sprintf("%v.api", SafeString(tableName))
		filename := filepath.Join(dirAbs, name)
		err := ioutil.WriteFile(filename, []byte(codes.Api), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vaddlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiInsert), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vdellogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiDelete), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vuplogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiUpdate), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vInfologic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiFindOne), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vlistlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.ApiFindList), os.ModePerm)
		if err != nil {
			return err
		}
	}
	fmt.Println("ApiDone.")
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
