package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Sign interface {
	Sign(request *http.Request)
}
type Wps4Sign struct {
	AK string
	SK string
}

func NewWps4Sign(AK string, SK string) *Wps4Sign {
	return &Wps4Sign{AK: AK, SK: SK}
}
func (s *Wps4Sign) Sign(request *http.Request) {
	//all, err := ioutil.ReadAll(request.Body)
	all, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataString := time.Now().UTC().Format(http.TimeFormat)
	method := request.Method
	path := request.URL.String()
	split := strings.Split(path, "docmini")
	path = split[1]
	request.Header.Set("Content-Type", "application/json")
	contentType := request.Header.Get("Content-Type")
	var authorization string
	//request.Body = ioutil.NopCloser(bytes.NewBuffer(all))
	request.Body = io.NopCloser(bytes.NewBuffer(all))
	body := string(all)
	if body == "" {
		authorization = "WPS-4" + method + path + contentType + dataString
	} else {
		hash := sha256.Sum256([]byte(body))
		authorization = "WPS-4" + method + path + contentType + dataString + hex.EncodeToString(hash[:])
	}
	mac := hmac.New(sha256.New, []byte(s.SK))
	mac.Write([]byte(authorization))
	authorization = "WPS-4 " + s.AK + ":" + hex.EncodeToString(mac.Sum(nil))
	request.Header.Set("Wps-Docs-Date", dataString)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Wps-Docs-Authorization", authorization)
}
