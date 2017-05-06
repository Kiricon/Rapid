<h1 align="center">Rapid</h1>

[![Build Status](https://travis-ci.org/Kiricon/Rapid.svg?branch=master)](https://travis-ci.org/Kiricon/Rapid)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kiricon/Rapid)](https://goreportcard.com/report/github.com/Kiricon/Rapid)
[![codecov](https://codecov.io/gh/Kiricon/Rapid/branch/master/graph/badge.svg)](https://codecov.io/gh/Kiricon/Rapid)
[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.svg)](https://godoc.org/github.com/Kiricon/Rapid)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<p align="center">
Rapid is a lightweight micro-framework built for quickly developing webservers and rest apis.<br/>

With a syntax and routing system inspired by Express.js, Flask & Laravel, developing end points is easy to pick up. 
</p>

<br/>

<h2 align="center">Install Rapid</h2>
<p align="center">
<code align="center">go get https://github.com/Kiricon/Rapid</code>
</p>

<br/>
<h2 align="center">Hello World</h2>

```Go
package main

import (
	r "github.com/Kiricon/Rapid"
)

func main() {

    r.Get("/", func(c r.Connection) {
        c.Send("Hello World")
    })

	r.Listen(3000)
}
```
