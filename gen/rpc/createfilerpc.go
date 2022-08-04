package rpc

import (
	"fmt"
	"gozc/tools/pathx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateFileRpc(modelList map[string]*CodeTuple, srcPath string) error {

	dirAbs := filepath.Dir(srcPath)

	err := pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	dirAbs = filepath.Join(dirAbs, "rpc")

	is, _ := IsPathExist(dirAbs)
	if is == false {
		err = os.Mkdir(dirAbs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for tableName, codes := range modelList {

		name := fmt.Sprintf("%v.proto", SafeString(tableName))
		filename := filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.Rpc), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vaddlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.RpcInsert), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vdellogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.RpcDelete), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vuplogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.RpcUpdate), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vfindonelogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.RpcFindOne), os.ModePerm)
		if err != nil {
			return err
		}

		name = fmt.Sprintf("%vlistlogic.go", SafeString(tableName))
		filename = filepath.Join(dirAbs, name)
		err = ioutil.WriteFile(filename, []byte(codes.RpcFindList), os.ModePerm)
		if err != nil {
			return err
		}
	}
	fmt.Println("RpcDone.")
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
