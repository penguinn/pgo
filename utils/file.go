package utils

import "os"

func FileExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}
