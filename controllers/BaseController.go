package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type BaseController struct {
	beego.Controller
}

func isEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

//创建随机token
func GenToken() string {
	str := rand.Int63n(math.MaxInt8)
	str2 := time.Now().UTC().UnixNano()
	has := md5.Sum([]byte(strconv.Itoa(int(str)) + strconv.Itoa(int(str2))))
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithData(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
