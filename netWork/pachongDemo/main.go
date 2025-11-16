package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Movie 存储电影信息的结构体
type Movie struct {
	Title  string // 电影名称
	Score  string // 评分
	Intro  string // 简介
	Rank   int    // 排名
	Detail string // 详情页 URL
}

func main() {
	//初始化一个上下文，设置协程超时时间为30秒，超时自动取消所有任务
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	//确保main方法结束之后会退出所有的协程
	defer cancelFunc()

	//定义任务列表
	var urls []string
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("https://movie.douban.com/top250?start=%d", 25*i)
		urls = append(urls, url)
	}

	//并发控制，限制一下线程个数
	const maxConcurrent = 3

	//设置一个信号channel，限制并发数量
	signalChan := make(chan struct{}, maxConcurrent)

	//创建一个接受爬取结果的channel
	resultChan := make(chan []Movie, len(urls))

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			//进入协程，我们占用一下信号量，往信号channel放入一个空结构体
			//在并发中，我们一般使用sync,WaitGroup来保证所有任务都完成，用信号量chan来控制并发的数量
			signalChan <- struct{}{}
			defer func() { <-signalChan }()

			//检查contex是否已经取消或者终止，或者超时
			select {
			case <-ctx.Done():
				fmt.Println("任务取消，原因：", ctx.Err())
				return
			default:
			}
			//爬取单页数据
			movies, err := crawPage(url)
			if err != nil {
				fmt.Printf("爬取失败 %s：%v\n", url, err)
				return

			}

			//将结果放入resultChan
			resultChan <- movies
		}(url)
	}

	//等所有任务完成，关闭resultChan
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	//收集所有结果，保存到文件
	var allMovie []Movie
	for pageMovies := range resultChan {
		allMovie = append(allMovie, pageMovies...)
	}

	//写入json文件
	if err := saveToFile(allMovie, "douban_top250.json"); err != nil {
		fmt.Printf("保存文件失败：%v\n", err)
		return
	}
	fmt.Printf("爬取完成，共 %d 部电影，已保存到 douban_top250.json\n", len(allMovie))
}

// 爬取网页
func crawPage(url string) ([]Movie, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//构建请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("请求构建失败，原因：%v", err)
		return nil, err
	}

	//发送请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf("请求send失败，原因：%v", err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("状态码错误%d", response.StatusCode)
	}

	//用goquery解析html
	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(response.Body))
	if err != nil {
		return nil, err
	}

	var movies []Movie
	// 定位电影列表项（通过浏览器 F12 查看页面结构获取选择器）
	doc.Find(".grid_view li").Each(func(i int, s *goquery.Selection) {
		// 提取排名（从 1 开始）
		rankText := s.Find(".pic em").Text()
		rank, _ := strconv.Atoi(rankText)

		// 提取电影名称
		title := s.Find(".title:nth-child(1)").Text()

		// 提取评分
		score := s.Find(".rating_num").Text()

		// 提取简介（可能为空，需要处理）
		intro := s.Find(".inq").Text()

		// 提取详情页 URL
		detail, _ := s.Find(".pic a").Attr("href")

		movies = append(movies, Movie{
			Rank:   rank,
			Title:  title,
			Score:  score,
			Intro:  intro,
			Detail: detail,
		})
	})

	return movies, nil

}

// saveToFile 将电影列表保存为 JSON 文件（IO 操作）
func saveToFile(movies []Movie, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 格式化 JSON（缩进 2 空格，便于阅读）
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(movies)
}
