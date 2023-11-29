package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type RuleGroupSearch struct{
    monit.RuleGroup
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
