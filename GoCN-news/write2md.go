package GoCN_news

import (
	"log"
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

func WriteToMd(newsList []string, newsUrlList []string, title string) bool {
	todayTimestamp := time.Now().Unix()                               //获得时间戳
	todayTimeStr := time.Unix(todayTimestamp, 0).Format("2006-01-02") //把时间戳转换成时间,并格式化为年月日
	fileName := "./GOCN-news/" + todayTimeStr + "-" + "GOCN每日新闻.md"
	exist, _ := PathExists(fileName)
	if exist {
		return false
	} else {
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("WriteToMd error")
		}
		_, err = f.WriteString("---\n")
		_, err = f.WriteString("title: " + title + "\n")
		_, err = f.WriteString("date: " + time.Unix(todayTimestamp, 0).Format("2006-01-02 15:04:05") + "\n")
		_, err = f.WriteString("tags:\n")
		_, err = f.WriteString("- goNews\n")
		_, err = f.WriteString("---\n")
		//完成后，延迟关闭
		defer f.Close()
		for index, news := range newsList {
			log.Printf("news is:%s,url is:%s", news, newsUrlList[index])
			_, err = f.WriteString(news + ":" + newsUrlList[index] + "\n")
			if err != nil {
				log.Fatal("write file err %v", err)
			}
			f.WriteString("\n")
		}
		return true
	}

}
