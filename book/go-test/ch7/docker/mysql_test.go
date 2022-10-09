package docker

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/containerd/containerd/log"
)

func TestMysql(t *testing.T) {
	mysql, _ := mysqltestcontainer.Create("test")
	db := mysql.GetDb()
	err := db.Ping()
	if err != nil {
		log.L.Errorln(err.Error())
	}
}
