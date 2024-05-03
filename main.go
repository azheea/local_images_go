package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var port = ":5000"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error retrieving file from form:", err)
		return
	}
	defer file.Close()

	// 创建本地文件
	localFile, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error creating local file:", err)
		return
	}
	defer localFile.Close()

	// 复制文件内容到本地文件
	_, err = io.Copy(localFile, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error copying file:", err)
		return
	}

	// 返回上传成功的消息和图片访问链接
	fmt.Fprintf(w, "http://0.0.0.0"+port+"/uploads/%s", handler.Filename)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求中的图片文件名
	filename := r.URL.Path[len("/uploads/"):]
	// 读取图片文件
	http.ServeFile(w, r, "./uploads/"+filename)
	fmt.Printf("请求了图片 %s \n", filename)
}

func main() {
	// 创建上传目录
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating uploads directory:", err)
		return
	}

	// 设置路由
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploads/", imageHandler)

	// 启动服务器
	fmt.Println("Local image server started on http://0.0.0.0" + port)
	http.ListenAndServe(port, nil)
}
