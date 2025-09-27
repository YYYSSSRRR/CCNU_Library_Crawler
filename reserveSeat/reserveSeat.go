package reserveseat

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Reserve(client *http.Client,devid string ,start string, end string,start_time string, end_time string)(error){
	//构建url
	const baseQueryUrl="http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx"
	u,_:=url.Parse(baseQueryUrl)
	q:=u.Query()
	q.Set("dialogid","")
	q.Set("dev_id",devid)
	q.Set("lab_id","")
	q.Set("kind_id","")
	q.Set("room_id","")
	q.Set("type","dev")
	q.Set("prop","")
	q.Set("test_id","")
	q.Set("term","")
	q.Set("Vnumber","")
	q.Set("classkind","")
	q.Set("test_name","")
	q.Set("start",start)
	q.Set("end",end)
	q.Set("start_time",start_time)
	q.Set("end_time",end_time)
	q.Set("up_file","")
	q.Set("memo","")
	q.Set("act","set_resv")
	q.Set("_",fmt.Sprint(time.Now().UnixMilli()))
	u.RawQuery=q.Encode()
	req,_:=http.NewRequest("GET",u.String(),nil)
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With","XMLHttpRequest")

	//发请求
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Printf("请求发送失败")
		return err
	}
	defer resp.Body.Close()

	data,_:=io.ReadAll(resp.Body)
	fmt.Printf("%s\n",data)

	return nil;

}