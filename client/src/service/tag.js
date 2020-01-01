import { doGet, doPost, doPut, doDelete } from './apiBase';

export { insertTag, getTag, getAllTag, updateTag, deleteTag }

const tagAPIBaseURL = "/tags";

function insertTag(tag) {
    return doPost(tagAPIBaseURL, tag)
}

function getTag(id) {
    return doGet(tagAPIBaseURL + "/" + id)
}

function getAllTag() {
    return doGet(tagAPIBaseURL)
}

function updateTag(tag) {
    return doPut(tagAPIBaseURL, tag)
}

function deleteTag(id) {
    return doDelete(tagAPIBaseURL + "/" + id)
}