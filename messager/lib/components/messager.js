import Style from './messager.scss'
import React from 'react'
import Router from 'react-router'
import { IndexRoute, Route } from 'react-router'
import Rooms from './rooms'
import Room from './room'
import NewRoom from './new-room'

export default class Messager extends React.Component {

  render () {
    return (
      <Router>
        <Route path="/">
          <IndexRoute component={ Rooms } />
          <Route path="/new" component={ NewRoom } />
          <Route path="/:path" component={ Room } />
        </Route>
      </Router>
    )
  }

}
