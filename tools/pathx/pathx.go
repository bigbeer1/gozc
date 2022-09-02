package pathx

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// LoadTemplate gets template content by the specified file.
func LoadTemplate(category, file, builtin string) (string, error) {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !FileExists(file) {
		return builtin, nil
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

var goctlHome string

// GetTemplateDir returns the category path value in GoctlHome where could get it by GetGoctlHome.
func GetTemplateDir(category string) (string, error) {
	env := os.Getenv("GOZC_PATH")
	if env == "" {
		return "", errors.New("请先设置环境变量GOZC_PATH路径指向Tpl文件夹")
	}

	return filepath.Join(env, category), nil
}

const (
	NL              = "\n"
	goctlDir        = ".goctl"
	gitDir          = ".git"
	autoCompleteDir = ".auto_complete"
	cacheDir        = "cache"
)
