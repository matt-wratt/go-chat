import React from 'react'
import { Link } from 'react-router'
import Api from '../api'

export default class NewRoom extends React.Component {

  submit (event) {
    event.preventDefault()
    Api.create(this.refs.name.value).then(() => {
      this.props.history.pushState(null, '/')
    }).catch(error => {
      this.setState({ error: error.message })
    })
  }

  render () {
    this.state = this.state || {}
    const error = this.state.error || null
    return (
      <form onSubmit={ this.submit.bind(this) }>
        <h1>New Room</h1>
        <Link to="/">Back</Link>
        <br />
        <br />
        <div className="error">{ error }</div>
        <label htmlFor="name">Name</label>
        <input id="name" name="name" ref="name" />
        <input type="submit" />
      </form>
    )
  }

}
