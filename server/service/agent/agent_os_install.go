package agent

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"go.uber.org/zap"
)

type AgentService struct{}

type OsInstall struct {
	Sn    string `json:"sn"`
	Start bool   `json:"start"`
}

func (agentService *AgentService) AgentOsInstall(ipAddress string, sn string) (err error) {
	address := fmt.Sprintf("%s:%d", ipAddress, global.GVA_CONFIG.Agent.Port)

	var osInstall = &OsInstall{
		Sn:    sn,
		Start: true,
	}

	content, err := json.Marshal(osInstall)
	if err != nil {
		global.GVA_LOG.Error("任务内容序列化失败!", zap.Error(err))
		return err
	}

	var req = &request.ReqTLSListening{
		Group:   "OsInstall",
		Operate: "ReInstall",
		Content: string(content),
	}

	message, err := json.Marshal(req)
	if err != nil {
		global.GVA_LOG.Error("任务请求序列化失败!", zap.Error(err))
		return err
	}

	// 建立连接
	conn, err := tls.Dial("tcp", address, global.GVA_AGENT_KEY)
	//conn, err := net.Dial("tcp", address)

	if err != nil {
		global.GVA_LOG.Error("Agent连接建立失败!", zap.Error(err))
		return err
	}
	defer conn.Close()

	// 发送消息给服务器
	_, err = conn.Write([]byte(message))
	if err != nil {
		global.GVA_LOG.Error("OsInstall请求失败!", zap.Error(err))
		return err
	}

	// 读取服务器的响应
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		global.GVA_LOG.Error("节点响应失败!", zap.Error(err))
		return err
	}

	var resp response.Response
	err = json.Unmarshal(buf[:n], &resp)
	if err != nil {
		global.GVA_LOG.Error("节点响应解析失败!", zap.Error(err))
		return err
	}

	if resp.Code != 0 {
		err = fmt.Errorf(resp.Msg)
		global.GVA_LOG.Error("安装异常!", zap.Error(err))
		return err
	}

	return err
}
