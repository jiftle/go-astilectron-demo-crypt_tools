package main

import (
	"encoding/json"
	//	"io/ioutil"
	//	"os"
	//	"os/user"
	//	"path/filepath"
	//	"sort"
	//	"strconv"

	// "github.com/asticode/go-astichartjs"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
)

type CryptParam struct {
	Plain  string `json:"plain,omitempty"`
	Key    string `json:"key,omitempty"`
	Cipher string `json:"cipher,omitempty"`
}

// handleMessages handles messages
// 处理消息
func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "btn_crypt":
		// Unmarshal payload
		// 解码载荷
		var param CryptParam
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &param); err != nil {
				payload = err.Error()
				return
			}

		}

		// ------------- 执行加密
		// x := HexStrToBytes("11111111111111111111111111111111")
		// key := HexStrToBytes("11111111111111111111111111111111")
		x := HexStrToBytes(param.Plain)
		key := HexStrToBytes(param.Key)

		// ----------- 拼凑24字节的密码 -------------
		var k []byte = make([]byte, 0)
		k = append(k, key[0:16]...)
		k = append(k, key[0:8]...)
		astilog.Infof("--> 切片: %v\n", k)

		astilog.Infof("x: %v\nkey: %v\n", x, k)
		x1 := encrypt_triple_des(x, k)
		c1 := BytesToHexStr(x1)
		astilog.Infof("-->> 密文: %v\n", string(c1))
		//		x2 := decrypt_triple_des(x1, k)
		//		astilog.Infof("-->> 明文: %v\n", BytesToHexStr(x2))

		// -------------- 返回值 -----------------
		//payload = BytesToHexStr(x2)
		payload = c1
		break
	}
	return
}
