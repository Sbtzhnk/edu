package profile

import (
	"github.com/google/uuid"
)

type profile struct {
	isBlock bool
	login   string
	passwd  string
	uuid    uuid.UUID
}

func (p *profile) lockProfile() bool {
	p.isBlock = true
	return true
}

func (p *profile) unlockProfile() bool {
	p.isBlock = false
	return true
}

func CreateProfile(login string, passwd string, isBLock bool) *profile {
	profile := profile{
		isBlock: isBLock,
		login:   login,
		passwd:  passwd,
		uuid:    uuid.New(),
	}
	return &profile
}

func (p profile) ChangeStatus(flag bool) {
	if flag {
		p.lockProfile()
	} else {
		p.unlockProfile()
	}
}
