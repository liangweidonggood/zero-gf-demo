<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  created() {
    this.test();
  },
  methods:{
    test(){
      console.log(123)
      // 组件中使用post方法
      this.axios.post('/api/api/v1/user/one',{"id":2}
      ).then(res => {

        // 成功回调
        console.log(res)
      }, res => {
        // 错误回调
        console.log(45678)
        console.log(res)
      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
