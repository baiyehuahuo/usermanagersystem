package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
	"usermanagersystem/consts"

	"github.com/pkg/errors"

	"github.com/jordan-wright/email"
)

var eacc EmailAuthCodeController

type EmailAuthCodeController interface {
	CheckAuthCodeByEmail(target string, code int) (success bool)
	SendAuthCodeByEmail(target string) (err error)
}

func EmailAuthCodeControllerCreate() {
	rand.Seed(time.Now().Unix())
	eacc = &emailAuthCodeControllerImpl{}
}

func GetEACC() EmailAuthCodeController {
	return eacc
}

// just impl !!!!!!!!!!!

type emailAuthCodeControllerImpl struct {
}

func (e *emailAuthCodeControllerImpl) CheckAuthCodeByEmail(target string, code int) (success bool) {
	correct, ok := cc.GetAuthCode(target)
	success = ok && correct == code
	if success {
		cc.DeleteAuthCode(target)
	}
	return success
}

func (e *emailAuthCodeControllerImpl) SendAuthCodeByEmail(target string) (err error) {
	config := Config.EmailConfig
	em := email.NewEmail()
	em.From = fmt.Sprintf("%s<%s>", consts.AuthEmailUser, config.Email)

	em.To = []string{target}

	em.Subject = consts.AuthEmailSubject

	authCode := rand.Intn(consts.AuthCodeRandRange)
	cc.SetAuthCode(target, authCode)
	em.Text = []byte(fmt.Sprintf("Hello World!! %06d", authCode))

	err = em.Send(config.Addr, smtp.PlainAuth("", config.UserName, config.Password, config.Host))
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	return nil
}
