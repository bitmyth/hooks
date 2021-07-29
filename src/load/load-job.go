package load

import (
    "io/ioutil"
    "log"
)

func Load(file string) []byte {
    jobDefinition, e := ioutil.ReadFile(file)
    if e != nil {
        log.Fatalln(e.Error())
    }
    return jobDefinition
}
