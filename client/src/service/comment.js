import { doGet, doPost, doPut, doDelete } from './apiBase';

export { insertComment, getComment, getAllComment, updateComment, deleteComment }

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

function deleteComment(id) {
    return doDelete(commentAPIBaseURL + "/" + id)
}