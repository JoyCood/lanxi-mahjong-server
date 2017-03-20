//web server
package req

import (
	"basic/utils"
	"data"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"github.com/gorilla/websocket"
)

type web_msg_test struct {
	Userid  string `json:userid` //角色ID
	Itemid  uint32 `json:itemid` //物品id或类型
	Amount  int32  `json:amount` //数量
}

func TestWeb(t *testing.T) {
	web_client()
}

//召唤机器人
func web_client() {
	//var addr string = "localhost:8084"
	addr := data.Conf.Host + ":" + strconv.Itoa(data.Conf.WebPort)
	fmt.Println("addr -> ", addr)
	u := url.URL{Scheme: "ws", Host: addr, Path: "/"}
	var Key  string = "XG0e2Ye/KAUJRXaMNnJ5UH1haBvh2FXOoAggE6f2Utw"
	var Now string = strconv.FormatInt(utils.Timestamp(), 10)
	var Num string = "33"
	var Sign string = utils.Md5(Key+Now)
	var Token string = Sign+Now+Num
	c, _, err := websocket.DefaultDialer.Dial(u.String(),
	http.Header{"Token":{Token}})
	//fmt.Printf("c -> %+v\n", c)
	if err != nil {
		fmt.Printf("dial err -> %v\n", err)
	}
	msg1 := web_msg_test{
		Userid: "16007",
		Itemid: 4,
		Amount: 100,
	}
	msg2, _ := json.Marshal(msg1)
	msg := Key+"|"+string(msg2)
	if c != nil {
		c.WriteMessage(websocket.TextMessage, []byte(msg))
		defer c.Close()
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		fmt.Printf("recv: %s", message)
	}
}
