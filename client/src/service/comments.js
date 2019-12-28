import axios from 'axios';

export function insertComment(id, body) {
    const apiURL = "/api/v1/comments";
    return axios.post(apiURL, { "id": id, "body": body })
}

export function getComment(id) {
    const apiURL = "/api/v1/comments";
    return axios.get(apiURL + "/" + id)
}