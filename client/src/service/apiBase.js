import axios from 'axios';

export { doGet, doPost }

const apiPathPrefix = "/api/v1";

function doGet(url) {
    return axios.get(apiPathPrefix + url)
}

function doPost(url, data) {
    return axios.post(apiPathPrefix + url, data)
}
