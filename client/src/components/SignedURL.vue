<template>
  <div>
    <input v-model="csvFileName" size="20" />
    <button @click="downlaodCSV">CSVダウンロード</button>
  </div>
</template>

<script>
import { getURLToCSVDonwload } from "../service/signedurl";
export default {
  name: "SignedURL",
  data: function() {
    return {
      csvFileName: ""
    };
  },
  methods: {
    downlaodCSV: function() {
      getURLToCSVDonwload(this.csvFileName)
        .then(res => {
          if (res.status === 200) {
            res.text().then(data => window.open(data));
          } else {
            res.text().then(data => window.alert(data));
          }
        })
        .catch(error => window.alert(error));
    }
  }
};
</script>