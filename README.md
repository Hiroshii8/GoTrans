# GoTrans
Simple Translator API with golang

you can directly use the the translator by importing translate package to your apps.

# How to Use
```
import translate "github.com/Hiroshii8/GoTrans/translate"

// init translate
trans := translate.New()

// call Request method for translate the text
transResp, err := trans.Request(fromLanguage, toLanguage, text)
```
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