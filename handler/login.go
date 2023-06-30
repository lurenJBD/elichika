package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/utils"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func StartUp(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var mask64 string
	req := gjson.Parse(reqBody)
	req.ForEach(func(key, value gjson.Result) bool {
		if value.Get("mask").String() != "" {
			mask64 = value.Get("mask").String()
			return false
		}
		return true
	})
	// fmt.Println("Request data:", req.String())
	// fmt.Println("Mask:", mask64)

	mask, err := base64.StdEncoding.DecodeString(mask64)
	if err != nil {
		panic(err)
	}
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	// fmt.Println("Random Bytes:", randomBytes)

	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	// fmt.Println("Session Key:", newKey64)

	startupBody := utils.ReadAllText("assets/startup.json")
	startupBody, _ = sjson.Set(startupBody, "authorization_key", newKey64)
	resp := SignResp(ctx.GetString("ep"), startupBody, config.StartUpKey)
	// fmt.Println("Response:", resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func Login(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var mask64 string
	req := gjson.Parse(reqBody)
	req.ForEach(func(key, value gjson.Result) bool {
		if value.Get("mask").String() != "" {
			mask64 = value.Get("mask").String()
			return false
		}
		return true
	})
	fmt.Println("Request data:", req.String())
	fmt.Println("Mask:", mask64)

	mask, err := base64.StdEncoding.DecodeString(mask64)
	if err != nil {
		panic(err)
	}
	randomBytes := encrypt.RSA_DecryptOAEP(mask, "privatekey.pem")
	fmt.Println("Random Bytes:", randomBytes)

	serverEventReceiverKey, err := hex.DecodeString(config.ServerEventReceiverKey)
	if err != nil {
		panic(err)
	}

	unknownKey, err := hex.DecodeString(config.UnknownKey)
	if err != nil {
		panic(err)
	}

	newKey := utils.Xor(randomBytes, []byte(config.SessionKey))
	newKey = utils.Xor(newKey, serverEventReceiverKey)
	newKey = utils.Xor(newKey, unknownKey)
	newKey64 := base64.StdEncoding.EncodeToString(newKey)
	// fmt.Println("Session Key:", newKey64)

	loginBody := utils.ReadAllText("assets/login.json")
	loginBody, _ = sjson.Set(loginBody, "session_key", newKey64)
	resp := SignResp(ctx.GetString("ep"), loginBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
