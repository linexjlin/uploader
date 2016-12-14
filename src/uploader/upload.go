// 处理/upload 逻辑
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func isExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil /*&& !info.IsDir() */ {
		return true
	}
	return false
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	log.Println(path)

	if err != nil {
		return true
	} else {
		return info.IsDir()
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //获取请求的方法

	path := BasePath + r.URL.Path

	if r.Method == "GET" {
		if exists := isExists(path); exists {
			http.ServeFile(w, r, path)
		} else {
			http.NotFound(w, r)
		}
	} else {
		if isdir := isDir(path); !isdir { //not a directory, is a file, can't create a file inside a file
			w.Write([]byte("Can not create a file inside a file"))
			return
		}
		checkError(checkCreatePath(path))
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("f")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		//log.Fprintf(w, "%v", handler.Header)
		newFilePath := path + "/" + handler.Filename
		checkError(checkBackUp(newFilePath))
		f, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func checkBackUp(fileName string) error {
	_, err := os.Stat(fileName)
	if err == nil {
		os.Rename(fileName, fileName+".bk."+time.Now().Format("2006-01-02_15-04-05"))
	}
	return nil
}

func checkCreatePath(path string) error {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		checkError(os.MkdirAll(path, 0776))
	}
	return nil
}
