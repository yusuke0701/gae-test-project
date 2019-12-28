import axios from 'axios';

export function getURLToCSVDonwload(csvFileName) {
    const apiURL = "/api/url/csv-download";
    return axios
        .get(apiURL + "/" + csvFileName)
}