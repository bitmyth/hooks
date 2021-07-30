package job

type Job struct {
    Name    string
    Url     string
    WorkDir string
    Command []string
}

type Jobs struct {
    Jobs []Job
}
