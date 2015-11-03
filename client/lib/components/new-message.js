import React from 'react'
import { Link } from 'react-router'
import Api from '../api'

export default class NewRoom extends React.Component {

  submit (event) {
    event.preventDefault()
    Api.message(this.props.room.path, this.refs.message.value).then(() => {
      this.setState({ error: null })
      this.refs.message.value = ''
      this.refs.message.focus()
    }).catch(error => {
      this.setState({ error: error.message })
    })
  }

  render () {
    this.state = this.state || {}
    const error = this.state.error || null
    return (
      <form onSubmit={ this.submit.bind(this) }>
        <label htmlFor="message">New Message</label>
        <input id="message" name="message" ref="message" />
        <input type="submit" />
        <br />
        <div className="error">{ error }</div>
      </form>
    )
  }

}
