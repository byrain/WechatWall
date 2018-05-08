package util

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"net/http"
	"sort"

	"github.com/go-chi/render"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

var (
	msgStartElementLiteral = []byte("<xml>")
	msgEndElementLiteral   = []byte("</xml>")

	msgToUserNameStartElementLiteral = []byte("<ToUserName>")
	msgToUserNameEndElementLiteral   = []byte("</ToUserName>")

	msgEncryptStartElementLiteral = []byte("<Encrypt>")
	msgEncryptEndElementLiteral   = []byte("</Encrypt>")

	cdataStartLiteral = []byte("<![CDATA[")
	cdataEndLiteral   = []byte("]]>")
)

type CipherRequestHttpBody struct {
	XMLName            struct{} `xml:"xml"`
	ToUserName         string   `xml:"ToUserName"`
	Base64EncryptedMsg []byte   `xml:"Encrypt"`
}

func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func RenderJSON(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, msg)
	return
}

func SecureCompare(given, actual []byte) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare(given, actual) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare(actual, actual) == 1 {
		return false
	}
	return false
}

func ProcessTextMsg(mixedMsg *core.MixedMsg) {
	// var haveObject = request.GetText(mixedMsg)

	// userId := textObject.MsgHeader.
	return
}
