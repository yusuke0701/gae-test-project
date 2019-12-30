import { doGet, doPost } from './apiBase';

export { insertAccount, getAccount, getAllAccount, updateAccount, login, logout }

const accountAPIBaseURL = "/accounts";

function insertAccount(id, body) {
    return doPost(accountAPIBaseURL, { "id": id, "body": body })
}

function getAccount(id) {
    return doGet(accountAPIBaseURL + "/" + id)
}

function getAllAccount() {
    return doGet(accountAPIBaseURL)
}

function updateAccount(id, body) {
    return doPost(accountAPIBaseURL, { "id": id, "body": body })
}

function login(id, password) {
    return doPost("/login", { "id": id, "password": password })
}

function logout(id) {
    return doPost("/logout", { "id": id })
}