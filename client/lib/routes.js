import React from 'react'
import Router from 'react-router'
import { IndexRoute, Route } from 'react-router'
import Rooms from './components/rooms'
import Room from './components/room'
import NewRoom from './components/new-room'

const routes = (
  <Router>
    <Route path="/">
      <IndexRoute component={ Rooms } />
      <Route path="/new" component={ NewRoom } />
      <Route path="/:path" component={ Room } />
    </Route>
  </Router>
)

export default routes
