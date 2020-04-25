package main

import (
    // "fmt"
    "log"
    "os"
    "os/user"
    "path"
    "path/filepath"
    "strings"
    "io/ioutil"
)

func ListSubdir(path string) ([]os.FileInfo, error) {

    usr, _ := user.Current()
    dir := usr.HomeDir

    if path == "~" {
        path = dir
    } else if strings.HasPrefix(path, "~/") {
        path = filepath.Join(dir, path[2:])
    }

    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    // for _, f := range files {
    //     fmt.Println(f.Name())
    // }

    return files, err
}



func IsDir(path string) (bool, error) {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return false, err
    }
    return fileInfo.IsDir(), err
}

func IsFile(name string) (bool, error) {
    fileInfo, err := os.Stat(name)
    if err != nil {
        return false, err
    }
    if fileInfo.Mode().IsRegular() {
        return true, nil
    }
    return false, nil
}

func FilenameWithoutExtension(name string) string {
    return strings.TrimSuffix(name, path.Ext(name))
}