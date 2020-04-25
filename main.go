package main

import (
    "fmt"
    "log"
    "os"
    "os/user"
    "strings"
    "flag"
    "context"
    "path/filepath"
    "net/http"
    // "strconv"
    // "time"

    // bolt "go.etcd.io/bbolt"
    // "github.com/jawher/mow.cli"
    "github.com/KyleBanks/goodreads"

    // "golang.org/x/oauth2"
    // "golang.org/x/oauth2/google"
    googlebooks "google.golang.org/api/books/v1"
    // "google.golang.org/api/googleapi/transport"
)

var defaultPath = "~/Documents/archives/books"

func main() {

    pathPtr := flag.String("path", defaultPath, " The directory for library.")

    base := *pathPtr
    _ = base

    // db, err := bolt.Open("my.db", 0666, &bolt.Options{Timeout: 3 * time.Second})
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer db.Close()

    // initialize Goodreads API
    key := os.Getenv("GOODREADS_API_KEY")
    if key == "" {
        fmt.Println("Missing required env var: GOODREADS_API_KEY")
        os.Exit(1)
    }
    client := goodreads.NewClient(key)

    // Google Books API
    ctx := context.Background()
    svc, err := googlebooks.New(http.DefaultClient)
    if err != nil {
        log.Fatalf("Unable to create Books service: %v", err)
    }

    books, _ := ListSubdir(*pathPtr)

    for _, book := range books {
        if book.Name()[0:1] == "." {
            continue
        }
        fmt.Println(book.Name())
        record, err := client.SearchBooks(book.Name(), 10, goodreads.AllFields)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(record)

        volumes, err := svc.Volumes.List(book.Name()).Context(ctx).Do()
        if err != nil {
            log.Fatal(err)
        }

        for i, v := range volumes.Items {
        fmt.Printf("%d. %s\n", i+1, v.VolumeInfo.Title)
        }
        break
    }
}


func pathWalker(path string) ([]string, error) {

    fileList := make([]string, 0)

    usr, _ := user.Current()
    dir := usr.HomeDir

    if path == "~" {
        path = dir
    } else if strings.HasPrefix(path, "~/") {
        path = filepath.Join(dir, path[2:])
    }

    e := filepath.Walk(path, func(currentPath string, f os.FileInfo, err error) error {
        fileList = append(fileList, currentPath)
        return err
    })

    if e != nil {
        panic(e)
    }

    for _, file := range fileList {
        fmt.Println(file)
        if ret, _ := IsDir(file); ret == true {

        }

        if ret, _ := IsFile(file); ret == true {
            ext := filepath.Ext(file)
            base := filepath.Base(file)
            filename := FilenameWithoutExtension(base)
            fmt.Println(filename)
            _ = ext
        }

    }
    return fileList, nil
}