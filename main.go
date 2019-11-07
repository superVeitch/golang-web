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
		fmt.Fprintf(w, "{\"code\":0,\"message\":\"0\",\"ttl\":1,\"data\":{\"id\":333,\"username\":\"wangweiqi\",\"email\":\"wangweiqi@bilibili.com\",\"token\":\"b87c0de1-dff0-4343-838f-947b06f2ef65\",\"CTime\":\"2019-10-23T19:15:12+08:00\",\"UTime\":\"2019-10-23T19:15:12+08:00\"}}")
	})
	http.HandleFunc("/cover/admin/inc/detail", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "{\"code\":0,\"message\":\"0\",\"ttl\":1,\"data\":{\"name\":\"thum_service\",\"env\":\"uat\",\"cover_total\":0,\"func_cover_total\":1,\"hit_line_total\":2,\"miss_line\":3,\"snapshot_id\":2,\"coverages\":[{\"path\":\"src/app/admin/cmd\",\"cover\":80,\"func_cover\":70,\"branch_cover\":70,\"is_dir\":true,\"hit_line\":2333,\"miss_line\":34}, {\"path\":\"src/app/admin/cmd/main.go\",\"cover\":560,\"func_cover\":90,\"branch_cover\":20,\"is_dir\":false,\"hit_line\":2333,\"miss_line\":34}]}}")
	})
	http.HandleFunc("/cover/admin/file/detail", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "{\"code\":0,\"message\":\"0\",\"ttl\":1,\"data\":{\"name\":\"thum_service\",\"env\":\"uat\",\"cover_total\":0,\"func_cover_total\":1,\"hit_line_total\":2,\"miss_line\":3,\"snapshot_id\":2,\"coverages\":[{\"num\": 1,\"hit\": 3,\"content\": \"import (\"},{\"num\": 2,\"hit\": 90,\"content\": \"  fmt\"},{\"num\": 3,\"hit\": 14,\"content\": \"func main(){\"},{\"num\": 4,\"hit\": 43,\"content\": \"flat.prase()\"}]}}")
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
