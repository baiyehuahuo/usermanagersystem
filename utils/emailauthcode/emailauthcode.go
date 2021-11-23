package emailauthcode

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"usermanagersystem/consts"
	"usermanagersystem/utils/configread"
	"usermanagersystem/utils/rediscontrol"

	"github.com/jordan-wright/email"
)

type emailAuthCodeControllerImpl struct {
	rc rediscontrol.RedisController
}

func (e *emailAuthCodeControllerImpl) SendAuthCodeByEmail(target string) error {
	config := configread.Config.EmailConfig
	em := email.NewEmail()
	em.From = fmt.Sprintf("%s<%s>", consts.AuthEmailUser, config.Email)

	em.To = []string{target}

	em.Subject = consts.AuthEmailSubject

	authCode := rand.Intn(consts.AuthCodeRandRange)
	em.Text = []byte(fmt.Sprintf("Hello World!! %06d", authCode))

	err := em.Send(config.Addr, smtp.PlainAuth("", config.UserName, config.Password, config.Host))
	if err != nil {
		return err
	}
	return nil
}
