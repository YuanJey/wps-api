package api_router

const (
	//GetRootDept GET
	GetRootDept = "/org/dev/v1/companies/{company_id}/depts/root"
	//获取部门信息 GET
	GetDeptInfo = "/org/dev/v1/companies/{company_id}/depts/{dept_id}"
	//批量获取部门信息 POST
	BatchGetDeptInfo = "/org/dev/v1/companies/{company_id}/depts"
	//获取子部门列表 GET
	GetSubDeptInfo = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/depts"
	//创建子部门 POST
	CreateSubDept = "/org/dev/v1/companies/{company_id}/depts/{parent_id}"
	//批量创建子部门 POST
	BatchCreateSubDept = "/org/dev/v1/batch/companies/{company_id}/depts/{parent_id}"
	//修改部门 POST
	ReviseDept = "/org/dev/v1/update/companies/{company_id}/depts/{dept_id}"
	//修改部门名称 POST
	ReviseDeptName = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/rename"
	//移动部门 POST
	MoveDept = "/org/dev/v1/companies/{company_id}/depts/{dept_id}/rename"
	//删除部门 POST
	DeleteDept = "/org/dev/v1/delete/companies/{company_id}/depts/{dept_id}"
	//批量删除部门 POST
	BatchDeleteDept = "/org/dev/v1/delete/companies/{company_id}/depts"
	//根据三方union-id获取部门信息 GET
	GetDeptInfoByThirdID = "/org/dev/v1/companies/{company_id}/thirddepts"
	//根据三方union-id批量获取部门信息 POST
	BatchGetDeptInfoByThirdID = "/org/dev/v1/companies/{company_id}/thirddepts"
	//根据三方union-id修改部门名称 POST
	ReviseDeptNameByThirdID = "/org/dev/v1/companies/{company_id}/thirddepts/rename"
	//根据三方union-id删除部门 POST
	DeleteDeptByThirdID = "/org/dev/v1/delete/companies/{company_id}/thirddepts"
	//根据三方union-id批量删除部门 POST
	BatchDeleteDeptByThirdID = "/org/dev/v1/delete/batch/companies/{company_id}/thirddepts"
	//绑定第三方部门 POST
	BindThirdDept = "/org/dev/v1/batch/companies/{company_id}/bind/thirddepts"
	//解除绑定第三方部门 POST
	RelieveBindThirdDept = "/org/dev/v1/batch/companies/{company_id}/unbind/thirddepts"
)
