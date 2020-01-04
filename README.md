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

make sure the depedency by typing 
``
dep ensure -v -vendor-only
``

and then run the apps with 
``
go run app.go
``

## current folder tree

```
├── app.go
├── entity
│   └── json
│       └── json.go
├── Gopkg.lock
├── Gopkg.toml
├── handler
│   └── handler.go
├── README.md
├── server
│   └── server.go
├── translate
│   └── translate.go
├── util
│   ├── validator.go
│   └── validator_test.go
└── vendor
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
