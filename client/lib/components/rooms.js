import React from 'react'
import { Link } from 'react-router'
import Api from '../api'

export default class Rooms extends React.Component {

  componentDidMount () {
    this.interval = setInterval(() => {
      Api.rooms()
        .then(rooms => this.setState({ rooms, error: null }))
        .catch(error => this.setState({ error: error.message }))
    }, 300)
  }

  componentWillUnmount () {
    clearInterval(this.interval)
  }

  render () {
    this.state = this.state || {}
    const rooms = (this.state.rooms || []).map(room => {
      return (
        <div>
          <Link to={ `/${room.path}` }>{ room.name } ({room.messages.length})</Link>
        </div>
      )
    })
    const error = this.state.error || null
    return (
      <div>
        <h1>Rooms</h1>
        <Link to="/new">New Room</Link>
        <br />
        <div className="error">{ error }</div>
        <br />
        { rooms }
      </div>
    )
  }

}
