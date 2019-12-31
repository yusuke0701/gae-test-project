import axios from 'axios';

export { doGet, doPost, doPut, doDelete }

const apiOrigin = process.env.VUE_APP_API_ORIGIN
const apiPathPrefix = '/api/v1';

const axiosInstance = axios.create({
    baseURL: apiOrigin + apiPathPrefix
})

function doGet(url) {
    return axiosInstance.get(url)
}

function doPost(url, data) {
    return axiosInstance.post(url, data)
}

function doPut(url, data) {
    return axiosInstance.put(url, data)
}

function doDelete(url) {
    return axiosInstance.delete(url)
}