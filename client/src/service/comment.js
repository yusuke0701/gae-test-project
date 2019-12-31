import { doGet, doPost, doPut } from './apiBase';

export { insertComment, getComment, getAllComment, updateComment }

const commentAPIBaseURL = "/comments";

function insertComment(comment) {
    return doPost(commentAPIBaseURL, comment)
}

function getComment(id) {
    return doGet(commentAPIBaseURL + "/" + id)
}

function getAllComment() {
    return doGet(commentAPIBaseURL)
}

function updateComment(comment) {
    return doPut(commentAPIBaseURL, comment)
}