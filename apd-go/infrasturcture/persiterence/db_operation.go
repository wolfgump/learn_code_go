package persiterence

import (
	"fmt"
	"zm-apd-go/domain/entity"

	"github.com/jinzhu/gorm"

	//init mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DbOperation db
type DbOperation struct {
	Db *gorm.DB
}

//NewDbOperation get a new db connection
func NewDbOperation(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*DbOperation, error) {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	DBURL := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbName)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	return &DbOperation{db}, nil
}

//Close connect
func (dbOpeartion *DbOperation) Close() error {
	return dbOpeartion.Db.Close()
}

//AutoMigrate auto create table schemal
func (dbOpeartion *DbOperation) AutoMigrate() error {
	return dbOpeartion.Db.AutoMigrate(&entity.Product{}).Error
}
