package domain_info

import "github.com/sevndata/blockchain_box/service_info/cert_info"

// DomainInfo 域名信息
type DomainInfo struct {
	FileName    string                 // 文件名称
	DomainName  string                 // 域名名称
	Certificate *cert_info.Certificate // 证书信息
}
