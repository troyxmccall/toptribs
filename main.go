package main

import (
  "os"

  "github.com/codegangsta/cli"
  "github.com/troyxmccall/toptribs/command"
)

func main() {
  app := cli.NewApp()
  app.Name = "toptribs"
  app.Usage = "Show the top contributors for specified files"
  app.Version = "0.1.0"
  app.Action = command.Contrib
  app.Flags = []cli.Flag{command.ContribFlag}
  app.Run(os.Args)
}
