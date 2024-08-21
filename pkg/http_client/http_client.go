package http_client

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"wpsApi/pkg/log"
	"wpsApi/pkg/sign"
	"wpsApi/pkg/utils"
)

func GetByte(operationID string, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Info(operationID, "api request body len :", len(all))
	return all, nil
}
func Get(operationID string, url string, req interface{}, resp interface{}, sign sign.Sign) error {
	body := strings.NewReader("")
	if req != nil {
		jsonStr, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(jsonStr))
	}
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return err
	}
	if sign != nil {
		sign.Sign(request)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "api request err url is "+url, httpResponse.Status, string(result))
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+url+string(result))
	}
	log.Info(operationID, "api request success url is "+url, string(result))
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
func Post(operationID string, url string, req interface{}, resp interface{}, sign sign.Sign) error {
	body := strings.NewReader("")
	if req != nil {
		jsonStr, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(jsonStr))
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	if sign != nil {
		sign.Sign(request)
	}
	client := http.Client{Timeout: 5 * time.Second}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(httpResponse.Body)
	if httpResponse.StatusCode != 200 {
		log.Error(operationID, "api request err url is "+url, httpResponse.Status, string(result))
		return utils.Wrap(errors.New(httpResponse.Status), "status code failed "+url+string(result))
	}
	log.Info(operationID, "api request success url is "+url, string(result))
	err = utils.JsonStringToStruct(string(result), &resp)
	if err != nil {
		return err
	}
	return nil
}
