package log

import (
	"fmt"
	"log"
	"log-api/util"
	"os"
	"time"
)

type logConfig struct{
	log *log.Logger
	todayLogFolder string
}

var conf logConfig

func DEBUG(keyword string, format string, a ...interface{}){
	s := fmt.Sprintf(format, a...)
	getLog().Printf("==%s==[DEBUG]%s", keyword, s)
}

func INFO(keyword string, format string, a ...interface{}){
	s := fmt.Sprintf(format, a...)
	getLog().Printf("==%s==[INFO]%s", keyword, s)
}

func WARN(keyword string, format string, a ...interface{}){
	s := fmt.Sprintf(format, a...)
	getLog().Printf("==%s==[WARN]%s", keyword, s)
}

func ERROR(keyword string, format string, a ...interface{}){
	s := fmt.Sprintf(format, a...)
	getLog().Printf("==%s==[ERROR]%s", keyword, s)
}

func getLog() *log.Logger{
	todayLogFolder := util.GetTodayLogFolder()
	if nil == conf.log || conf.todayLogFolder != todayLogFolder {
		logPath := todayLogFolder + time.Now().Format("2006-01-02 15-04-05") + ".log"
		logFile, err := os.Create(logPath)
		if err != nil {
			fmt.Println(err);
			return nil
		}
		conf.log = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
		conf.todayLogFolder = todayLogFolder
	}
	return conf.log
}

