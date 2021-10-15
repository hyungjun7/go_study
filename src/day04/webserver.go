package main

import (
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Go에서 소문자로 시작하는 필드는 private 이기 때문에
//소문자로 필드를 만든 경우 JSON 응답 메세지에 포함되지 않는다.

type JsonRes struct {
	Message string `json:"message"`
	Author  string `json:"-"`          //해당 필드를 출력하지 않도록 설정함
	Date    string `json:",omitempty"` //값이 비어있다면 필드를 출력하지 않음
	Id      int    `json:"id,string"`  //출력을 문자열로 변환한 다음 id 로 필드명을 변경
}

//json 데이터를 주고 받을 때 구조체에 태그를 설정하면 성능 향상이 있음.

type JsonReq struct {
	Name string `json:"name"`
}

func main() {
	//서버 실행할 포트 설정
	port := 8080

	//라우터 등록
	http.HandleFunc("/index", handler)

	log.Printf("server starting on %d", port)
	//서버 시작
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

//등록한 라우터에서 사용되는 핸들러
func handler(res http.ResponseWriter, req *http.Request) {

	//아래에서 마샬을 사용하여 디코딩한 것보다 성능향상이 있음
	var reqMsg JsonReq
	decoder := json2.NewDecoder(req.Body)
	err := decoder.Decode(&reqMsg)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	//body, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	http.Error(res, "Bad Request", http.StatusBadRequest)
	//}
	//
	//var reqMsg JsonReq
	//
	//err = json2.Unmarshal(body, &reqMsg)
	//if err != nil {
	//	http.Error(res, "Bad Request", http.StatusBadRequest)
	//}

	jsonMsg := JsonRes{
		Message: "hello " + reqMsg.Name,
		Id:      1,
	}

	//NewEncoder 를 사용하면 임시 객체에 담는 것을 스킵할 수 있음.
	//또한 마샬링하는 것보다 성능이 더 좋음
	encoder := json2.NewEncoder(res)
	_ = encoder.Encode(&jsonMsg)
	//data, err := json2.Marshal(jsonMsg)
	//if err != nil {
	//	panic("err!")
	//}
	//fmt.Println(string(data))
	//_, _ = fmt.Fprint(res, string(data))
}
