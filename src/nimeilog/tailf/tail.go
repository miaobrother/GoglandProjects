package tailf

import (
	//"fmt"
	"github.com/hpcloud/tail"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

type TextMsg struct {
	Msg string
	Topic string
}

var(
	tailObjMgr *TailObjMgr
)

type CollConf struct {
	LogPath string
	Topic string
}

type TailObj struct {
	tail *tail.Tail
	conf CollConf

}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan chan *TextMsg
}

func InitTail(conf []CollConf,chanSize int)(err error)  {
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg,chanSize),
	}
	if len(conf) == 0{
		err = fmt.Errorf("invalid config for log coll,conf is %v\n",conf)
		return
	}

	for _,v := range conf{
		obj := &TailObj{
			conf:v,
		}
		tails,errs := tail.TailFile(v.LogPath,tail.Config{
			ReOpen:true,
			Follow:true,
			MustExist:false,
			Poll:true,
		})
		if errs != nil{
			fmt.Printf("tail got logPath failed,err is:%v\n",err)
			return
		}
		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs,obj)


		go readFromTail(obj)

	}
	return
}



func readFromTail(tailObj *TailObj){

	for true{
		line,ok := <- tailObj.tail.Lines
		if !ok{
			logs.Warn("tail file close reopen,filename:%s\n",tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		textMsg := &TextMsg{
			Msg:line.Text,
			Topic:tailObj.conf.Topic,

		}
		tailObjMgr.msgChan <- textMsg

	}

}

func GetOneLine() (msg *TextMsg)  {
	msg = <- tailObjMgr.msgChan
	//fmt.Println(msg.Msg,msg.Topic)
	return
}