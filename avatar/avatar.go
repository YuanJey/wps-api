package avatar

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/YuanJey/wps-api/pkg/api_req"
	"github.com/YuanJey/wps-api/pkg/api_resp"
	"github.com/YuanJey/wps-api/pkg/consts"
	"github.com/YuanJey/wps-api/pkg/http_client"
	"github.com/YuanJey/wps-api/pkg/log"
	"github.com/YuanJey/wps-api/pkg/sign"
	"reflect"
)

type Avatar struct {
	addr      string
	sign      *sign.Sign
	companyId string
}

func NewAvatar(addr string, sign sign.Sign) *Avatar {
	return &Avatar{addr: addr, sign: &sign, companyId: "1"}
}
func (a *Avatar) CallFunc(funcName string, args ...interface{}) (resp []interface{}, err error) {
	value := reflect.ValueOf(a)
	method := value.MethodByName(funcName)
	if method.IsValid() {
		a := make([]reflect.Value, 0)
		for _, arg := range args {
			a = append(a, reflect.ValueOf(arg))
		}
		Resp := method.Call(a)
		for i := range Resp {
			if Resp[i].Kind() == reflect.Interface && !Resp[i].IsNil() {
				err = Resp[i].Interface().(error)
			} else if !Resp[i].IsNil() {
				resp = append(resp, Resp[i].Interface().(interface{}))
			}
		}
		return resp, err
	} else {
		log.Error("CallFunc method not found ", funcName, args)
		return nil, errors.New("CallFunc method not found")
	}
}

func (a *Avatar) UserAvatarUpload(operationID string, url string) (*api_resp.AvatarUploadResp, error) {
	getByte, err := http_client.GetByte(operationID, url)
	if err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(getByte)
	req := api_req.AvatarUploadReq{FileBase64: base64Str, NeedSession: true}
	resp := api_resp.AvatarUploadResp{}
	err = http_client.Post(operationID, fmt.Sprintf(a.addr+consts.AvatarUploadPath, a.companyId), req, &resp, *a.sign)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (a *Avatar) DeptAvatarUpload(operationID string, url string) (*api_resp.AvatarUploadResp, error) {
	getByte, err := http_client.GetByte(operationID, url)
	if err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(getByte)
	req := api_req.AvatarUploadReq{FileBase64: base64Str}
	resp := api_resp.AvatarUploadResp{}
	err = http_client.Post(operationID, fmt.Sprintf(a.addr+consts.AvatarUploadPath, a.companyId), req, &resp, *a.sign)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
