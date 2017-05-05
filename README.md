<h1 align="center">Rapid</h1>

[![Build Status](https://travis-ci.org/Kiricon/Rapid.svg?branch=master)](https://travis-ci.org/Kiricon/Rapid)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<p align="center">
Rapid is a lightweight micro-framework built for quickly developing webservers and rest apis.<br/>

With a syntax and routing system inspired by Express.js, Flask & Laravel, developing end points is easy to pick up. 
</p>

<h3 align="center">Install Rapid</h3>

`go get https://github.com/Kiricon/Rapid`


<h2 align="center">Hello World</h3>

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