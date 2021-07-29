package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "hook/src/hook"
    "hook/src/job"
    "hook/src/load"
    "log"
    "net/http"
    "os"
)

var jobs *job.Jobs

func main() {
    file := os.Getenv("FILE")
    port := load.Default(os.Getenv("PORT"), "8000")

    data := load.Load(file)

    err := yaml.Unmarshal([]byte(data), &jobs)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    d, _ := yaml.Marshal(&jobs)
    fmt.Printf("Loaded jobs:\n---\n%s\n", string(d))

    handler := hook.MakeHandler(jobs)
    for _, job := range jobs.Jobs {
        http.HandleFunc(job.Url, handler)
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
