import axios from 'axios';

export function getURLToCSVDonwload(csvFileName) {
    const apiURL = "/api/v1/url/csv-download";
    return axios
        .get(apiURL + "/" + csvFileName)
}