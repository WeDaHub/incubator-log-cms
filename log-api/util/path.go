package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"errors"
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

func CreateFolder(folder string) string{
	exists, _ :=  PathExists(folder)
	if !exists {
		os.MkdirAll(folder, os.ModePerm)
	}
	return folder
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	path = strings.Replace(path, "\\", "/", -1)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func GetDataGlobalFolder() string{
	path, _ := GetCurrentPath()
	return CreateFolder(path + "./data/")
}

func GetLogRootFolder() string{
	return CreateFolder(GetDataGlobalFolder() + "log/")
}

func GetTodayLogFolder() string{
	toady := time.Now().Format("2006-01-02/")
	return CreateFolder(GetLogRootFolder() + toady)
}

func GetConfigFolder() string{
	return CreateFolder(GetDataGlobalFolder() + "config/")
}

