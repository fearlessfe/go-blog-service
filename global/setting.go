package global

import (
	"github.com/fearlessfe/go-blog-service/pkg/logger"
	"github.com/fearlessfe/go-blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine *gorm.DB
	Logger *logger.Logger
)
