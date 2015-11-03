import React from 'react'
import { Link } from 'react-router'
import Api from '../api'
import NewMessage from './new-message'

export default class Room extends React.Component {

  componentDidMount () {
    this.interval = setInterval(() => {
      Api.room(this.props.params.path)
        .then(room => this.setState({ room, error: null }))
        .catch(error => this.setState({ error: error.message }))
    })
  }

  componentWillUnmount () {
    clearInterval(this.interval)
  }

  room () {
    return (this.state || {}).room || {}
  }

  render () {
    const room = this.room()
    const messages = (room.messages || []).map(message => {
      return (
        <div>{ message.value }</div>
      )
    })
    return (
      <div>
        <h1>{ room.name }</h1>
        <Link to="/">Back</Link>
        <br />
        <br />
        { messages }
        <br />
        <NewMessage room={ room }/>
      </div>
    )
  }

}
