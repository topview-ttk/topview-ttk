package db

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var syncLock sync.Mutex
var db *gorm.DB

type ttkDSConfig struct {
	DSN                string `yaml:"datasource.ttk-user.dsn"`
	MaxIdleConnections int    `yaml:"datasource.ttk-user.gorm.max-idle-connections"`
	MaxOpenConnections int    `yaml:"datasource.ttk-user.gorm.max-open-connections"`
	ConnMaxLifetime    int    `yaml:"datasource.ttk-user.gorm.conn-max-lifetime"`
}

func setupDB(config *ttkDSConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ttk_",
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	return db, nil
}

func init() {
	dir, _ := os.Getwd()
	yamlPath := filepath.Join(dir, "configs/db-cfg.yaml")
	file, err := os.ReadFile(yamlPath)
	if err != nil {
		logrus.Errorf("Read db-cfg fail: %s", err)
		panic(err)
		return
	}

	config := &ttkDSConfig{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logrus.Errorf("Unmarshal db fail: %s", err)
		panic(err)
		return
	}

	syncLock.Lock()
	defer syncLock.Unlock()

	db, err = setupDB(config)
	if err != nil {
		logrus.Errorf("Connect to db fail: %s", err)
		panic(err)
		return
	}

}

func GetDb() *gorm.DB {
	return db
}
