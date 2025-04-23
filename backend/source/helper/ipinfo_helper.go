package helper

import (
	"abnet_backend/source/logger"
	"fmt"
	"net"
	"path/filepath"
	"sync"

	"github.com/oschwald/maxminddb-golang"
)

// 按 IPinfo Lite 数据库模式定义结构
type LiteRecord struct {
	Network       string `maxminddb:"network"`
	Country       string `maxminddb:"country"`
	CountryCode   string `maxminddb:"country_code"`
	Continent     string `maxminddb:"continent"`
	ContinentCode string `maxminddb:"continent_code"`
	ASN           string `maxminddb:"asn"`
	ASName        string `maxminddb:"as_name"`
	ASDomain      string `maxminddb:"as_domain"`
}

// IPInfo 存储返回给客户端的IP信息结构体
type IPInfo struct {
	IP          string `json:"ip"`
	ASN         string `json:"asn"`
	ASNName     string `json:"asn_name"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
}

var (
	ipinfoDB   *maxminddb.Reader
	dbInitOnce sync.Once
	dbInitErr  error
)

// InitIPInfoDatabase 初始化 ipinfo 的 MMDB 数据库
func InitIPInfoDatabase(dbPath string) error {
	var err error

	// 打开 MaxMind 数据库
	ipinfoDB, err = maxminddb.Open(dbPath)
	if err != nil {
		return fmt.Errorf("无法打开 ipinfo 数据库: %v", err)
	}

	return nil
}

// GetIPInfoFromMMDB 从 ipinfo 数据库获取 IP 信息
func GetIPInfoFromMMDB(ip net.IP) (*IPInfo, error) {
	var record LiteRecord
	err := ipinfoDB.Lookup(ip, &record)
	if err != nil {
		return nil, fmt.Errorf("查询IP信息失败: %v", err)
	}

	ipInfo := &IPInfo{
		IP:          ip.String(),
		ASN:         record.ASN,
		ASNName:     record.ASName,
		CountryCode: record.CountryCode,
		Country:     record.Country,
	}

	return ipInfo, nil
}

// EnsureIPInfoDBInitialized 确保 ipinfo 数据库已初始化
func EnsureIPInfoDBInitialized() error {
	dbInitOnce.Do(func() {
		// 数据库文件路径
		dbPath := filepath.Join("data", "ipinfo.mmdb")

		dbInitErr = InitIPInfoDatabase(dbPath)
		if dbInitErr != nil {
			logger.Error("初始化 ipinfo 数据库失败: %v", dbInitErr)
		}
	})

	return dbInitErr
}
