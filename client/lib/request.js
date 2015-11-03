function request(method, url, opts = {}) {
  return new Promise((resolve, reject) => {
    const headers = opts.headers || {}
    const data = opts.data || {}
    const xhr = new XMLHttpRequest
    xhr.open(method, url)
    Object.keys(headers).forEach(header => {
      xhr.setRequestHeader(header, headers[header])
    })
    xhr.onload = event => resolve(xhr.responseText)
    xhr.onerror = reject
    if (data) {
      xhr.send(new Blob([JSON.stringify(data)], {type : 'application/json'}))
    } else {
      xhr.send()
    }
  })
}

export default {
  get: request.bind(null, 'GET'),
  post: request.bind(null, 'POST'),
  put: request.bind(null, 'PUT'),
}
