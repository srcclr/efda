package main

import "fmt"
import "github.com/google/go-querystring/query"
import "github.com/srcclr/efda/golang/dep/dep-override/sub"
import "github.com/srcclr/efda/golang/dep/dep-override/sub2"
import "github.com/russellhaering/gosaml2"

type Options struct {
    Query   string `url:"q"`
    ShowAll bool   `url:"all"`
    Page    int    `url:"page"`
}

func main() {
    fmt.Println("Intro to Go!")
    opt := Options{ "foo", true, 2 }
    v, _ := query.Values(opt)
    fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"

    sub.Foo()

    sub2.Bar()

    ep := &saml2.ErrParsing {
        Tag: "tag",
        Value: "some text",
    }

    ep.Error()
}
