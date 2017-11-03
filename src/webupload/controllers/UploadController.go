package controllers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"yhl/help"
)

type UploadController struct {
	help.BaseController
}

func (this *UploadController) Webupload() {
	help.Log.Info(this.Ctx.Input.Domain())
	method := this.Ctx.Input.Method()
	help.Log.Info("method:" + method)
	if method == "OPTIONS" {
		this.StopRun()
	}
	filename := this.Ctx.Input.Query("name")
	chunks := this.Ctx.Input.Query("chunks")
	chunk := this.Ctx.Input.Query("chunk")
	f, h, err := this.GetFile("file")
	if err != nil {
		help.Log.Info(err.Error())
	}
	defer f.Close()
	ext := filepath.Ext(h.Filename)
	filename = help.Md5(filename) + ext
	prefix := "tmp/"
	part := prefix + filename + "_" + chunk + ".part"
	this.SaveToFile("file", part)
	count, err := strconv.Atoi(chunks)
	cache := help.Cache
	cache.Incr(filename)
	num, err := strconv.Atoi(string(cache.Get(filename).([]uint8)))
	dir := "uploads/"
	if this.Ctx.Input.Domain() == "frontend.feichangjuzu.com" {
		dir = "uploads-test/"
	}
	y, m, d := help.Date()
	dir = dir + fmt.Sprintf("%d/%d/%d/", y, m, d)
	if num == count {
		log.Println("==================== num:", num)
		go func(prefix, filename, dir, ext string) {
			cache.Delete(filename)
			outDir := "" + dir
			if !help.PathExist(outDir) {
				os.MkdirAll(outDir, os.ModePerm)
			}
			outfile := outDir + fmt.Sprintf("%d%d", time.Now().Unix(), help.RandNum(10000, 99999)) + ext
			out, _ := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0600)
			bWriter := bufio.NewWriter(out)
			for i := 0; i < count; i++ {
				infile := prefix + filename + "_" + strconv.Itoa(i) + ".part"
				in, _ := os.Open(infile)
				bReader := bufio.NewReader(in)
				for {
					buffer := make([]byte, 1024)
					readCount, err := bReader.Read(buffer)
					if err == io.EOF {
						os.Remove(infile)
						break
					} else {
						bWriter.Write(buffer[:readCount])
					}
				}
				log.Println("==================== i:", i)
			}
			bWriter.Flush()

		}(prefix, filename, dir, ext)
	}

	help.Log.Info("filename:" + filename + " chunks:" + chunks + " chunk:" + chunk)
	this.SendResJsonp(0, "ok", dir+filename)
}
