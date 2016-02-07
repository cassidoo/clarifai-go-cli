Clarifai CLI
=============

A simple Go-based command line interface for the [Clarifai](http://clarifai.com/) /v1/tag endpoint.

##To Run

First, make sure you have [Go](https://golang.org/dl/) installed, and your
[`GOPATH`](https://golang.org/doc/code.html#GOPATH) is set up.

Download [clarifaicli.go](https://github.com/cassidoo/clarifaicli/blob/master/clarifaicli.go)
or clone this repository, and run:

```
go install
```

Note: You might have to restart your terminal for the command to register.

Next, get an access token from the [Clarifai Developer site](http://developer.clarifai.com)
and an image URL you'd like to tag.

Now, you can run the following and get tags with ease!

```
clarifaicli accesstoken imageurl.jpg
```
