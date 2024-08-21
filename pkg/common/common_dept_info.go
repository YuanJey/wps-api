package common

import "fmt"

type Dept struct {
	Name            string `json:"name"`
	Source          string `json:"source"`
	ThirdPlatformId string `json:"third_platform_id"`
	ThirdUnionId    string `json:"third_union_id"`
	Weight          int    `json:"weight"`
	ParentId        string `json:"parent_id"`
	HandleNumber    int
	Children        []*Dept
}

func Build(nodes []*Dept) {
	temp := make(map[string]*Dept)
	for i := range nodes {
		temp[nodes[i].ThirdUnionId] = nodes[i]
	}

	for _, dept := range temp {
		if parent, ok := temp[dept.ParentId]; ok {
			parent.Children = append(parent.Children, dept)
		}
	}
	var delIds []string
	for id, dept := range temp {
		if _, ok := temp[dept.ParentId]; ok {
			delIds = append(delIds, id)
		}
	}
	for i := range delIds {
		delete(temp, delIds[i])
	}
	fmt.Println(temp)
}

func ToMap(nodes []*Dept) map[string]*Dept {
	temp := make(map[string]*Dept)
	for i := range nodes {
		temp[nodes[i].ThirdUnionId] = nodes[i]
	}
	return temp
}
func BuildMap(nodes []*Dept) map[string]*Dept {
	temp := make(map[string]*Dept)
	for i := range nodes {
		temp[nodes[i].ThirdUnionId] = nodes[i]
	}
	for _, dept := range temp {
		if parent, ok := temp[dept.ParentId]; ok {
			parent.Children = append(parent.Children, dept)
		}
	}
	return temp
}
func IncrementBuild(nodes []*Dept) map[string]*Dept {
	temp := make(map[string]*Dept)
	for i := range nodes {
		temp[nodes[i].ThirdUnionId] = nodes[i]
	}

	for _, dept := range temp {
		if parent, ok := temp[dept.ParentId]; ok {
			parent.Children = append(parent.Children, dept)
		}
	}
	return temp
}
