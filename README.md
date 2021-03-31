# get-GoCN-news
爬取获得gocn的新闻，并同步到博客，每日更新从我做起

具体部署方法：https://blog.csdn.net/weixin_44024220/article/details/105960728

### 参考库：

https://github.com/anaskhan96/soup

博客框架：hexo

### 特性

支持每天定时爬取

自动推送到hexo页面上

### 效果展示：

https://greenpipig.github.io/

### 使用方法：

首先修改update.sh中的路径文件，修改为自己的博客路径

每次检索时间为3小时一次

go build main.go

nohup ./main &

#### 踩坑：

对于html解析时该库无法解析空格，推荐使用doc.FindAllStrict此方法

### 步骤

```
1、get-GoCN-news\getNews
	htmlPara.go中github.com/greenpipig/soup修改为自己github仓库的soup,fork一个https://github.com/anaskhan96/soup
2、get-GoCN-news\vendor\github.com\greenpipig\soup路径名改为自己的，如get-GoCN-news\vendor\github.com\JiangJiaWei520\soup
3、get-GoCN-news\vendor路径下vendor.json中属性字段rootPath的值github.com/greenpipig/get-GoCN-news改为github.com/JiangJiaWei520/get-GoCN-news
4、go.mod文件去掉
5、get-GoCN-news路径下main.go文件中的
	"github.com/greenpipig/get-GoCN-news/GoCN-news"   
	"github.com/greenpipig/get-GoCN-news/cron" 		  
	"github.com/greenpipig/get-GoCN-news/getNews"	  
	改为
		"github.com/JiangJiaWei520/get-GoCN-news/GoCN-news"   
	"github.com/JiangJiaWei520/get-GoCN-news/cron" 		  
	"github.com/JiangJiaWei520/get-GoCN-news/getNews"
6、安装go环境
7、cd get-GoCN-news
8、go mod init
9、go mod tidy
10、go mod vendor
	go mod init 初始化模块（例如把原本dep管理的依赖关系转换过来）
	go mod tidy 增加缺失的包，移除没用的包
	go mod vendor 把依赖拷贝到 vendor/ 目录下
注意：删除go.mod下require github.com/greenpipig/soup v1.1.2-0.20200506083017-b50f4923c2f9
11、go build main.go 非首次可以执行main.exe
	爬取
12、nohup ./main & 
	生成GoCN-news.log文件
13、执行./update.sh

**go mod
```