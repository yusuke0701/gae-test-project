import { doGet, doPost, doPut, doDelete } from './apiBase';

export { insertAccount, getAccount, getAllAccount, updateAccount, deleteAccount, login, logout }

const accountAPIBaseURL = "/accounts";

function insertAccount(account) {
    return doPost(accountAPIBaseURL, account)
}

function getAccount(id) {
    return doGet(accountAPIBaseURL + "/" + id)
}

function getAllAccount() {
    return doGet(accountAPIBaseURL)
}

function updateAccount(account) {
    return doPut(accountAPIBaseURL, account)
}

function deleteAccount(id) {
    return doDelete(accountAPIBaseURL + "/" + id)
}

function login(id, password) {
    return doPost("/login", { "id": id, "password": password })
}

function logout(id) {
    return doPost("/logout", { "id": id })
}