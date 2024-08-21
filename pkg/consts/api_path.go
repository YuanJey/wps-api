package consts

const (
	// CreateCompaniesMembersPath 创建企业成员
	CreateCompaniesMembersPath = "/org/dev/v1/companies/%s/members"

	//GetCompanyMembersPath 获取企业成员
	GetCompanyMembersPath = "/org/dev/v1/companies/%s/members/%s"
	// BatchGetCompanyMembersPath 批量获取企业成员信息
	BatchGetCompanyMembersPath = "/org/dev/v1/batch/companies/%s/members"
	// BatchGetCompanyMembersByThirdIdPath 根据第三方id批量获取企业成员
	BatchGetCompanyMembersByThirdIdPath = "/org/dev/v1/batch/companies/%s/thirdusers"

	// BatchDisableCompanyMembersPath 批量禁用企业成员
	BatchDisableCompanyMembersPath = "/org/dev/v1/batch/companies/%s/members/disable"
	// BatchEnableCompanyMembersPath 批量启用企业成员
	BatchEnableCompanyMembersPath = "/org/dev/v1/batch/companies/%s/members/enable"

	// UpdateThirdMemberInfoPath 根据第三方union-id修改企业成员信息
	UpdateThirdMemberInfoPath = "/org/dev/v1/companies/%s/thirdusers/rename"
	// UpdateMemberInfoPath 修改企业成员信息
	UpdateMemberInfoPath = "/org/dev/v1/companies/%s/members/%s"

	// ChangeCompanyMembersDeptPath 调整企业账户的归属部门
	ChangeCompanyMembersDeptPath = "/org/dev/v1/companies/%s/change_account_dept"
	GetDepartmentMembersPath     = "/org/dev/v1/companies/%s/depts/%s/members"
)

const (
	// GetRootDeptPath 获取根部门
	GetRootDeptPath = "/org/dev/v1/companies/%s/depts/root"
	// GetDeptInfoByThirdId 根据三方union-id获取部门信息
	GetDeptInfoByThirdId = "/org/dev/v1/companies/%s/thirddepts?third_platform_id=1&third_union_id=&s"
	// BatchGetThirdDeptListPath 根据三方union-id批量获取部门信息
	BatchGetThirdDeptListPath = "/org/dev/v1/companies/%s/thirddepts"
	// GetDeptListPath 批量获取部门信息
	GetDeptListPath = "/org/dev/v1/companies/%s/depts"
	// GetSubDeptListPath 获取子部门列表
	GetSubDeptListPath = "/org/dev/v1/companies/%s/depts/%s/depts?offset=%s&limit=%s"

	// BatchCreateDeptPath 批量创建子部门
	BatchCreateDeptPath = "/org/dev/v1/batch/companies/%s/depts/%s"

	// UpdateDeptPath 修改部门信息
	UpdateDeptPath = "/org/dev/v1/update/companies/%s/depts/%s"
	// ReThirdDeptNamePath 根据三方union-id修改部门名称
	ReThirdDeptNamePath = "/org/dev/v1/companies/%s/thirddepts/rename"
	// MoveDeptPath 移动部门
	MoveDeptPath = "/org/dev/v1/companies/%s/depts/move"

	// BatchDeleteDeptPath 批量删除部门
	BatchDeleteDeptPath = "/org/dev/v1/delete/companies/%s/depts"
	// BatchDeleteThirdDeptPath 根据三方union-id批量删除部门
	BatchDeleteThirdDeptPath = "/org/dev/v1/delete/batch/companies/%s/thirddepts"
)
const (
	AvatarUploadPath = "/avatar/dev/v1/companies/%s/custom_avatar"
)
