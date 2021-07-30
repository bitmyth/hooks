package hook

import (
    "hook/src/job"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func MakeHandler(jobs *job.Jobs) func(writer http.ResponseWriter, request *http.Request) {

    return func(writer http.ResponseWriter, request *http.Request) {

        path := request.URL.Path
        log.Println("request path:", path)

        for _, job := range jobs.Jobs {
            if job.Url == path {
                log.Println("matched", job.Url)

                if job.WorkDir != "" {
                    if err := os.Chdir(job.WorkDir); err != nil {
                        log.Printf("Error:\tCould not move into the directory (%s)\n", job.WorkDir)
                        return
                    }
                }

                c := &Cli{}
                cmd := exec.Command(job.Command[0], job.Command[1:]...)
                o, e := c.Run(cmd)
                if len(e) > 0 {
                    log.Println(string(e))
                    writer.Write(e)
                }
                writer.Write(o)

                return
            }
        }

        writer.Write([]byte("no job"))
    }
}
