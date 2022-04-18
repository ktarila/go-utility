# go-stimulus

Starter project with [golang](https://golang.org) backend and [Stimulus JS](https://stimulusjs.org) on the front-end. 
StimulusJS, as per their website, is suitable for the applications that have server-side rendered HTML at their core,
and want to sprinkle Javascript to make them sparkle. This project also uses Webpack to bundle Javascript and CSS, 
and Babel to transpile ES6 code back to vanilla JavaScript so that every environment (e.g. browser) can interpret it.

Golang backend is using [Fiber - Express inspired web framework written in Go](https://github.com/gofiber) for APIs, 
serve HTML, and other static content such as Javascript and CSS. HTML can be rendered in server-side using 
[HTML Template Engine ](https://github.com/gofiber/template) supported by Fiber.


#### To run the project

```
$ cd go-stimulus
$ go get 
$ cd webapp
$ yarn install
$ yarn start
```
go to http://localhost:3000

