package email

import (
	code "MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/common/code"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/config"
	"MuXiFresh-Be-2.0/common/globalKey"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

type EmailInfo struct {
	Host     string
	Port     string
	UserName string
	Password string
}

var eInfo EmailInfo

func Load(c config.Config) {
	copier.Copy(&eInfo, c.EmailConf)
}

func Send(Email string, Type string) error {
	if Type != globalKey.AuthRegister && Type != globalKey.AuthChPass {
		return fmt.Errorf("invalid email type")
	}
	//生成一个验证码
	randCode := code.RandStringBytes(8)
	html := randCode
	subject := fmt.Sprintf("%v验证码", Type)
	//发送
	e := &email.Email{
		To:      []string{Email},
		From:    eInfo.UserName,
		Subject: subject,
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(eInfo.Host+":"+eInfo.Port, smtp.PlainAuth("", eInfo.UserName, eInfo.Password, eInfo.Host))
	if err != nil {
		return err
	}
	//存到redis
	err = code.SetEm(Type, Email, randCode)
	return err
}
