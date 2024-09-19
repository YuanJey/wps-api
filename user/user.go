package user

import (
	"errors"
	"fmt"
	"github.com/YuanJey/wps-api/dept"
	"github.com/YuanJey/wps-api/pkg/api_req"
	"github.com/YuanJey/wps-api/pkg/api_resp"
	"github.com/YuanJey/wps-api/pkg/config"
	"github.com/YuanJey/wps-api/pkg/consts"
	"github.com/YuanJey/wps-api/pkg/http_client"
	"github.com/YuanJey/wps-api/pkg/log"
	"github.com/YuanJey/wps-api/pkg/sign"
	"github.com/YuanJey/wps-api/pkg/utils"
	"reflect"
	"strconv"
)

type ApiUser interface {
	// CreateCompaniesMembers 创建企业成员
	CreateCompaniesMembers(operationID string, req api_req.CreateCompaniesMembersReq) (*api_resp.CreateCompaniesMembersResp, error)

	//GetCompanyMembersPath
	GetCompanyMembers(operationID string, accountId string) (*api_resp.GetCompanyMembersResp, error)
	// BatchGetCompanyMembers 批量获取企业成员
	BatchGetCompanyMembers(operationID string, accountList []string) (*api_resp.BatchGetCompanyMembersResp, error)
	// BatchGetCompanyMembersByThirdId 根据第三方id批量获取企业成员
	BatchGetCompanyMembersByThirdId(operationID string, req api_req.BatchGetCompanyMembersByThirdIdReq) (*api_resp.BatchGetCompanyMembersResp, error)

	// BatchDisableCompanyMembers 批量禁用企业成员
	BatchDisableCompanyMembers(operationID string, accountList []string) (*api_resp.CommonResp, error)
	// BatchEnableCompanyMembers 批量启用企业成员
	BatchEnableCompanyMembers(operationID string, accountList []string) (*api_resp.CommonResp, error)

	//UpdateThirdMemberInfo 根据第三方union-id修改企业成员信息
	UpdateThirdMemberInfo(operationID, thirdUnionId string, req api_req.UpdateMemberInfoReq) (*api_resp.CommonResp, error)
	// UpdateMemberInfo 修改企业成员信息
	UpdateMemberInfo(operationID string, accountId string, req api_req.UpdateMemberInfoReq) (*api_resp.CommonResp, error)
	// ChangeCompanyMembersDept 调整企业账户的归属部门
	ChangeCompanyMembersDept(operationID string, req api_req.ChangeCompanyMembersDeptReq) (*api_resp.CommonResp, error)
	// GetDepartmentMembersPath 获取部门成员列表
	GetDepartmentMembers(operationID string, deptId, offset, limit string) (*api_resp.GetDepartmentMembersResp, error)
	// BatchDeleteCompanyMembersPath 批量删除用户
	BatchDeleteCompanyMembers(operationID string, accounts []string) (*api_resp.CommonResp, error)
	//BatchDisableThirdMembers
	BatchDisableThirdMembers(operationID string, req api_req.BatchDisableThirdMembersReq) (*api_resp.CommonResp, error)
	//BatchEnableThirdMembers
	BatchEnableThirdMembers(operationID string, req api_req.BatchEnableThirdMembersReq) (*api_resp.CommonResp, error)
	//ChangeMemberDeptWeightPath
	ChangeMemberDeptWeight(operationID, deptId, accountId string, req api_req.ChangeMemberDeptWeightReq) (*api_resp.CommonResp, error)
	GetCompanyMembersByStatus(operationID string, status, offset, limit string) (*api_resp.BatchGetCompanyMembersResp, error)
	//BatchBindThirdMemberPath
	BatchBindThirdMember(operationID string, req api_req.BatchBindThirdMemberReq) (*api_resp.CommonResp, error)
}
type User struct {
	addr      string
	sign      *sign.Sign
	companyId string
}

