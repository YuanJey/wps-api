package api_resp

type AvatarUploadResp struct {
	CommonResp
	Data struct {
		Key  string `json:"key"`
		Size int64  `json:"size"`
	} `json:"data"`
}
