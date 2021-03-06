package zfService

import (
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils"
	"funnel/app/utils/fetch"
	"net/http"
	"strings"
)

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()
	_, err := f.Get(zf.ZfLoginHome())

	if err != nil {
		return nil, err
	}
	if len(f.Cookie) < 1 {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}

	s, err := f.Get(zf.ZfLoginKaptcha())
	if err != nil {
		return nil, err
	}
	captcha, err := utils.BreakCaptcha(s)
	if err != nil {
		return nil, err
	}
	loginData := genLoginData(username, captcha, password, f)

	s, err = f.PostForm(zf.ZfLoginHome(), loginData)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(s), "验证码输入错误") {
		return nil, errors.ERR_WRONG_Captcha
	}
	if strings.Contains(string(s), "用户名或密码不正确") {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	var cookie *http.Cookie
	for _, v := range f.Cookie {
		if v.Name == "JSESSIONID" {
			cookie = v
		}
	}
	if cookie == nil {
		return nil, errors.ERR_UNKNOWN_LOGIN_ERROR
	}
	return service.SetUser(service.ZFPrefix, username, password, cookie)
}
