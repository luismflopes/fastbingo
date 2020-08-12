import axios from 'axios'

const _axios = axios.create({
  baseURL: 'http://localhost:8081/api/v1',
  withCredentials: false,
  crossdomain: true,
  headers: {
  }
})

export default _axios
