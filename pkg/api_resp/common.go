package api_resp

type CommonResp struct {
	Code   int64  `json:"code"`
	Detail string `json:"detail"`
	Msg    string `json:"msg"`
}
type Failure struct {
	Code   string `json:"code"`
	DeptId string `json:"dept_id"`
	Result string `json:"resutl"`
}
type Success struct {
	Code   string `json:"code"`
	DeptId string `json:"dept_id"`
	Result string `json:"resutl"`
}
type Dept struct {
	Id              string `json:"id"`
	CompanyId       string `json:"company_id"`
	Weight          int    `json:"weight"`
	ThirdUnionId    string `json:"third_union_id"`
	ThirdPlatformId string `json:"third_platform_id"`
	Ctime           int    `json:"ctime"`
	Mtime           int    `json:"mtime"`
	Type            string `json:"type"`
	ParentId        string `json:"parent_id"`
	Name            string `json:"name"`
	Alias           string `json:"alias"`
	LeaderId        string `json:"leader_id"`
	CreatorId       string `json:"creator_id"`
	AbsPath         string `json:"abs_path"`
	IdPath          string `json:"id_path"`
	Source          string `json:"source"`
	Synced          bool   `json:"synced"`
}
type Member struct {
	AccountId         string        `json:"account_id"`
	Address           string        `json:"address"`
	CompanyId         string        `json:"company_id"`
	CompanyUid        string        `json:"company_uid"`
	Ctime             int           `json:"ctime"`
	DefDeptId         string        `json:"def_dept_id"`
	DeptNum           int           `json:"dept_num"`
	Email             string        `json:"email"`
	EmployeeId        string        `json:"employee_id"`
	EmployeeType      string        `json:"employee_type"`
	EmploymentStatus  string        `json:"employment_status"`
	Gender            string        `json:"gender"`
	IsInnerSuperAdmin bool          `json:"is_inner_super_admin"`
	Leader            string        `json:"leader"`
	LoginName         string        `json:"login_name"`
	MobilePhone       string        `json:"mobile_phone"`
	Mtime             int           `json:"mtime"`
	NickName          string        `json:"nick_name"`
	PreStatus         int           `json:"pre_status"`
	Role              string        `json:"role"`
	Source            string        `json:"source"`
	Status            string        `json:"status"`
	Telephone         string        `json:"telephone"`
	ThirdPlatformId   string        `json:"third_platform_id"`
	ThirdUnionId      string        `json:"third_union_id"`
	Title             string        `json:"title"`
	WorkPlace         string        `json:"work_place"`
	CustomFields      []CustomField `json:"custom_fields"`
	DeptList          []Dept        `json:"depts"`
	Synced            bool          `json:"synced"`
}

type CustomField struct {
	FieldId string `json:"field_id"`
	Text    string `json:"text"`
}
