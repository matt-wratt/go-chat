import request from './request'

const parse = text => JSON.parse(text)

export default {
  rooms () {
    return request.get('/api/rooms').then(parse)
  },

  create (name) {
    return request.post('/api/rooms', { data: { name }})
  },

  room (room) {
    return request.get(`/api/rooms/${ room }`).then(parse)
  },

  message (room, value) {
    return request.post(`/api/rooms/${ room }`, { data: { value } })
  }
}
