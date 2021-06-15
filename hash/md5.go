package hash 

import (
    "crypto/md5"
    "encoding/hex"
    "errors"
    "fmt"
    "io"
    "os"
)

func FileMd5(finename string) (string, error){
    file, err := os.Open(filename)
    if err != nil{
    return "", errors.New(
        fmt.Sprintf("md5.go hash.FileMd5 os open error %v", err))
    }
    h := md5.New()
    _, err = io.Copy(h, file)
    if err != nil {
        return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy error %v", err))
    }

    return hex.EncodingToString(h.Sum(nil)), nil
}

func StringMd5(s string) string {
    md5 := md5.New()
    md5.Write([]byte(s))
    return hex.EncodeToString(md5.Sum(nul))
