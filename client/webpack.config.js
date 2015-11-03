'use strict'

const ExtractTextPlugin = require('extract-text-webpack-plugin')

module.exports = {

  entry: {
    index: './index.html',
    application: './lib/application.js'
  },

  output: {
    path: __dirname + '/dist',
    filename: '[name].js'
  },

  module: {
    loaders: [
      { test: /\.js$/, loader: 'babel-loader', exclude: /node_modules/ },
      { test: /\.scss$/, loader: ExtractTextPlugin.extract('style-loader', 'css-loader!sass-loader') },
      { test: /\.html$/, loader: "file?name=[name].[ext]" }
    ]
  },

  plugins: [
    new ExtractTextPlugin("application.css")
  ]
}
