package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "hook/src/hook"
    "hook/src/job"
    "hook/src/load"
    "log"
    "net/http"
    _ "net/http/pprof"
    "os"
    "os/signal"
    "syscall"
)

var jobs *job.Jobs

func main() {
    file := load.Default(os.Getenv("FILE"), "jobs.yaml")
    port := load.Default(os.Getenv("PORT"), "8000")

    data := load.Load(file.(string))

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

    go func() {
        _, _ = os.Stdout.Write([]byte(fmt.Sprintf("Server listen on port %s\n", port)))
        if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
        //log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown Server ...")

}
