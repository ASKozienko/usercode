package main

import (
        "context"

        "github.com/corezoid/gitcall-go-runner/gitcall"
)

func main() {
        gitcall.Handle(func (_ context.Context, data map[string]interface{}) error {
        data["foo"] = "bar"
       
        return nil
   })
}