func (u *User) BatchBindThirdMember(operationID string, req api_req.BatchBindThirdMemberReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchBindThirdMemberPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchBindThirdMemberPath err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) GetCompanyMembersByStatus(operationID string, status, offset, limit string) (*api_resp.BatchGetCompanyMembersResp, error) {
	resp := api_resp.BatchGetCompanyMembersResp{}
	err := http_client.Get(operationID, fmt.Sprintf(u.addr+consts.GetCompanyMembersByStatusPath, u.companyId, status, offset, limit), nil, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "GetCompanyMembersByStatusPath err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) ChangeMemberDeptWeight(operationID, deptId, accountId string, req api_req.ChangeMemberDeptWeightReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.ChangeMemberDeptWeightPath, u.companyId, deptId, accountId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "ChangeMemberDeptWeightPath err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) GetDepartmentMembers(operationID string, deptId, offset, limit string) (*api_resp.GetDepartmentMembersResp, error) {
	getMembersResp := api_resp.GetDepartmentMembersResp{}
	err := http_client.Get(operationID, fmt.Sprintf(u.addr+consts.GetDepartmentMembersPath, u.companyId, deptId, offset, limit), nil, &getMembersResp, *u.sign)
	if err != nil {
		log.Error(operationID, "GetDepartmentMembersPath err ", err.Error())
		return nil, err
	}
	return &getMembersResp, nil
}

func NewUser(sign sign.Sign, addr string) *User {
	return &User{sign: &sign, addr: addr, companyId: "1"}
}

func (u *User) CallFunc(funcName string, args ...interface{}) (resp []interface{}, err error) {
	value := reflect.ValueOf(u)
	method := value.MethodByName(funcName)
	if method.IsValid() {
		a := make([]reflect.Value, 0)
		for _, arg := range args {
			a = append(a, reflect.ValueOf(arg))
		}
		Resp := method.Call(a)
		for i := range Resp {
			if Resp[i].Kind() == reflect.Interface && !Resp[i].IsNil() {
				err = Resp[i].Interface().(error)
			} else if !Resp[i].IsNil() {
				resp = append(resp, Resp[i].Interface().(interface{}))
			}
		}
		return resp, err
	} else {
		log.Error("CallFunc method not found ", funcName, args)
		return nil, errors.New("CallFunc method not found")
	}
}
func (u *User) ChangeCompanyMembersDept(operationID string, req api_req.ChangeCompanyMembersDeptReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.ChangeCompanyMembersDeptPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "ChangeCompanyMembersDept err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) CreateCompaniesMembers(operationID string, req api_req.CreateCompaniesMembersReq) (*api_resp.CreateCompaniesMembersResp, error) {
	resp := api_resp.CreateCompaniesMembersResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.CreateCompaniesMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "CreateCompaniesMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}
func (u *User) GetCompanyMembers(operationID string, accountId string) (*api_resp.GetCompanyMembersResp, error) {
	resp := api_resp.GetCompanyMembersResp{}
	err := http_client.Get(operationID, fmt.Sprintf(u.addr+consts.GetCompanyMembersPath, u.companyId, accountId), nil, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "GetCompanyMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}
func (u *User) BatchGetCompanyMembers(operationID string, accountList []string) (*api_resp.BatchGetCompanyMembersResp, error) {
	req := api_req.BatchGetCompanyMembersReq{AccountIds: accountList}
	resp := api_resp.BatchGetCompanyMembersResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchGetCompanyMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchGetCompanyMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) BatchGetCompanyMembersByThirdId(operationID string, req api_req.BatchGetCompanyMembersByThirdIdReq) (*api_resp.BatchGetCompanyMembersResp, error) {
	resp := api_resp.BatchGetCompanyMembersResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchGetCompanyMembersByThirdIdPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchGetCompanyMembersByThirdId err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) BatchDisableCompanyMembers(operationID string, accountList []string) (*api_resp.CommonResp, error) {
	req := api_req.BatchDisableCompanyMembersReq{AccountIds: accountList}
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchDisableCompanyMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchDisableCompanyMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) BatchEnableCompanyMembers(operationID string, accountList []string) (*api_resp.CommonResp, error) {
	req := api_req.BatchEnableCompanyMembersReq{AccountIds: accountList}
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchEnableCompanyMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchEnableCompanyMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) BatchDisableThirdMembers(operationID string, req api_req.BatchDisableThirdMembersReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchDisableThirdMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchDisableThirdMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}
func (u *User) BatchEnableThirdMembers(operationID string, req api_req.BatchEnableThirdMembersReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.BatchEnableThirdMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchEnableThirdMembers err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) UpdateThirdMemberInfo(operationID, thirdUnionId string, req api_req.UpdateMemberInfoReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.UpdateThirdMemberInfoPath, u.companyId, config.Config.WPS.PlatformId, thirdUnionId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "UpdateThirdMemberInfo err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

func (u *User) UpdateMemberInfo(operationID string, accountId string, req api_req.UpdateMemberInfoReq) (*api_resp.CommonResp, error) {
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.UpdateMemberInfoPath, u.companyId, accountId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "UpdateMemberInfo err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

// BatchDeleteCompanyMembersPath
func (u *User) BatchDeleteCompanyMembers(operationID string, accounts []string) (*api_resp.CommonResp, error) {
	req := api_req.DeleteBatchCompanyMembersReq{AccountIds: accounts}
	resp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(u.addr+consts.DeleteBatchCompanyMembersPath, u.companyId), req, &resp, *u.sign)
	if err != nil {
		log.Error(operationID, "BatchDeleteCompanyMembersPath err ", err.Error())
		return nil, err
	}
	return &resp, nil
}

type AllWpsUser struct {
	WidUserList map[string]*api_resp.Member
	TidUserList map[string]*api_resp.Member
}

// 获取全部企业用户
func (u *User) GetAllUser(operationID string, allDept *dept.AllWpsDept) (*AllWpsUser, error) {
	var allUser []api_resp.Member
	for id := range allDept.WidDeptList {
		users, err := u.getDeptUsers(operationID, id)
		if err != nil {
			log.Error(operationID, "getDeptUsers err ", err.Error())
			return nil, err
		}
		allUser = append(allUser, users...)
	}
	//users, err := u.getDeptUsers(operationID, allDept.RootDept.Id)
	//if err != nil {
	//	return nil, err
	//}
	//allUser = append(allUser, users...)
	wu := make(map[string]*api_resp.Member)
	tu := make(map[string]*api_resp.Member)
	for i := range allUser {
		log.Info(operationID, "user ", allUser[i])
		tu[allUser[i].ThirdUnionId] = &allUser[i]
		wu[allUser[i].AccountId] = &allUser[i]
	}
	mapUser := AllWpsUser{WidUserList: wu, TidUserList: tu}
	return &mapUser, nil
}
func (u *User) getDeptUsers(operationID string, deptId string) ([]api_resp.Member, error) {
	var deptList []api_resp.Member
	page := 1
	size := 1000
	for {
		offset := (page - 1) * size
		list, err := u.GetDepartmentMembers(operationID, deptId, strconv.Itoa(offset), strconv.Itoa(size))
		if err != nil {
			return deptList, err
		}
		page++
		//var temp []api_resp.Member
		for i := range list.Data.DeptMembers {
			member := api_resp.Member{}
			if err := utils.CopyStructFields(&member, list.Data.DeptMembers[i].AccountInfo); err != nil {
				log.Error(operationID, "CopyStructFields failed ", err.Error())
				continue
			}
			//temp = append(temp, member)
			deptList = append(deptList, member)
		}
		//deptList = append(deptList, temp...)
		if len(list.Data.DeptMembers) < size {
			return deptList, nil
		}
	}
}
