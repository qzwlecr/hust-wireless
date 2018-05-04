package hust_wireless

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (u *User) Init(userName string, userPasswd string) {
	u.name = userName
	u.password = userPasswd
	u.index = ""
}

func (u *User) Login() (status int, queryString string) {
	try, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal("请连接到HUST_WIRELESS!!!!")
	}
	defer try.Body.Close()

	tryBody, _ := ioutil.ReadAll(try.Body)
	queryString = matchURL(string(tryBody))
	if queryString == "" {
		return 0, "已经登录,请勿重复登录!"
	}

	req := url.Values{
		"userId":      {u.name},
		"password":    {u.password},
		"queryString": {queryString},
		"service":     {""},
		"operatorPwd": {""},
		"validcode":   {""},
	}

	resp, err := http.PostForm("http://192.168.50.3:8080/eportal/InterFace.do?method=login", req)
	if err != nil {
		log.Fatal("无法连接到登录服务器!!!!")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	resS := Responses{}
	err = json.Unmarshal(respBody, &resS)
	if resS.Result == "success" {
		u.index = resS.UserIndex
		if strings.Index(resS.Message, "在线") != -1 {
			return 4, "登录失败,请重试!"
		} else {
			return 1, "登录成功!"
		}
	} else {
		if resS.Message == "密码不匹配,请输入正确的密码!" {
			u.index = resS.UserIndex
			return 2, "密码不匹配,请重试!"
		} else {
			if resS.Message == "帐号不存在,请输入正确的帐号。" {
				return 3, "账户不存在,请重试!"
			}
		}
	}
	return 4, "登录失败,请重试!"
}

func (u *User) Logout() {
	r := url.Values{"userIndex": {u.index}}

	_, err := http.PostForm("http://192.168.50.3:8080/eportal/InterFace.do?method=logout", r)
	if err != nil {
		log.Fatal("无法连接到登录服务器!!!!")
	}

	return
}

func matchURL(s string) string {
	if strings.Index(s, "192.168.50.3:8080") == -1 {
		return ""
	}
	i := strings.Index(s, "'")
	if i >= 0 {
		j := strings.Index(s[i+1:], "'")
		if j >= 0 {
			return s[i+1 : j+i+1]
		}
	}
	return ""
}
