/*
Package third_party_service Third-party services(zh-cn:第三方服务)
*/
package third_party_service

import (
	"github.com/sevndata/blockchain_box/service_info"
	"github.com/sevndata/blockchain_box/service_info/domain_info"
)

var (
	// SupportedThirdPartyServices 已支持的第三方服务Map
	SupportedThirdPartyServices = map[string]ThirdPartyServiceFlag{
		ThirdPartyServiceFlagRedis.ServiceName(): ThirdPartyServiceFlagRedis,
		ThirdPartyServiceFlagMysql.ServiceName(): ThirdPartyServiceFlagMysql,
		ThirdPartyServiceFlagNginx.ServiceName(): ThirdPartyServiceFlagNginx,
	}
)

// ThirdPartyServiceFlag 第三方服务标识
type ThirdPartyServiceFlag int

// server常量
const (
	ThirdPartyServiceFlagNone ThirdPartyServiceFlag = iota
	ThirdPartyServiceFlagRedis
	ThirdPartyServiceFlagMysql
	ThirdPartyServiceFlagNginx
)

// GetThirdPartyServiceFlagFromStr 根据server name 称获取 第三方service Flag
// str 服务名
func GetThirdPartyServiceFlagFromStr(str string) ThirdPartyServiceFlag {
	switch str {
	case "redis-server":
		return ThirdPartyServiceFlagRedis
	case "mysql":
		return ThirdPartyServiceFlagMysql
	case "nginx":
		return ThirdPartyServiceFlagNginx
	default:
		return ThirdPartyServiceFlagNone
	}
}

// ServiceName 服务名称
func (tpsf ThirdPartyServiceFlag) ServiceName() string {
	switch tpsf {
	case ThirdPartyServiceFlagRedis:
		return "redis-server"
	case ThirdPartyServiceFlagMysql:
		return "mysql"
	case ThirdPartyServiceFlagNginx:
		return "nginx"
	default:
		return "none"
	}
}

// IsCSupported 检查服务是否受工具支持
// serviceName 服务名
func IsCSupported(serviceName string) bool {
	_, exists := SupportedThirdPartyServices[serviceName]
	return exists
}

type ServiceInfoThirdParty struct {
	ServiceInfoBase       service_info.ServiceInfoBase       // 基础信息
	ThirdPartyServiceFlag ThirdPartyServiceFlag              // 第三方服务标识
	DomainInfo            map[string]*domain_info.DomainInfo // 域名信息
}
