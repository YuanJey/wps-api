package api_req

type GetThirdDeptListReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}
type GetDeptListReq struct {
	DeptIds []string `json:"dept_ids"`
}
type BatchCreateDeptReq struct {
	Detps []CreateDeptInfo `json:"detps"`
}
type CreateDeptInfo struct {
	Name            string `json:"name"`
	Source          string `json:"source"`
	ThirdPlatformId string `json:"third_platform_id"`
	ThirdUnionId    string `json:"third_union_id"`
	Weight          int    `json:"weight"`
}

type UpdateDeptReq struct {
	Name   string `json:"name"`
	Weight int    `json:"weight,omitempty"`
}

type ReThirdDeptNameReq struct {
	Name            string `json:"name"`
	ThirdPlatformId string `json:"third_platform_id"`
	ThirdUnionId    string `json:"third_union_id"`
}

type MoveDeptReq struct {
	DeptId     string `json:"dept_id"`
	ToParentId string `json:"to_parent_id"`
}
type BatchDeleteDeptReq struct {
	DeptIds []string `json:"dept_ids"`
}
type BatchDeleteThirdDeptReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}
type BatchBindThirdDeptReq struct {
	Itmes      []BatchBindThirdDeptItem `json:"itmes"`
	PlatformId string                   `json:"platform_id"`
}
type BatchBindThirdDeptItem struct {
	DeptId  string `json:"dept_id"`
	UnionId string `json:"union_id"`
}
