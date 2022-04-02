package config

import (
	"os"
	"strings"
	"sync"

	"github.com/Yujiman/e_commerce/goods/order/aggregatorOrder/internal/utils"
)

var (
	onceIP    sync.Once
	ipChecker *utils.IPChecker
)

func GetIPChecker() *utils.IPChecker {
	onceIP.Do(func() {
		ipChecker = &utils.IPChecker{
			Whitelist: strings.Split(os.Getenv("ALLOWED_IP"), ";"),
		}
	})

	return ipChecker
}
