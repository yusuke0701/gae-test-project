import { doGet, doPost } from './apiBase';

export { insertComment, getComment, getAllComment, updateComment }

const commentAPIBaseURL = "/comments";

function insertComment(id, body) {
    return doPost(commentAPIBaseURL, { "id": id, "body": body })
}

function getComment(id) {
    return doGet(commentAPIBaseURL + "/" + id)
}

function getAllComment() {
    return doGet(commentAPIBaseURL)
}

function updateComment(id, body) {
    return doPost(commentAPIBaseURL, { "id": id, "body": body })
}