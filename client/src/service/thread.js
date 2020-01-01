import { doGet, doPost, doPut, doDelete } from './apiBase';

export { insertThread, getThread, getAllThread, updateThread, deleteThread }

const threadAPIBaseURL = "/threads";

function insertThread(thread) {
    return doPost(threadAPIBaseURL, thread)
}

function getThread(id) {
    return doGet(threadAPIBaseURL + "/" + id)
}

function getAllThread() {
    return doGet(threadAPIBaseURL)
}

function updateThread(thread) {
    return doPut(threadAPIBaseURL, thread)
}

function deleteThread(id) {
    return doDelete(threadAPIBaseURL + "/" + id)
}