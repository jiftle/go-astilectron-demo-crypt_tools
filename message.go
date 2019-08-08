package main

import (
	"encoding/json"
	//	"io/ioutil"
	//	"os"
	//	"os/user"
	//	"path/filepath"
	//	"sort"
	//	"strconv"
	"strings"

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

type RetValue struct {
	Ret     bool   `json:"ret"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

// handleMessages handles messages
// 处理消息
func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "btn_crypt":
		// 定义返回值
		var ret RetValue

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
		plainLen := strings.Count(param.Plain, "") - 1
		keyLen := strings.Count(param.Key, "") - 1
		// ------------- 参数合法性校验 ----------------
		if (plainLen == 0 || plainLen%2 != 0) || (keyLen == 0 || keyLen%2 != 0) {
			ret.Ret = false
			ret.Message = "必须是2的倍数"

			payload = ret
			return
		}

		// ------------- 执行加密
		plain := HexStrToBytes(param.Plain)
		key := HexStrToBytes(param.Key)

		// ----------- 拼凑24字节的密码 -------------
		var k []byte = make([]byte, 0)
		k = append(k, key[0:16]...)
		k = append(k, key[0:8]...)
		//astilog.Infof("--> 切片: %v\n", k)

		astilog.Infof("-->> 明文: %v\n密码: %v\n", BytesToHexStr(plain), BytesToHexStr(k))
		x1 := encrypt_triple_des(plain, k)
		cipher := BytesToHexStr(x1)
		astilog.Infof("-->> 密文: %v\n", string(cipher))

		// -------------- 返回值 -----------------
		ret.Ret = true
		ret.Message = ""
		ret.Value = cipher

		payload = ret
		break
	}
	return
}
