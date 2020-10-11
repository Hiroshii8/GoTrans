# GoTrans

Simple Translator API with golang

you can directly use the the translator by importing translate package to your apps.

## How to Use Translate package

```
import translate "github.com/Hiroshii8/GoTrans/translate"

// init translate
trans := translate.New()

// call Request method for translate the text
transResp, err := trans.Request(fromLanguage, toLanguage, text)
```

## How to run this apps

You can run this apps by typing

``
go run app.go
``

for testing the GoTrans via http or grpc

## current folder tree

```
├── app.go
├── entity
│   └── json
│       └── json.go
├── go.mod
├── go.sum
├── grpc
│   ├── Makefile
│   └── proto
│       ├── translate.pb.go
│       └── translate.proto
├── handler
│   └── handler.go
├── README.md
├── server
│   ├── grpc.go
│   └── server.go
└── translate
    └── translate.go
```

## Coming Soon Feature

```
- Auto reload apps
- NSQ
- Running in Docker (if i'm not lazy enough ^v^)
- Grace Stop
```

## To Any Contributor

Feel free to contribute in here, this repo is meant to be a learning repository to anyone for learning to create webservice with GoLang. So happy learning (^v^)/

## How To Contribute

1. Fork this repo to yours
2. Create feature that you desire to
3. Create pull request and mention me for the update you want to merge. I will review as fast as i can.
