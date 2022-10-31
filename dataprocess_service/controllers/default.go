package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"time"
)

type MainController struct {
	beego.Controller
	Body        []byte
	RequestID   string
	RequestTime time.Time
	code        int
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	//logger.
	//log.Fatal("this is Fatal")
	c.TplName = "index.tpl"
}

type HttpData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Result struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	OriginErr string `json:"origin_error"`
}

func (c *MainController) Prepare() {
	c.code = 200
	c.Body = c.Ctx.Input.RequestBody
	req := c.Ctx.Request
	addr := req.RemoteAddr // "IP:port" "192.168.1.150:8889"
	fmt.Println(addr)
	c.RequestID = uuid.New().String()
	//log.SetOutput("uuid")
	//formatter := &MyJSONFormatter{
	//	JSONPrefix: "ip: ",
	//	Otherdata:  addr,
	//	RequestID:  c.RequestID,
	//}

	//log.SetFormatter(formatter)

	//logger := log.WithFields(log.Fields{
	//	//"user_id":    "",
	//	"ip":         addr,
	//	"request_id": c.RequestID,
	//})
	//logger.Info("request start.......")
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	c.RequestTime = time.Now()
	c.Ctx.Output.Header("X-Request-ID", c.RequestID)
}

type MyJSONFormatter struct {
	JSONPrefix string
	Otherdata  string
	RequestID  string
}

func (my *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	// fmt.Println(entry.Data["msg"])

	entry.Data["msg"] = fmt.Sprintf("%s%s", my.JSONPrefix, my.Otherdata)
	entry.Data["request_id"] = fmt.Sprintf("%s", my.RequestID)

	json, err := jsoniter.Marshal(&entry.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON , %w", err)
	}
	return append(json, '\n'), nil

}

type LogTrace struct {
}

// 实例化
func NewLogTrace() LogTrace {
	return LogTrace{}
}

//func (hook LogTrace) Levels() []log.Level {
//	return logrus.AllLevels
//}

func (hook LogTrace) Fire(entry *log.Entry) error {
	ctx := entry.Context
	if ctx != nil {
		traceId := ctx.Value("trace_id")
		if traceId != nil {
			entry.Data["trace_id"] = traceId
		}
	}
	return nil
}
