package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "clarifaicli"
  app.Usage = "Use Clarifai from the command line!"
  app.Action = func(c *cli.Context) {
    cmd := exec.Command("curl", "-H", "Authorization: Bearer " + c.Args()[0], "--data-urlencode", "url=" + c.Args()[1], "https://api.clarifai.com/v1/tag/")
    data, err := cmd.CombinedOutput()
    if err != nil {
      log.Fatalf("Unable to execute curl because:%v", err)
    }
    fmt.Fprintln(os.Stdout, string(data))
  }

  app.Run(os.Args)
}

