package main

import (
    "plugin"
    "os"
    "os/exec"
    "io/ioutil"
    "path/filepath"
    "log"
)

func buildPlugin(taskNumber string) {
    taskPath := filepath.Join("src/", taskNumber + ".go")
    cmd := exec.Command("go", "build", "-buildmode=plugin", taskPath)
    stderr, err := cmd.StderrPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    msg, err := ioutil.ReadAll(stderr)
    if err != nil || string(msg) != "" {
        log.Fatal(err, string(msg))
    }
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func runPlugin(taskNumber string) {
    so := taskNumber + ".so"
    plug, err := plugin.Open(so)
    if err != nil {
        log.Fatal(err)
    }
    fn, err := plug.Lookup("Fn")
    if err != nil {
        log.Fatal(err)
    }
    fn.(func())()
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Provide task number as the first argument")
    }
    taskNumber := os.Args[1]
    log.Println("Build task")
    buildPlugin(taskNumber)
    log.Println("Run task")
    runPlugin(taskNumber)
}

