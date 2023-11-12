package backendgcf

import (
	"fmt"
	"testing"
	module "github.com/nugisorange/backendgcf/module"

)

// user
func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "testgis")
	var userdata User
	userdata.Username = "faisal"
	userdata.Role = "admin"
	userdata.Password = "sankuyges"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}

func TestGetAllUserFromUsername(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "testgis")
	ahay := module.GetUserFromUsername(mconn, "user", "faisal")
	fmt.Println(ahay)
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "testgis")
	ahay := module.GetAllUser(mconn, "user")
	fmt.Println(ahay)
}