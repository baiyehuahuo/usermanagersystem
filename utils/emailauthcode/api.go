package emailauthcode

import (
	"math/rand"
	"time"
	"usermanagersystem/utils/rediscontrol"
)

var EACC EmailAuthCodeController

type EmailAuthCodeController interface {
	SendAuthCodeByEmail(target string) error
}

func EmailAuthCodeControllerCreate() {
	rand.Seed(time.Now().Unix())
	EACC = &emailAuthCodeControllerImpl{
		rc: rediscontrol.New(),
	}
}
