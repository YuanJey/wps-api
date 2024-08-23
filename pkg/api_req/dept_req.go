package api_req

type GetThirdDeptListReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}
type GetDeptListReq struct {
	DeptIds []int `json:"dept_ids"`
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
	Weight int    `json:"weight"`
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
