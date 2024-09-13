package api_resp

type GetDeptResp struct {
	CommonResp
	Data struct {
		Dept
		//Id              string `json:"id"`
		//CompanyId       string `json:"company_id"`
		//Weight          int    `json:"weight"`
		//ThirdUnionId    string `json:"third_union_id"`
		//ThirdPlatformId string `json:"third_platform_id"`
		//Ctime           int    `json:"ctime"`
		//Mtime           int    `json:"mtime"`
		//Type            string `json:"type"`
		//ParentId        string `json:"parent_id"`
		//Name            string `json:"name"`
		//Alias           string `json:"alias"`
		//LeaderId        string `json:"leader_id"`
		//CreatorId       string `json:"creator_id"`
		//AbsPath         string `json:"abs_path"`
		//IdPath          string `json:"id_path"`
		//Source          string `json:"source"`
	} `json:"data"`
}
type GetDeptListResp struct {
	CommonResp
	Data struct {
		Depts []Dept `json:"depts"`
		Total int    `json:"total"`
	} `json:"data"`
}

type CreateSubDeptResp struct {
	CommonResp
	Data struct {
		AbsPath         string `json:"abs_path"`
		Alias           string `json:"alias"`
		CompanyId       string `json:"company_id"`
		CreatorId       string `json:"creator_id"`
		Ctime           int64  `json:"ctime"`
		Id              string `json:"id"`
		IdPath          string `json:"id_path"`
		LeaderId        string `json:"leader_id"`
		Mtime           int64  `json:"mtime"`
		Name            string `json:"name"`
		ParentId        string `json:"parent_id"`
		Source          string `json:"source"`
		ThirdPlatformId string `json:"third_platform_id"`
		ThirdUnionId    string `json:"third_union_id"`
		Weight          int    `json:"weight"`
	} `json:"data"`
}
type BatchCreateDeptResp struct {
	CommonResp
	Data struct {
		Failure      []Failure `json:"failure"`
		FailureCount int       `json:"failure_count"`
		Success      []Success `json:"success"`
		SuccessCount int       `json:"success_count"`
	} `json:"data"`
}

type BatchDeleteDeptResp struct {
	CommonResp
	Data struct {
		Failure      []Failure `json:"failure"`
		FailureCount int       `json:"failure_count"`
		Success      []Success `json:"success"`
		SuccessCount int       `json:"success_count"`
	} `json:"data"`
}
