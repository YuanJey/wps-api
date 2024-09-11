package dept

import (
	"errors"
	"fmt"
	"github.com/YuanJey/wps-api/pkg/api_req"
	"github.com/YuanJey/wps-api/pkg/api_resp"
	"github.com/YuanJey/wps-api/pkg/consts"
	"github.com/YuanJey/wps-api/pkg/http_client"
	"github.com/YuanJey/wps-api/pkg/log"
	"github.com/YuanJey/wps-api/pkg/sign"
	"reflect"
	"strconv"
)

type ApiDept interface {
	// GetRootDept 获取根部门
	GetRootDept(operationID string) (*api_resp.GetDeptResp, error)
	// GetDeptInfoByThirdId 根据三方union-id获取部门信息
	GetDeptInfoByThirdId(operationID string, thirdId string) (*api_resp.GetDeptResp, error)
	// BatchGetThirdDeptList 根据三方union-id批量获取部门信息
	BatchGetThirdDeptList(operationID string, unionIds []string) (*api_resp.GetDeptListResp, error)
	// GetDeptInfo
	GetDeptInfo(operationID string, deptId string) (*api_resp.GetDeptResp, error)
	// GetDeptList 批量获取部门信息
	GetDeptList(operationID string, deptIds []int) (*api_resp.GetDeptListResp, error)
	// BatchCreateDept 批量创建子部门
	BatchCreateDept(operationID string, parentId string, deptList api_req.BatchCreateDeptReq) (*api_resp.BatchCreateDeptResp, error)
	// GetSubDeptList 获取子部门列表
	GetSubDeptList(operationID string, deptId, offset, limit string) (*api_resp.GetDeptListResp, error)
	UpdateDeptInfo(operationID string, deptId string, deptReq api_req.UpdateDeptReq) (*api_resp.CommonResp, error)
	ReThirdDeptNamePath(operationID string, req api_req.ReThirdDeptNameReq) (*api_resp.CommonResp, error)
	MoveDept(operationID string, req api_req.MoveDeptReq) (*api_resp.CommonResp, error)
	BatchDeleteDept(operationID string, req api_req.BatchDeleteDeptReq) (*api_resp.CommonResp, error)
	BatchDeleteThirdDept(operationID string, req api_req.BatchDeleteThirdDeptReq) (*api_resp.CommonResp, error)
	BatchBindThirdDept(operationID string, req api_req.BatchBindThirdDeptReq) (*api_resp.CommonResp, error)
}
type Dept struct {
	addr      string
	sign      *sign.Sign
	companyId string
}

func (d *Dept) GetDeptInfo(operationID string, deptId string) (*api_resp.GetDeptResp, error) {
	getDeptResp := api_resp.GetDeptResp{}
	err := http_client.Get(operationID, fmt.Sprintf(d.addr+consts.GetDeptInfoPath, d.companyId, deptId), nil, &getDeptResp, *d.sign)
	if err != nil {
		log.Error(operationID, "GetDeptInfo err ", err.Error())
		return nil, err
	}
	return &getDeptResp, nil
}

func NewDept(sign sign.Sign, addr string) *Dept {
	return &Dept{sign: &sign, addr: addr, companyId: "1"}
}

func (d *Dept) CallFunc(funcName string, args ...interface{}) (resp []interface{}, err error) {
	value := reflect.ValueOf(d)
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
func (d *Dept) A(a int, b int) {
	fmt.Println(a + b)
}
func (d *Dept) GetRootDept(operationID string) (*api_resp.GetDeptResp, error) {
	getDeptResp := api_resp.GetDeptResp{}
	err := http_client.Get(operationID, fmt.Sprintf(d.addr+consts.GetRootDeptPath, d.companyId), nil, &getDeptResp, *d.sign)
	if err != nil {
		log.Error(operationID, "GetRootDept err ", err.Error())
		return &getDeptResp, err
	}
	return &getDeptResp, nil
}

func (d *Dept) GetDeptInfoByThirdId(operationID string, thirdId string) (*api_resp.GetDeptResp, error) {
	getDeptResp := api_resp.GetDeptResp{}
	err := http_client.Get(operationID, fmt.Sprintf(d.addr+consts.GetDeptInfoByThirdId, d.companyId, thirdId), nil, &getDeptResp, *d.sign)
	if err != nil {
		log.Error(operationID, "GetRootDept err ", err.Error())
		return &getDeptResp, err
	}
	return &getDeptResp, nil
}
func (d *Dept) BatchGetThirdDeptList(operationID string, unionIds []string) (*api_resp.GetDeptListResp, error) {
	thirdDeptListReq := api_req.GetThirdDeptListReq{PlatformId: "1", UnionIds: unionIds}
	getDeptListResp := api_resp.GetDeptListResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.BatchGetThirdDeptListPath, d.companyId), thirdDeptListReq, &getDeptListResp, *d.sign)
	if err != nil {
		log.Error(operationID, "BatchGetThirdDeptList err ", err.Error())
		return nil, err
	}
	return &getDeptListResp, nil
}
func (d *Dept) GetDeptList(operationID string, deptIds []int) (*api_resp.GetDeptListResp, error) {
	GetDeptListReq := api_req.GetDeptListReq{DeptIds: deptIds}
	getDeptListResp := api_resp.GetDeptListResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.GetDeptListPath, d.companyId), GetDeptListReq, &getDeptListResp, *d.sign)
	if err != nil {
		log.Error(operationID, "GetDeptList err ", err.Error())
		return nil, err
	}
	return &getDeptListResp, nil
}

