package main

import (
	"fmt"

	"github.com/tidwall/gjson"

	common "github.com/iotexproject/w3bstream/examples/wasm_common_go"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	common.Log(fmt.Sprintf("start received: %d", rid))
	message, err := common.GetDataByRID(rid)
	if err != nil {
		common.Log("error: " + err.Error())
		return -1
	}
	res := string(message)

	Points := gjson.Get(res, "Points")

	common.Log("wasm get Points from json: " + Connections.String())

	if Points.Int() > 1000 {
		common.SendTx(fmt.Sprintf(
			`{
				"to": "%s",
				"value": "0",
				"data": "1249c58b"
			}`,
			"0x3ac682102B7Bfef43D2f48e17e14CDcC055D11fd",//W3BStreamNFT contract address
		))
		common.Log("W3BStreamNFT has been sent to the private key account address")
	}

	return 0
}