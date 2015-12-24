package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "github.com/ghodss/yaml"
)

func main() {
  byteArray, _ := ioutil.ReadAll(os.Stdin)
  y, err := yaml.YAMLToJSON(byteArray)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  fmt.Println(string(y))
}
