package main

import (
    "bytes"
    "fmt"
    "gopkg.in/yaml.v2"
    "log"
    "net/http"
    "os/exec"
)

type Job struct {
    Name    string
    Url     string
    Command []string
}
type T struct {
    Jobs []Job
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {

    path := request.URL.Path
    log.Println("request path:", path)

    fmt.Printf("%v", t)

    for _, job := range t.Jobs {
        if job.Url == path {
            log.Println("matched", job.Url)
            c := &Cli{}
            cmd := exec.Command(job.Command[0], job.Command[1])
            o, e := (c.Run(cmd))
            if string(e) != "" {
                writer.Write([]byte(e))
            }
            writer.Write(o)

            return
        }
    }

    writer.Write([]byte("no job"))
}

var t *T

func main() {
    data := `
jobs:
  - name: test
    url: /test
    command: ['date','-h']
  - name: test2
    url: /test2
    command: ['touch','test2.txt']
`
    err := yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    d, _ := yaml.Marshal(&t)
    fmt.Printf("load jobs:\n%s\n\n", string(d))

    for _, job := range t.Jobs {
        http.HandleFunc(job.Url, IndexHandler)
    }

    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

type Cli struct {
    OutStr string
    ErrStr string
}

func (c *Cli) Run(cmd *exec.Cmd) ([]byte, []byte) {
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        log.Printf("cmd.Run() failed with %s\n", err)
    }

    return stdout.Bytes(), stderr.Bytes()
}
