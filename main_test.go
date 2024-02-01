package backendgcf

import (
	"context"
	"fmt"
	"testing"

	gcf "github.com/nugisOrange/petback"
	"github.com/whatsauth/watoken"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("FaisalAsh", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Faisal",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGOSTRING", "gisChap04", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Faisal",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGOSTRING", "gisChap04", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGOSTRING", "gisChap04")
	pass, _ := gcf.HashPass("faisalganteng")
	data := RegisterStruct{
		Username: "faisal",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}