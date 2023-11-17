package backendgcf

import (
	"fmt"
	"testing"
	"github.com/whatsauth/watoken"
)
 
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("faisaltes", privateKey)
	fmt.Println(hasil, err)
}

func TestInsertUser(t *testing.T) {
	mconn := GetConnectionMongo("MONGOSTRING", "petasal")
	var userdata User
	userdata.Username = "shidiq"
	userdata.Password = "user123"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}