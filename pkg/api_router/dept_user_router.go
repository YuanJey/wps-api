package api_router

const (
	//修改部门成员的权重 POST
	ReviseDeptUserWeight = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/members/{account_id}/weight"
	//添加账户到某部门 POST
	AddUserToDept = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/members/{account_id}"
	//获取某部门成员 GET
	GetDeptUser = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/members/{account_id}"
	//获取账户所在的部门列表 GET
	GetUserDeptList = "/org/dev/v1/companies/{company_id}/members/{account_id}/depts"
	//获取部门成员列表 GET
	GetDeptUserList = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/members"
	//调整企业账户的归属部门 POST
	ReviseUserDept = "/org/dev/v1/companies/{company_id}/change_account_dept"
	//账户退出某部门 POST
	UserQuitDept = "/org/dev/v1/delete/companies/{company_id}/depts/{dept_id}/members/{account_id}"
)
