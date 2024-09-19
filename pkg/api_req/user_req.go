package api_req

import "pl.ghgame.cn/gitea/yuanjie/wps-api/pkg/api_resp"

type CreateCompaniesMembersReq struct {
	Account          string                 `json:"account"`
	Avatar           string                 `json:"avatar,omitempty"`
	DefDeptId        string                 `json:"def_dept_id"`
	DeptIds          []DeptIds              `json:"dept_ids"`
	Email            string                 `json:"email,omitempty"`
	EmployeeId       string                 `json:"employee_id,omitempty"`
	EmploymentStatus string                 `json:"employment_status,omitempty"`
	EmploymentType   string                 `json:"employment_type,omitempty"`
	Gender           string                 `json:"gender,omitempty"`
	MobilePhone      string                 `json:"mobile_phone,omitempty"`
	NickName         string                 `json:"nick_name"`
	Password         string                 `json:"password,omitempty"`
	Source           string                 `json:"source"`
	Telephone        string                 `json:"telephone,omitempty"`
	ThirdPlatformId  string                 `json:"third_platform_id"`
	ThirdUnionId     string                 `json:"third_union_id"`
	Title            string                 `json:"title,omitempty"`
	CustomFields     []api_resp.CustomField `json:"custom_fields,omitempty"`
}
type DeptIds struct {
	DeptId string `json:"dept_id"`
	Weight int    `json:"weight"`
}
type BatchGetCompanyMembersReq struct {
	AccountIds []string `json:"account_ids"`
}
type BatchDisableCompanyMembersReq struct {
	AccountIds []string `json:"account_ids"`
}
type BatchEnableCompanyMembersReq struct {
	AccountIds []string `json:"account_ids"`
}
type BatchGetCompanyMembersByThirdIdReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}

type UpdateMemberInfoReq struct {
	Email            string                 `json:"email,omitempty"`
	EmployeeId       string                 `json:"employee_id,omitempty"`
	EmploymentStatus string                 `json:"employment_status,omitempty"`
	EmploymentType   string                 `json:"employment_type,omitempty"`
	Gender           string                 `json:"gender,omitempty"`
	MobilePhone      string                 `json:"mobile_phone,omitempty"`
	NickName         string                 `json:"nick_name"`
	Telephone        string                 `json:"telephone,omitempty"`
	Title            string                 `json:"title,omitempty"`
	CustomFields     []api_resp.CustomField `json:"custom_fields,omitempty"`
	Avatar           string                 `json:"avatar,omitempty"`
}
type ChangeCompanyMembersDeptReq struct {
	AccountId  string    `json:"account_id"`
	DefDeptId  string    `json:"def_dept_id"`
	NewDeptIds []NewDept `json:"new_dept_ids"`
	OldDeptIds []string  `json:"old_dept_ids"`
}
type NewDept struct {
	DeptId string `json:"dept_id"`
	Weight int    `json:"weight"`
}

type DeleteBatchCompanyMembersReq struct {
	AccountIds []string `json:"account_ids"`
}
type BatchDisableThirdMembersReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}

type BatchEnableThirdMembersReq struct {
	PlatformId string   `json:"platform_id"`
	UnionIds   []string `json:"union_ids"`
}
type ChangeMemberDeptWeightReq struct {
	Weight int `json:"weight"`
}
type BatchBindThirdMemberReq struct {
	Itmes      []BatchBindThirdMemberItem `json:"itmes"`
	PlatformId string                     `json:"platform_id"`
}
type BatchBindThirdMemberItem struct {
	AccountId string `json:"account_id"`
	UnionId   string `json:"union_id"`
}
