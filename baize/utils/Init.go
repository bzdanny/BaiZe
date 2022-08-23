package utils

import (
	"github.com/bzdanny/BaiZe/baize/utils/logger"
	"github.com/bzdanny/BaiZe/baize/utils/snowflake"
	"github.com/bzdanny/BaiZe/baize/utils/token"
)

func Init() {
	snowflake.Init()
	logger.Init()
	token.Init()
}
