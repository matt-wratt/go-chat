'use strict'

const webpack = require('webpack')

module.exports = {

  entry: [
    'webpack/hot/only-dev-server',
    './lib/application.js'
  ],

  output: {
    filename: 'application.js'
  },

  module: {
    loaders: [
      { test: /\.js$/, loaders: ['react-hot', 'babel'], exclude: /node_modules/ },
      { test: /\.js$/, loader: 'babel-loader', exclude: /node_modules/ },
      { test: /\.scss$/, loader: 'style-loader!css-loader!sass-loader' }
    ]
  },

  plugins: [
    new webpack.NoErrorsPlugin()
  ]

}
