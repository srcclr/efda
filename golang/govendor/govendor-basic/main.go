package main

import "strconv"
import "fmt"
import "crypto/rsa"
import "crypto/rand"
import "github.com/google/go-querystring/query"
import "github.com/google/go-github/github"
import "github.com/square/go-jose"

type Options struct {
    Query   string `url:"q"`
    ShowAll bool   `url:"all"`
    Page    int    `url:"page"`
}

func main() {
    // Use go-querystring
    fmt.Println("Intro to Go!")
    opt := Options{ "foo", true, 2 }
    v, _ := query.Values(opt)
    fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"

    // Use go-github
    client := github.NewClient(nil)

    // list public repositories for org "github"
    publicRepos := &github.RepositoryListByOrgOptions{Type: "public"}
    repos, _, _ := client.Repositories.ListByOrg("github", publicRepos)

    fmt.Println("Hello from sub! Num repos: " + strconv.Itoa(len(repos)))

    // Use go-jose
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    publicKey := &privateKey.PublicKey
    encrypter, err := jose.NewEncrypter(jose.RSA_OAEP, jose.A128GCM, publicKey)
    if err != nil {
        panic(err)
    }

    var plaintext = []byte("Lorem ipsum dolor sit amet")
    object, err := encrypter.Encrypt(plaintext)
    if err != nil {
        panic(err)
    }

    fmt.Println(object)
}
