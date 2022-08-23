package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/Slimyi/FemboyDesktopFlooder/r34dl"
)

func openFile(fname string, dir string) {
	open := exec.Command("cmd.exe", "/c", dir+fname)
	err := open.Run()
	if err != nil {
		fmt.Println("error with file" + dir + fname)
	}
}

func writeFile(result r34dl.Posts, dir string, n int) {
	fileurl := result.Post[n].FileURL
	fileres, err := http.Get(fileurl)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(fileres.Body)
	if err != nil {
		panic(err)
	}
	var ftype string
	if fileurl[len(fileurl)-4:] == ".jpg" || fileurl[len(fileurl)-5:] == ".jpeg" {
		os.WriteFile(dir+result.Post[n].ID+".jpg", data, os.ModePerm)
		ftype = ".jpg"
	} else if fileurl[len(fileurl)-4:] == ".png" {
		os.WriteFile(dir+result.Post[n].ID+".png", data, os.ModePerm)
		ftype = ".png"
	} else if fileurl[len(fileurl)-4:] == ".mp4" {
		os.WriteFile(dir+result.Post[n].ID+".mp4", data, os.ModePerm)
		ftype = ".mp4"
	} else {
		fileres.Body.Close()
	}
	fileres.Body.Close()
	if ftype != "" {
		openFile(result.Post[n].ID+ftype, dir)
	}
}

func initial(wg *sync.WaitGroup, p int) {
	defer wg.Done()
	result := r34dl.Fetch("femboy", 50, p)
	fmt.Println(len(result.Post))
	filenum := 0
	var stdout bytes.Buffer
	userCom := exec.Command("cmd.exe", "/c", "echo", "%username%")
	userCom.Stdout = &stdout
	err := userCom.Run()
	if err != nil {
		panic(err)
	}
	outStr := string(stdout.Bytes()[:len(stdout.Bytes())-2])
	dir := "/Users/" + outStr + "/Desktop/"
	fmt.Println(dir)
	for i := 0; i < len(result.Post); i++ {
		writeFile(result, dir, i)
		filenum++
	}
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go initial(wg, i)
	}
	wg.Wait()
}
