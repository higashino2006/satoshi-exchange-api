package test_lib

import (
	"encoding/json"
	"errors"
	"reflect"
	"se-api/src/internal/config"
	"se-api/src/internal/db"
)

func Init() error {
	err := config.Init()
	if err != nil {
		return err
	}
	err = db.Init()
	if err != nil {
		return err
	}
	return nil
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

	err = db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error
	if err != nil {
		return err
	}
	for _, item := range tables {
		if item == "schema_migrations" {
			continue
		}
		err = db.Exec("TRUNCATE TABLE " + item).Error
		if err != nil {
			return err
		}
	}
	err = db.Exec("SET FOREIGN_KEY_CHECKS = 1").Error
	if err != nil {
		return err
	}

	return nil
}

func IsJSONEqual(json1, json2 string) (bool, error) {
	var obj1, obj2 interface{}

	// JSON文字列をインターフェース{}にアンマーシャル
	if err := json.Unmarshal([]byte(json1), &obj1); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(json2), &obj2); err != nil {
		return false, err
	}

	// reflect.DeepEqualを使用して比較
	return reflect.DeepEqual(obj1, obj2), nil
}
