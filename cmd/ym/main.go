package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "github.com/ghodss/yaml"
    "encoding/json"
)

func main() {
  byteArray, _ := ioutil.ReadAll(os.Stdin)
  var o interface{}
  var y []byte

  err := json.Unmarshal(byteArray, &o)
  if err == nil {
    // JSON decoding succeeded, it's JSON
    y, err = yaml.JSONToYAML(byteArray)
  } else {
    // JSON decoding failed, it's probably YAML
    y, err = yaml.YAMLToJSON(byteArray)
  }

  if err != nil {
    fmt.Fprintf(os.Stderr, "ym: %v", err)
  }
  fmt.Println(string(y))
}
