# 诗歌爬虫
## 功能
爬取[古诗网](https://so.gushiwen.org/shiwen/default_0AA1.aspx)前20页的诗歌并输出到屏幕。

分为并发和非并发两种模式，其中`crawler.go`为非并发模式、`concurrence_crawler.go`为并发模式。非并发模式运行时间大约为2.5s，并发模式运行耗时大约为1.5s

**注意**：并发模式中一首诗可能不会被连续输出，原因是不同goroutine同时在打印，后期可以用互斥锁修复。