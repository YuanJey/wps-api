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

//	type GetDepartmentMembersResp struct {
//		CommonResp
//		Data struct {
//			DeptMembers []Member `json:"dept_members"`
//			Total       int      `json:"total"`
//		} `json:"data"`
//	}
type GetDepartmentMembersResp struct {
	Code int `json:"code"`
	Data struct {
		DeptMembers []DeptMembers `json:"dept_members"`
		Total       int           `json:"total"`
	} `json:"data"`
	Detail string `json:"detail"`
	Msg    string `json:"msg"`
}
type DeptMembers struct {
	Id          string `json:"id"`
	CompanyId   string `json:"company_id"`
	Ctime       int    `json:"ctime"`
	Mtime       int    `json:"mtime"`
	AccountInfo struct {
		Member
	} `json:"account_info"`
	DeptId     string `json:"dept_id"`
	AccountId  string `json:"account_id"`
	CompanyUid string `json:"company_uid"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	Sort       int    `json:"sort"`
	Weight     int    `json:"weight"`
	Source     string `json:"source"`
}
