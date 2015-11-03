import fs from 'fs'
import http from 'http'
import React from 'react'
import { renderToString } from 'react-dom/server'
import { match, RoutingContext } from 'react-router'
import routes from './routes'

const PORT = process.env.PORT || 8080

const readFile = path => new Promise((resolve, reject) => {
  fs.readFile(path, (err, data) => {
    if (err) {
      reject(err)
    } else {
      resolve(data.toString())
    }
  })
}).catch(err => {
  console.log(err)
  throw err
})

readFile('./dist/index.html').then(html => {
  const serve = ((req, res) => {
    readFile('./dist/' + req.url.replace(/^\//, '')).then(data => {
      res.statusCode = 200
      if (req.url.match(/\.css$/)) {
        res.setHeader('Content-Type', 'text/css')
      } else if (req.url.match(/\.js$/)) {
        res.setHeader('Content-Type', 'application/javascript')
      }
      res.end(data)
      console.log(res.statusCode, req.url)
    }).catch(err => {
      match({ routes, location: req.url }, (error, redirectLocation, renderProps) => {
        if (error) {
          res.statusCode = 500
          res.end(error.message)
        } else if (redirectLocation) {
          res.statusCode = 302
          res.setHeader('Location', redirectLocation.pathname + redirectLocation.search)
          res.end('Found')
        } else if (renderProps) {
          res.statusCode = 200
          const app = renderToString(<RoutingContext {...renderProps} />)
          const page = html.replace(/<!-- react-app -->/, app)
          res.end(page)
        } else {
          res.statusCode = 404
          res.end('Not found')
        }
      })
      console.log(res.statusCode, req.url)
    })
  })

  http
    .createServer(serve)
    .listen(PORT, () => console.log("Server listening on: http://localhost:%s", PORT))
})
