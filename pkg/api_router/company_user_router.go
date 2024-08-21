package api_router

const (
	//删除企业成员 POST
	DeleteCompanyUser = "/org/dev/v1/delete/batch/companies/{company_id}/members"
	//启用企业成员 POST
	EnableCompanyUser = "/org/dev/v1/batch/companies/{company_id}/members/enable"
	//根据第三方union-id删除企业成员 POST
	DeleteCompanyUserByThirdID = "/org/dev/v1/delete/batch/companies/{company_id}/thirdusers"
	//根据第三方union-id启用企业成员 POST
	EnableCompanyUserByThirdID = "/org/dev/v1/batch/companies/{company_id}/thirdusers/enable"
	//根据第三方union-id禁用企业成员 POST
	DisableCompanyUserByThirdID = "/org/dev/v1/batch/companies/{company_id}/thirdusers/disable"
	//禁用企业成员 POST POST
	DisableCompanyUser = "/org/dev/v1/batch/companies/{company_id}/members/disable"
	//绑定第三方用户 POST
	BindThirdUser = "/org/dev/v1/batch/companies/{company_id}/bind/thirdusers"
	//解除绑定第三方用户 POST
	RelieveBindThirdUser = "/org/dev/v1/batch/companies/{company_id}/unbind/thirdusers"
	//修改企业成员信息 POST
	ReviseCompanyUserInfo = "/org/dev/v1/companies/{company_id}/members/{account_id}"
	//创建企业成员 POST
	createCompanyUser = "/org/dev/v1/companies/{company_id}/members"
	//批量获取企业成员 POST
	BatchGetCompanyUserInfo = "/org/dev/v1/batch/companies/{company_id}/members"
	//根据第三方union-id修改企业成员信息 POST
	ReviseCompanyUserInfoByThirdID = "/org/dev/v1/companies/{company_id}/thirdusers"
	//根据第三方union-id批量获取企业成员 POST
	BatchGetCompanyUserInfoByThirdID = "/org/dev/v1/batch/companies/{company_id}/thirdusers"

	//根据第三方union-i	d获取企业成员 GET
	GetCompanyUserInfoByThirdID = "/org/dev/v1/companies/{company_id}/thirdusers"
	//获取企业成员 GET
	GetCompanyUserInfo = "/org/dev/v1/companies/{company_id}/members/{account_id}"
	//获取企业成员列表 GET
	GetCompanyUserList = "/org/dev/v1/companies/{company_id}/members"
)
