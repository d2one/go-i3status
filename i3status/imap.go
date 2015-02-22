package i3status

import (
	"code.google.com/p/go-imap/go1/imap"
	"strconv"
	"strings"
	"time"
)

const (
	Addr = "MAIL SERVER"
	User = "LOGIN"
	Pass = "PASSWORD"
)

type ImapWidget struct {
	BaseWidget
}

func NewImapWidget() *ImapWidget {
	instanceCount++
	w := ImapWidget{
		*NewBaseWidget(),
	}
	return &w
}

func (w *ImapWidget) basicLoop() {

	var (
		c   *imap.Client
		cmd *imap.Command
		rsp *imap.Response
	)

	c = Dial(Addr)
	c.Login(User, Pass)
	defer c.Logout(30 * time.Second)

	cmd, _ = imap.Wait(c.List("", "*"))

	msg := NewMessage()
	msg.Name = "MAIL"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		msg.FullText = ""
		for _, rsp = range cmd.Data {
			folder := rsp.MailboxInfo()
			c.Select(folder.Name, true)
			if c.Mailbox.Unseen > 0 {
				msg.Color = "#FFF000000"
				msg.FullText = "MAIL"
			}
		}

		w.Output <- *msg
		time.Sleep(10000 * time.Millisecond)
	}
}

func (w *ImapWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}

func Dial(addr string) (c *imap.Client) {
	var err error
	if strings.HasSuffix(addr, ":993") {
		c, err = imap.DialTLS(addr, nil)
	} else {
		c, err = imap.Dial(addr)
	}
	if err != nil {
		panic(err)
	}
	return c
}
