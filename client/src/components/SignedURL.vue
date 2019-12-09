<template>
  <div>
    <input v-model="csvFileName" size="20" />
    <button v-on:click="downlaodCSV">CSVダウンロード</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "SignedURL",
  data: function() {
    return {
      csvFileName: ""
    };
  },
  methods: {
    downlaodCSV: function() {
      const apiURL = "/api/url/csv-download";
      axios
        .get(apiURL + "/" + this.csvFileName)
        .then(response => response.text())
        .then(textData => window.open(textData))
        .catch(error => console.log(error));
    }
  }
};
</script>