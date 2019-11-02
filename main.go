package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {
	s := &http.Server{
		Addr:           ":9002",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/cover/admin/apps", mock)
	http.HandleFunc("/cover/admin/user", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "{\"username\":\"wagweiqi\"}")
	})
	http.HandleFunc("/cover/admin/add", add)
	log.Fatal(s.ListenAndServe())
}

func mock(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r)
	fmt.Fprintf(w, "{\"code\":0,\"message\":\"0\",\"ttl\":1,\"data\":{\"total\":0,\"page_num\":1,\"page_size\":10,\"apps\":[{\"id\":1,\"name\":\"thum_service\",\"env\":\"uat\",\"app_id\":\"appid\",\"language\": \"golang\",\"git_url\":\"xxxx.git\",\"rd_direct\": \"veitch\",\"test_direct\": \"veitch\",\"creator\":\"veitch\",\"ops_time\":\"2019-11-01 10:10:10\"},{\"id\":2,\"name\":\"thum_service1\",\"app_id\":\"appid\",\"env\":\"uat\",\"language\": \"golang\",\"git_url\":\"xxxx.git\",\"rd_direct\": \"veitch\",\"test_direct\": \"veitch\",\"creator\":\"veitch\",\"ops_time\":\"2019-11-01 10:10:10\"}]}}")
}

func add(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body)
	fmt.Fprintf(w, "添加成功")
}
