package api_req

import "github.com/YuanJey/wps-api/pkg/api_resp"

type CreateCompaniesMembersReq struct {
	Account          string                 `json:"account"`
	Avatar           string                 `json:"avatar"`
	DefDeptId        int                    `json:"def_dept_id"`
	DeptIds          []DeptIds              `json:"dept_ids"`
	Email            string                 `json:"email"`
	EmployeeId       string                 `json:"employee_id"`
	EmploymentStatus string                 `json:"employment_status"`
	EmploymentType   string                 `json:"employment_type"`
	Gender           string                 `json:"gender"`
	MobilePhone      string                 `json:"mobile_phone"`
	NickName         string                 `json:"nick_name"`
	Password         string                 `json:"password"`
	Source           string                 `json:"source"`
	Telephone        string                 `json:"telephone"`
	ThirdPlatformId  string                 `json:"third_platform_id"`
	ThirdUnionId     string                 `json:"third_union_id"`
	Title            string                 `json:"title"`
	CustomFields     []api_resp.CustomField `json:"custom_fields"`
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
	Email            string                 `json:"email"`
	EmployeeId       string                 `json:"employee_id"`
	EmploymentStatus string                 `json:"employment_status"`
	EmploymentType   string                 `json:"employment_type"`
	Gender           string                 `json:"gender"`
	MobilePhone      string                 `json:"mobile_phone"`
	NickName         string                 `json:"nick_name"`
	Telephone        string                 `json:"telephone"`
	Title            string                 `json:"title"`
	CustomFields     []api_resp.CustomField `json:"custom_fields"`
	Avatar           string                 `json:"avatar"`
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
