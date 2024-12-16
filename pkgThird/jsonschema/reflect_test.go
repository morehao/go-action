package jsonschema

import (
	"fmt"
	"testing"
)

type User struct {
	Name      string  `json:"name" doc:"名字"`
	Age       int     `json:"age" form:"age" doc:"年龄"`
	MaxChild  Child   `json:"maxChild" form:"maxChild" doc:"最大的孩子"`
	ChildList []Child `json:"childList" form:"childList" doc:"所有的孩子"`
	School
}
type School struct {
	SchoolId uint `json:"schoolId" form:"schoolId" doc:"学校id"`
}

type Child struct {
	Name string `json:"name" form:"name" doc:"姓名"`
	Age  int    `json:"age" form:"age" doc:"年龄"`
}

func Test_fn(t *testing.T) {
	res := BuildJsonschema(&User{}, false)
	fmt.Println(res)
}

type AppVersionListRes struct {
	List []AppVersionListItem `json:"list" doc:"列表"`
}

type AppVersionListItem struct {
	Id            uint   `json:"id" doc:"数据自增id"`
	VersionNo     string `json:"versionNo" doc:"版本号"`
	BuildNo       string `json:"buildNo" doc:"构建号"`
	OsType        uint   `json:"osType" doc:"系统类型，1：ios，2：安卓"`
	IsForce       uint   `json:"isForce" doc:"是否强更，1：是，2：否"`
	IsSilent      uint   `json:"isSilent" doc:"是否静默更新，1：是，2：否"`
	VersionUrl    string `json:"versionUrl" doc:"版本地址"`
	VersionDesc   string `json:"versionDesc" doc:"版本描述"`
	CreatedBy     uint   `json:"createdBy" doc:"创建人id"`
	CreateByName  string `json:"createByName" doc:"创建人名称"`
	CreatedAt     string `json:"createdAt" doc:"创建时间"`
	UpdatedBy     uint   `json:"updatedBy" doc:"更新人id"`
	UpdatedByName string `json:"updatedByName" doc:"更新人名称"`
	UpdatedAt     string `json:"updatedAt" doc:"更新时间"`
}

func Test_Res(t *testing.T) {
	res := BuildJsonschema(&AppVersionListRes{}, true)
	fmt.Println(res)
}
