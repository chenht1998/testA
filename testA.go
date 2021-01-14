package testA

import (
    log "github.com/sirupsen/logrus"
)

func Init(){
    log.Info("testA init")
}

func Exec(){
    log.Info("testA exec")
}
