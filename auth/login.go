package auth

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

//登录获取鉴权的cookie
func Login(client *http.Client,username string,password string) (*http.Cookie,error){
	const loginUrl="https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page="
	loginPageResp,err:=client.Get(loginUrl)
	if err!=nil{
		fmt.Printf("获取登录页失败:%v",err)
		return &http.Cookie{},err
	}
	defer loginPageResp.Body.Close()

	//loginPageResp是返回的消息流，转化成[]bytes字节切片
	htmlPageBytes,_:=io.ReadAll(loginPageResp.Body)

	//解析网页数据，提取lt,execution的值
	doc,_:=goquery.NewDocumentFromReader(strings.NewReader(string(htmlPageBytes)))

	lt,_:=doc.Find("input[name=lt]").Attr("value")
	execution,_:=doc.Find("input[name=execution]").Attr("value")

	formData:=url.Values{}
	formData.Set("username",username)
	formData.Set("password",password)
	formData.Set("lt",lt)
	formData.Set("execution",execution)
	formData.Set("_eventId","submit")
	formData.Set("submit","登录")

	//构建post请求
	postReq,_:=http.NewRequest("POST",loginUrl,strings.NewReader(formData.Encode()))
	postReq.Header.Set("Content-Type","application/x-www-form-urlencoded")
	postReq.Header.Set("Origin","https://account.ccnu.edu.cn")
	postReq.Header.Set("Referer","https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	postReq.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")

	//发送post请求，拿到响应体
	resp,err:=client.Do(postReq)
	if err!=nil{
		fmt.Printf("post请求发送失败:%v",err)
		return &http.Cookie{},err
	}
	defer resp.Body.Close()

	//拿到cookie
	u2,err:=url.Parse("http://kjyy.ccnu.edu.cn")
	if err!=nil{
		fmt.Printf("解析url失败:%v",err)
		return &http.Cookie{},err
	}
	cookies:=client.Jar.Cookies(u2)
	return cookies[0],nil
}