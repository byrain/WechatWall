package handler

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"net/http"

	"github.com/chanxuehong/util/security"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"

	"github.com/WechatWall/common"
	"github.com/WechatWall/util"
)

func Wx() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/ping", pingHandler)
	router.Get("/", wx_response)
	router.Post("/", wx_message)
	router.Get("/user_info", user_info)

	return router
}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "pong\n")
}

func user_info(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Println(req.Body)
	fmt.Println(req.RemoteAddr)
	Header := req.Header
	fmt.Println("Header: ", Header)
	fmt.Println("Header[Accept-Language]: ", Header["Accept-Language"])
	io.WriteString(w, "hi")
}

func wx_response(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	haveSignature := query.Get("signature")
	if haveSignature == "" {
		// errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
		return
	}

	timestamp := query.Get("timestamp")
	if timestamp == "" {
		// errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
		return
	}
	nonce := query.Get("nonce")
	if nonce == "" {
		// errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
		return
	}
	echostr := query.Get("echostr")
	if echostr == "" {
		// errorHandler.ServeError(w, r, errors.New("not found echostr query parameter"))
		return
	}

	token := common.Token

	wantSignature := util.Sign(token, timestamp, nonce)
	if !security.SecureCompareString(haveSignature, wantSignature) {
		fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
	}

	io.WriteString(w, echostr)
}

func wx_message(w http.ResponseWriter, req *http.Request) {

	buffer := util.TextBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer util.TextBufferPool.Put(buffer)
	if _, err := buffer.ReadFrom(req.Body); err != nil {
		// errorHandler.ServeError(w, r, err)
		return
	}
	requestBodyBytes := buffer.Bytes()
	msg := bytes.TrimSpace(requestBodyBytes)
	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		// log.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}

	switch mixedMsg.MsgHeader.MsgType {
	case "text":
		var haveObject = request.GetText(mixedMsg)
		// fmt.Println(haveObject)
		fmt.Printf("%s send %s. Avatar is nil", haveObject.MsgHeader.FromUserName, haveObject.Content)
	case "image":
		var haveObject = request.GetImage(mixedMsg)
		fmt.Println(haveObject)
	default:
		return
	}
}
