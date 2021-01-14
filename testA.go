package testA

import (
    log "github.com/sirupsen/logrus"
    "github.com/jordan0210/releaseGoModTest"
)

func Init(){
    log.Info("github testA init")
}

func Exec(){
    myGoMod.Main()
    log.Info("github testA exec")
}
