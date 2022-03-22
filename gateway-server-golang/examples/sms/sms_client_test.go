package sms

import (
	"fmt"
	cmpp "github.com/bigwhite/gocmpp"
	cmpputils "github.com/bigwhite/gocmpp/utils"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSMSClient(t *testing.T) {
	log.Println("Client example start!")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go startAClient(i + 1)
	}
	wg.Wait()
	log.Println("Client example ends!")
}

const (
	user           string        = "900001"
	password       string        = "888888"
	connectTimeout time.Duration = time.Second * 2
)

func startAClient(idx int) {
	c := cmpp.NewClient(cmpp.V30)
	defer wg.Done()
	defer c.Disconnect()
	err := c.Connect(":8888", user, password, connectTimeout)
	if err != nil {
		log.Printf("client %d: connect error: %s.", idx, err)
		return
	}
	log.Printf("client %d: connect and auth ok", idx)

	t := time.NewTicker(time.Second * 5)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			//submit a message
			cont, err := cmpputils.Utf8ToUcs2("测试gocmpp submit")
			if err != nil {
				fmt.Printf("client %d: utf8 to ucs2 transform err: %s.", idx, err)
				return
			}
			p := &cmpp.Cmpp3SubmitReqPkt{
				PkTotal:            1,
				PkNumber:           1,
				RegisteredDelivery: 1,
				MsgLevel:           1,
				ServiceId:          "test",
				FeeUserType:        2,
				FeeTerminalId:      "13500002696",
				FeeTerminalType:    0,
				MsgFmt:             8,
				MsgSrc:             "900001",
				FeeType:            "02",
				FeeCode:            "10",
				ValidTime:          "151105131555101+",
				AtTime:             "",
				SrcId:              "900001",
				DestUsrTl:          1,
				DestTerminalId:     []string{"13500002696"},
				DestTerminalType:   0,
				MsgLength:          uint8(len(cont)),
				MsgContent:         cont,
			}

			_, err = c.SendReqPkt(p)
			if err != nil {
				log.Printf("client %d: send a cmpp3 submit request error: %s.", idx, err)
			} else {
				log.Printf("client %d: send a cmpp3 submit request ok", idx)
			}
			break
		default:
		}

		// recv packets
		i, err := c.RecvAndUnpackPkt(0)
		if err != nil {
			log.Printf("client %d: client read and unpack pkt error: %s.", idx, err)
			break
		}

		switch p := i.(type) {
		case *cmpp.Cmpp3SubmitRspPkt:
			log.Printf("client %d: receive a cmpp3 submit response: %v.", idx, p)

		case *cmpp.CmppActiveTestReqPkt:
			log.Printf("client %d: receive a cmpp active request: %v.", idx, p)
			rsp := &cmpp.CmppActiveTestRspPkt{}
			err := c.SendRspPkt(rsp, p.SeqId)
			if err != nil {
				log.Printf("client %d: send cmpp active response error: %s.", idx, err)
				break
			}
		case *cmpp.CmppActiveTestRspPkt:
			log.Printf("client %d: receive a cmpp activetest response: %v.", idx, p)

		case *cmpp.CmppTerminateReqPkt:
			log.Printf("client %d: receive a cmpp terminate request: %v.", idx, p)
			rsp := &cmpp.CmppTerminateRspPkt{}
			err := c.SendRspPkt(rsp, p.SeqId)
			if err != nil {
				log.Printf("client %d: send cmpp terminate response error: %s.", idx, err)
				break
			}
		case *cmpp.CmppTerminateRspPkt:
			log.Printf("client %d: receive a cmpp terminate response: %v.", idx, p)
		}
	}
}

var wg sync.WaitGroup
