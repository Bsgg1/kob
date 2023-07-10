<template>
  <div>
    <div>name:{{ bot_name }}</div>
    <div>id: {{ bot_id }}</div>
  </div>
</template>
<script>
import $ from "jquery";
import {ref} from "vue";

export default{
  name:"app",
  setup: () => {
    let bot_name = ref("")
    let bot_id = ref("")

    $.ajax({
      url: "http://localhost:8081/bot/getInfo",
      type: "get",
      success: resp => {
        bot_name.value = resp.name,
        bot_id.value = resp.id
      }
    })
    return {
      bot_name,
      bot_id
    }
  }
}

</script>
<style>

</style>
