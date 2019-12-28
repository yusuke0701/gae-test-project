import axios from 'axios';

export function insertComment(id, body) {
    const apiURL = "/api/comments";
    return axios.post(apiURL, { "id": id, "body": body })
}

export function getComment(id) {
    const apiURL = "/api/comments";
    return axios.get(apiURL + "/" + id)
}