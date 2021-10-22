package utils

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

func BcryptMake(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, pwd)
    if err != nil {
        return false
    }
    return true
}

