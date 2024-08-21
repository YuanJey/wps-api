package api_resp

type CreateCompaniesMembersResp struct {
	CommonResp
	Data struct {
		Member
	} `json:"data"`
}
type BatchGetCompanyMembersResp struct {
	CommonResp
	Data struct {
		CompanyMembers []Member `json:"company_members"`
		Total          int      `json:"total"`
	} `json:"data"`
}
type GetCompanyMembersResp struct {
	CommonResp
	Data struct {
		Member
	} `json:"data"`
}
type GetDepartmentMembersResp struct {
	CommonResp
	Data struct {
		DeptMembers []Member `json:"dept_members"`
		Total       int      `json:"total"`
	} `json:"data"`
}
