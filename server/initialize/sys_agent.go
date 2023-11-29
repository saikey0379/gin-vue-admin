package initialize

import (
	"crypto/tls"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func Agent() {
	cert, err := tls.LoadX509KeyPair(global.GVA_CONFIG.Agent.ClientCrt, global.GVA_CONFIG.Agent.ClientKey)
	if err != nil {
		global.GVA_LOG.Error("密钥对读取失败!", zap.Error(err))
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过服务器端证书验证
		Certificates:       []tls.Certificate{cert},
	}

	global.GVA_AGENT_KEY = tlsConfig
}
