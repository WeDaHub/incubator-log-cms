<template>
<div id="vue_det">
    <h1>site : {{site}}</h1>
    <h1 @click="details()">url : {{url}}</h1>
</div>
</template>

<script type="text/javascript">
import Vue from 'vue'
   export default
   {
        data () {
            return {
            site: '菜鸟教程',
            url: 'www.runoob.com',
            alexa: '10000'
            }
        },
        methods: {
            details () {

            }

        }
    }
</script>
