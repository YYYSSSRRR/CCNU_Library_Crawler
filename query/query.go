package query

import (
	"encoding/json"
	"fmt"
	"io"
	"library/types"
	"net/http"
	"net/url"
	"time"
)

type rawUser struct{
	Name string `json:"name"`
}

func Query(client *http.Client,StudentId string) (*types.User,error){
	const baseQueryUrl string ="http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/data/searchAccount.aspx"
	u,_:=url.Parse(baseQueryUrl)
	q:=u.Query()
	q.Set("type","logonname")
	q.Set("ReservaApply","ReservaApply")
	q.Set("term",StudentId)
	q.Set("_",fmt.Sprint(time.Now().UnixMilli()))
	u.RawQuery=q.Encode()
	//最终的url就是u.String()

	req,_:=http.NewRequest("GET",u.String(),nil)
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With","XMLHttpRequest")

	result,err:=client.Do(req)
	if err!=nil{
		fmt.Printf("发送请求失败:%v",err)
		return nil,err
	}
	defer result.Body.Close()

	rawInfo,_:=io.ReadAll(result.Body)
	// fmt.Printf("%s\n",rawInfo)

	var rawUser []rawUser
	err1:=json.Unmarshal(rawInfo,&rawUser)
	if err1!=nil{
		fmt.Printf("json反序列化失败:%v",err1)
		return nil,err
	}

	if len(rawUser) == 0 {
		// fmt.Printf("未找到\n")
		// 返回 nil 表示未找到
		return nil, nil
	}

	user := types.User{
		StudentId: StudentId,
			Name: rawUser[0].Name,
			Grade: StudentId[:4],
	}

	return &user,nil
}