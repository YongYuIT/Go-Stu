package Probe

import (
	"testing"
	"github.com/emersion/go-imap/client"
	"fmt"
)

func TestImapConn(t *testing.T) {
	imap_add := "imap.mail.us-west-2.awsapps.com:993"
	c, err := client.DialTLS(imap_add, nil)
	defer c.Logout()
	if err != nil {
		t.Error(fmt.Errorf("conn fail to %s for %s\n", imap_add, err.Error()))
		return
	}
	user_name := "******@******"
	user_pwd := "******"
	if err := c.Login(user_name, user_pwd); err != nil {
		t.Error(fmt.Errorf("conn fail to %s for %s\n", user_name, err.Error()))
		return
	}
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		t.Error(fmt.Errorf("open fail to %s for %s\n", user_name, err.Error()))
		return
	}
	uid_next := mbox.UidNext
	fmt.Printf("success! %d \n", uid_next)
}
