package hook

import (
    "bytes"
    "log"
    "os/exec"
)


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
