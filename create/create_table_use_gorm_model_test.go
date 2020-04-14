package create

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTableUseGormModel(t *testing.T) {
	t.Run("create table", func(t *testing.T) {
		db, err := gorm.Open("mysql", "docker:docker@/test_database?charset=utf8&parseTime=True&loc=Local")
		assert.NoError(t, err)

		err = db.Set("gorm:table_options", "engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin").
			AutoMigrate(&User{}, &Chat{}).Error
		assert.NoError(t, err)
	})
}
