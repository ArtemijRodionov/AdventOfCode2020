package main

import (
    "plugin"
    "os"
    "os/exec"
    "path/filepath"
    "log"
)

func buildPlugin(taskNumber string) {
    taskPath := filepath.Join("src/", taskNumber + ".go")
    cmd := exec.Command("go", "build", "-buildmode=plugin", taskPath)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
}

func runPlugin(taskNumber string) {
    so := taskNumber + ".so"
    plug, err := plugin.Open(so)
    if err != nil {
        panic(err)
    }
    fn, err := plug.Lookup("Fn")
    if err != nil {
        panic(err)
    }
    fn.(func())()
}

func main() {
    if len(os.Args) < 2 {
        panic("Provide task number as the first argument")
    }
    taskNumber := os.Args[1]
    log.Println("Build task")
    buildPlugin(taskNumber)
    log.Println("Run task")
    runPlugin(taskNumber)
}

