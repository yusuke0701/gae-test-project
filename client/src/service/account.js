import { doGet, doPost } from './apiBase';

export { insertAccount, getAccount, getAllAccount, updateAccount, login, logout }

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
    return doPost(accountAPIBaseURL, account)
}

function login(id, password) {
    return doPost("/login", { "id": id, "password": password })
}

function logout(id) {
    return doPost("/logout", { "id": id })
}