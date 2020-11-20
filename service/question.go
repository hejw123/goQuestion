package service

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"question/models"
	"strconv"
)

// 问题列表
func List(w http.ResponseWriter, r *http.Request) {

	// 设置Header头
	w.Header().Set("content-type", "text/json")
	get := r.URL.Query()

	// 参数判断
	//iType  , err := strconv.Atoi(get.Get("type"))
	//if err != nil {
	//	t := map[string]interface{}{}
	//	t["msg"]  = "查询失败"
	//	t["code"] = 400
	//	t["data"] = []interface{}{}
	//	s  , _  := jsoniter.Marshal(t)
	//	w.Write([]byte(s))
	//	return
	//}
	iType := 1
	// 参数判断
	iClass, err := strconv.Atoi(get.Get("class"))
	fmt.Println(iClass)
	if err != nil {
		t := map[string]interface{}{}
		t["msg"] = "查询失败"
		t["code"] = 400
		t["data"] = []interface{}{}
		s, _ := jsoniter.Marshal(t)
		w.Write([]byte(s))
		return
	}

	Lists, err := models.GetClassQuestion(iClass, iType)
	if err != nil {
		t := map[string]interface{}{}
		t["msg"] = "查询失败"
		t["code"] = 400
		t["data"] = []interface{}{}
		s, _ := jsoniter.Marshal(t)
		w.Write([]byte(s))
		return
	}

	Data, _ := jsoniter.MarshalToString(Lists)
	t := map[string]interface{}{}
	t["msg"] = "查询成功"
	t["code"] = 200
	t["data"] = Data
	s, _ := jsoniter.Marshal(t)
	w.Write([]byte(s))
}

// 创建问题
func Create(w http.ResponseWriter, r *http.Request) {

	// 设置Header头
	w.Header().Set("content-type", "text/json")

	sQuestion := r.PostFormValue("q")
	sClass := r.PostFormValue("c")

	iClass, err := strconv.Atoi(sClass)

	if err != nil || iClass < 0 {
		w.Write([]byte("{'msg':'添加失败','code':401}"))
		return
	}

	Q := models.Question{
		Question: strconv.Quote(sQuestion),
		Class:    iClass,
		Type:     1,
	}

	models.InsertQuestion(&Q)
	t := map[string]interface{}{}
	t["msg"] = "添加成功"
	t["code"] = 200
	s, _ := jsoniter.Marshal(t)
	w.Write([]byte(s))
}

// 更新问题状态
func Update(w http.ResponseWriter, r *http.Request) {

	// 设置header头
	w.Header().Set("content-type", "text/json")

	sId := r.PostFormValue("i")
	sType := r.PostFormValue("t")

	iId, err := strconv.Atoi(sId)
	if err != nil || iId < 0 {
		w.Write([]byte("{'msg':'修改失败','code':401}"))
	}

	iType, err := strconv.Atoi(sType)
	if err != nil || iType > 5 {
		w.Write([]byte("{'msg':'修改失败','code':401}"))
	}

	models.UpdateQuestion(iId, iType)

	t := map[string]interface{}{}
	t["msg"] = "添加成功"
	t["code"] = 200
	s, _ := jsoniter.Marshal(t)
	w.Write([]byte(s))
}
