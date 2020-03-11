package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func fetch(url string) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}
func parseUrls(url string) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="contson"(.*?)>(.*?)</div>`)
	items := rp.FindAllStringSubmatch(body, -1)
	rContent := regexp.MustCompile("<p>(.*?)</p>")
	for _, item := range items {
		fmt.Println("--------------------------")
		content := rContent.FindAllStringSubmatch(item[2],-1)
		if len(content) == 0 {
			removeBr(item[2])
		}else{
			for _,v := range content {
				removeBr(v[1])
			}
		}
	}
}

func removeBr(s string){
	index := 0
	rp := regexp.MustCompile("<br />")
	list := rp.FindAllStringSubmatchIndex(s,-1)
	tem := s[index:]
	for _,v := range list {
		tem = s[index:v[0]]
		if strings.Contains(tem,"&nbsp;") {
			tem = tem[:len(tem)-6]
		}
		if len(tem) >= 1 {
			if !strings.Contains(tem,"　　"){
				tem = "　　" + tem
			}
		}
		fmt.Println(tem)
		index = v[1]
	}
	tem = s[index:]
	if strings.Contains(tem,"&nbsp;") {
		tem = tem[:len(tem)-6]
	}
	if len(tem) >= 1 {
		if !strings.Contains(tem,"　　") {
			tem = "　　" + tem
		}
	}
	fmt.Println(tem)
}

func main() {
	start := time.Now()
	for i:=1;i <= 20;i ++{
		parseUrls("https://so.gushiwen.org/shiwen/default_0AA" + strconv.Itoa(i) + ".aspx")
	}
	elapsed := time.Since(start)
	fmt.Println("\n***************")
	fmt.Printf("Took %s\n", elapsed)
	fmt.Println("***************")

}
