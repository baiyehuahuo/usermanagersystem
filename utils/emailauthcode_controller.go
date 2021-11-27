package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
	"usermanagersystem/consts"

	"github.com/jordan-wright/email"
)

var EACC EmailAuthCodeController

type EmailAuthCodeController interface {
	SendAuthCodeByEmail(target string) error
}

func EmailAuthCodeControllerCreate() {
	rand.Seed(time.Now().Unix())
	EACC = &emailAuthCodeControllerImpl{}
}

// todo 以下是实现

type emailAuthCodeControllerImpl struct {
}

func (e *emailAuthCodeControllerImpl) SendAuthCodeByEmail(target string) error {
	config := Config.EmailConfig
	em := email.NewEmail()
	em.From = fmt.Sprintf("%s<%s>", consts.AuthEmailUser, config.Email)

	em.To = []string{target}

	em.Subject = consts.AuthEmailSubject

	authCode := rand.Intn(consts.AuthCodeRandRange)
	CC.SetAuthCode(target, authCode)
	em.Text = []byte(fmt.Sprintf("Hello World!! %06d", authCode))

	err := em.Send(config.Addr, smtp.PlainAuth("", config.UserName, config.Password, config.Host))
	if err != nil {
		return err
	}
	return nil
}
