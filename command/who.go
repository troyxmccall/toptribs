package command

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/troyxmccall/toptribs/util"
)

func Contrib(c *cli.Context) {
  if len(c.Args()) == 0 {
    cli.ShowAppHelp(c)
    os.Exit(1)
  }

  names, err := getContributors(c.Args())
  if err != nil {
    log.Fatal(err)
  }

  number := c.Int("number")
  if number == 0 {
    number = 1
  }

  for i, name := range names {
    if i >= number {
      break
    }
    fmt.Println(name)
  }

  os.Exit(0)
}

var ContribFlag = cli.IntFlag{
  Name:  "number, n",
  Usage: "the number of contributors",
}

func getContributors(filenames []string) ([]string, error) {
  var err error
  contributors := make(map[string]int)

  for _, filename := range filenames {
    names, err := gitLog(filename)
    if err != nil {
      continue
    }

    for _, name := range names {
      contributors[name] += 1
    }
  }

  if err != nil {
    return []string{}, err
  }

  vs := util.NewValSorter(contributors)
  vs.Sort()

  return vs.Keys, nil
}

func gitLog(filename string) ([]string, error) {
  // Print only author's name
  cmd := exec.Command("git", "log", "--pretty=%an", filename)
  out, err := cmd.CombinedOutput()
  str := string(out)
  str = strings.Trim(str, "\n")
  return strings.Split(str, "\n"), err
}
