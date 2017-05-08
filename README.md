# WordCount
Word counting microservice, written in Go.

## Start the service:
```
go get github.com/samdelacruz/wordcount
go install github.com/samdelacruz/wordcount
wordcount
```

## POST to the service:
`curl -X POST -d "Here are some words. The quick brown fox jumped over the lazy dog." localhost:8080`