package main

import (
  "archive/zip"
  "bytes"
  "fmt"
  "golang.org/x/text/encoding/simplifiedchinese"
  "golang.org/x/text/transform"
  "io"
  "io/ioutil"
  "os"
  "path/filepath"
)

func main() {
  var path string = `F://test/zip/test.zip`
  var outDir string = "F://test/upzip/"
  upzip(path, outDir)
}

func upzip(path string, outDir string) {
  rc, err := zip.OpenReader(path)
  if err != nil {
    fmt.Println(err)
    defer rc.Close()
  }

  for _, file := range rc.Reader.File {
    i := bytes.NewReader([]byte(file.Name))
    decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
    content, _ := ioutil.ReadAll(decoder)
    if file.FileInfo().IsDir() {
      e1 := os.MkdirAll(outDir+string(content), 0644)
      if e1 != nil {
        fmt.Println(e1)
      }
      continue
    }
    //输出文件名
    _, filename := filepath.Split(string(content))
    fmt.Println(filename)

    f, e := file.Open()
    if e != nil {
      fmt.Println(e)
      continue
    }
    //fmt.Println("unzip:",string(content))
    defer f.Close()
    NewFile, e2 := os.Create(outDir + string(content))
    if e2 != nil {
      fmt.Println(e2)
      continue
    }
    io.Copy(NewFile, f)
    NewFile.Close()
  }
}
