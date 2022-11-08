package main

import (
    cmd "github.com/bogdzn/gork/cmd"
)

func main() {
    opts := &cmd.Options{}
    parser := cmd.GetCmdParser(opts)

    parser.ExecuteC()
}
