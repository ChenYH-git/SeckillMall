<template>
  <div>
    <router-view/>

    <div class="footer">
      CYH
    </div>
  
  </div>
  

</template>

<script>


export default {
  name: 'App',
  components: {
  },
}
</script>

<style lang="scss">
@import "static/css/init.css";
@import "static/css/common.css";

#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50; */
  /* margin-top: 60px; */
}

.footer {
  height: 80px;
  line-height: 80px;
  text-align: center;
  background: #333;
  color: #fff;
  font-size: 12px;
  letter-spacing: 1px;
  /* margin-top: 110px; */
  margin-top: 110px;
}
</style>
