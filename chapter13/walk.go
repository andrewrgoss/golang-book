package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error { // The function you pass to Walk is called for every file and folder in the root folder. (in this case '.')/ This recursively walks a folder (read the folder's contents, all the sub-folders, all the sub-sub-folders, …)
    fmt.Println(path)
    return nil
  })
}