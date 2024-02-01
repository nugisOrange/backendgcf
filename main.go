package backendgcf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"context"
	"os"

	gcf "github.com/nugisOrange/petback"
	"github.com/petapedia/peda"
	"github.com/whatsauth/watoken"
)

func GCHandlerFunc(publickey, Mongostring, dbname, colname string, r *http.Request) string {
	resp := new(gcf.Credential)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			koneksyen := GetConnectionMongo(Mongostring, dbname)
			datageo := GetAllGeoData(koneksyen, colname)
			jsoncihuy, _ := json.Marshal(datageo)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(jsoncihuy)
		}
	}
	return ReturnStringStruct(resp)
}

func GCFPostCoordinate(Mongostring, Publickey, dbname, colname string, r *http.Request) string {
	req := new(Credents)
	conn := GetConnectionMongo(Mongostring, dbname)
	resp := new(LonLatProperties)
	err := json.NewDecoder(r.Body).Decode(&resp)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(Publickey))
		if !existing {
			req.Status = strconv.Itoa(http.StatusNotFound)
			req.Message = "Kamu kayaknya belum punya akun"
		} else {
			if err != nil {
				req.Status = strconv.Itoa(http.StatusNotFound)
				req.Message = "error parsing application/json: " + err.Error()
			} else {
				req.Status = strconv.Itoa(http.StatusOK)
				Ins := InsertDataLonlat(conn, colname,
					resp.Coordinates,
					resp.Name,
					resp.Volume,
					resp.Type)
				req.Message = fmt.Sprintf("%v:%v", "Berhasil Input data", Ins)
			}
		}
	}
	return ReturnStringStruct(req)
}

func GCFUpdateName(publickey, Mongostring, dbname string, r *http.Request) string {
	req := new(Credents)
	resp := new(LonLatProperties)
	err := json.NewDecoder(r.Body).Decode(&resp)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			req.Status = strconv.Itoa(http.StatusNotFound)
			req.Message = "Kamu kayaknya belum punya akun"
		} else {
			if err != nil {
				req.Status = strconv.Itoa(http.StatusNotFound)
				req.Message = "error parsing application/json: " + err.Error()
			} else {
				req.Status = strconv.Itoa(http.StatusOK)
				Ins := UpdateNameGeo(Mongostring, dbname, context.Background(),
					LonLatProperties{
						Type:   resp.Type,
						Name:   resp.Name,
						Volume: resp.Volume,
					})
				req.Message = fmt.Sprintf("%v:%v", "Berhasil Update data", Ins)
			}
		}
	}
	return ReturnStringStruct(req)
}

func GCFDeleteLon(publickey, Mongostring, dbname string, r *http.Request) string {
	req := new(Credents)
	resp := new(LonLatProperties)
	err := json.NewDecoder(r.Body).Decode(&resp)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = strconv.Itoa(http.StatusNotFound)
		req.Message = "Header Login Not Exist"
	} else {
		existing := IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			req.Status = strconv.Itoa(http.StatusNotFound)
			req.Message = "Kamu kayaknya belum punya akun"
		} else {
			if err != nil {
				req.Status = strconv.Itoa(http.StatusNotFound)
				req.Message = "error parsing application/json: " + err.Error()
			} else {
				req.Status = strconv.Itoa(http.StatusOK)
				Ins := DeleteDataGeo(Mongostring, dbname, context.Background(),
					LonLatProperties{
						Type:   resp.Type,
						Name:   resp.Name,
						Volume: resp.Volume,
					})
				req.Message = fmt.Sprintf("%v:%v", "Berhasil Hapus data", Ins)
			}
		}
	}
	return ReturnStringStruct(req)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}

func Register(Mongoenv, dbname string, r *http.Request) string {
	resp := new(gcf.Credential)
	userdata := new(RegisterStruct)
	resp.Status = false
	conn := GetConnectionMongo(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := gcf.HashPass(userdata.Password)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(conn, userdata.Username, hash)
		resp.Message = "Berhasil Input data"
	}
	response := gcf.ReturnStringStruct(resp)
	return response
}

func Login(Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp gcf.Credential
	mconn := gcf.MongoCreateConnection(MongoEnv, dbname)
	var datauser peda.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		if peda.IsPasswordValid(mconn, Colname, datauser) {
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(Privatekey))
			if err != nil {
				resp.Message = "Gagal Encode Token : " + err.Error()
			} else {
				resp.Status = true
				resp.Message = "Selamat Datang"
				resp.Token = tokenstring
			}
		} else {
			resp.Message = "Password Salah"
		}
	}
	return gcf.ReturnStringStruct(resp)
}
