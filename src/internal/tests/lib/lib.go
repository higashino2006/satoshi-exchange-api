package test_lib

import (
	"errors"
	"se-api/src/internal/config"
	"se-api/src/internal/db"
	"se-api/src/internal/lib/common"
	"strings"
)

func JoinWithTestBackendUrl(subPath string) string {
	return common.JoinPaths(config.AppConfig.TEST_BACKEND_URL, subPath)
}

func TruncateAllTables() error {
	if !config.AppConfig.TEST_MODE {
		return errors.New("it only runs on test mode")
	}

	db := db.GetDB()
	var tables []string

	err := db.Raw("SHOW TABLES").Scan(&tables).Error
	if err != nil {
		return err
	}
	err = db.Exec("TRUNCATE TABLE " + strings.Join(tables, ",")).Error
	if err != nil {
		return err
	}
	return nil
}
