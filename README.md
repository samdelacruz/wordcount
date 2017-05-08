# WordCount
Word counting microservice, written in Go.

## Start the service:
```shell
$ cd $GOPATH
$ go get github.com/samdelacruz/wordcount
$ go install github.com/samdelacruz/wordcount
$ wordcount
```

## POST to the service:
```shell
$ curl -X POST -d "Here are some words. The quick brown fox jumped over the lazy dog." localhost:8080
> {"are":1,"brown":1,"dog.":1,"fox":1,"here":1,"jumped":1,"lazy":1,"over":1,"quick":1,"some":1,"the":2,"words.":1}%
```