func (d *Dept) BatchCreateDept(operationID string, parentId string, deptList api_req.BatchCreateDeptReq) (*api_resp.BatchCreateDeptResp, error) {
	BatchCreateDeptResp := api_resp.BatchCreateDeptResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.BatchCreateDeptPath, d.companyId, parentId), deptList, &BatchCreateDeptResp, *d.sign)
	if err != nil {
		log.Error(operationID, "BatchCreateDept err ", err.Error())
		return nil, err
	}
	return &BatchCreateDeptResp, nil
}

func (d *Dept) GetSubDeptList(operationID string, deptId, offset, limit string) (*api_resp.GetDeptListResp, error) {
	getDeptListResp := api_resp.GetDeptListResp{}
	err := http_client.Get(operationID, fmt.Sprintf(d.addr+consts.GetSubDeptListPath, d.companyId, deptId, offset, limit), nil, &getDeptListResp, *d.sign)
	if err != nil {
		log.Error(operationID, "GetSubDeptList err ", err.Error())
		return nil, err
	}
	return &getDeptListResp, nil
}
func (d *Dept) UpdateDeptInfo(operationID string, deptId string, deptReq api_req.UpdateDeptReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.UpdateDeptPath, d.companyId, deptId), deptReq, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "UpdateDeptInfo err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}
func (d *Dept) ReThirdDeptNamePath(operationID string, req api_req.ReThirdDeptNameReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.ReThirdDeptNamePath, d.companyId), req, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "ReThirdDeptNamePath err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}

func (d *Dept) MoveDept(operationID string, req api_req.MoveDeptReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.MoveDeptPath, d.companyId), req, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "MoveDept err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}

func (d *Dept) BatchDeleteDept(operationID string, req api_req.BatchDeleteDeptReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.BatchDeleteDeptPath, d.companyId), req, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "BatchDeleteDept err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}

func (d *Dept) BatchDeleteThirdDept(operationID string, req api_req.BatchDeleteThirdDeptReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.BatchDeleteThirdDeptPath, d.companyId), req, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "BatchDeleteThirdDept err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}

type AllWpsDept struct {
	RootDept    api_resp.Dept
	WidDeptList map[string]*api_resp.Dept
	TidDeptList map[string]*api_resp.Dept
}

func (d *Dept) GetAllWpsDept(operationID string) (*AllWpsDept, error) {
	m := make(map[string]*api_resp.Dept)
	wpsDept := AllWpsDept{}
	rootDept, err := d.GetRootDept(operationID)
	if err != nil {
		return nil, err
	}
	wpsDept.RootDept = rootDept.Data.Dept
	department, err := d.getAllDepartment(operationID, []api_resp.Dept{rootDept.Data.Dept})
	if err != nil {
		return nil, err
	}
	tm := make(map[string]*api_resp.Dept)
	for i := range department {
		m[department[i].Id] = &department[i]
		tm[department[i].ThirdUnionId] = &department[i]
	}
	wpsDept.WidDeptList = m
	wpsDept.TidDeptList = tm
	return &wpsDept, nil
}
func (d *Dept) getAllDepartment(operationID string, depts []api_resp.Dept) ([]api_resp.Dept, error) {
	var alldeptList []api_resp.Dept
	for _, dep := range depts {
		subDeptList, err := d.getSubDept(operationID, dep.Id)
		if err != nil {
			log.Error(operationID, "getSubDept err ", err.Error())
			continue
		}
		alldeptList = append(alldeptList, subDeptList...)
		department, err := d.getAllDepartment(operationID, subDeptList)
		if err != nil {
			log.Error(operationID, "getAllDepartment err ", err.Error())
			continue
		}
		alldeptList = append(alldeptList, department...)
	}
	return alldeptList, nil
}

// BatchBindThirdDept
func (d *Dept) BatchBindThirdDept(operationID string, req api_req.BatchBindThirdDeptReq) (*api_resp.CommonResp, error) {
	commonResp := api_resp.CommonResp{}
	err := http_client.Post(operationID, fmt.Sprintf(d.addr+consts.BatchBindThirdDeptPath, d.companyId), req, &commonResp, *d.sign)
	if err != nil {
		log.Error(operationID, "BatchBindThirdDept err ", err.Error())
		return nil, err
	}
	return &commonResp, nil
}
func (d *Dept) getSubDept(operationID string, deptId string) ([]api_resp.Dept, error) {
	var deptList []api_resp.Dept
	page := 1
	size := 1000
	for {
		offset := (page - 1) * size
		list, err := d.GetSubDeptList(operationID, deptId, strconv.Itoa(offset), strconv.Itoa(size))
		if err != nil {
			return deptList, err
		}
		page++
		deptList = append(deptList, list.Data.Depts...)
		if len(list.Data.Depts) < size {
			return deptList, nil
		}
	}
}
