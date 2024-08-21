package api_req

type AvatarUploadReq struct {
	FileBase64  string `json:"file_base64"`
	NeedSession bool   `json:"need_session"`
}
