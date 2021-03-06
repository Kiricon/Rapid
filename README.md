<p align="center">
    <img src="https://github.com/Kiricon/Rapid/blob/master/demo/public/imgs/logo.png?raw=true" height="200" />
</p>

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
<code align="center">go get github.com/Kiricon/Rapid</code>
</p>

<br/>


# Hello World

```Go
package main

import (
	"github.com/Kiricon/Rapid"
)

func main() {

    app := rapid.App()

    app.Get("/", func(c rapid.Connection) {
        c.Send("Hello World")
    })


    app.Listen(3000)
}
```

## Features

- Routing
- Url parameters
- Templating
- Static file serving
- Zero Dependencies (Only standard Library calls)
