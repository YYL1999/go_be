<template>
       <div id="app">
         <home-header></home-header>
         <router-view/>
         <home-footer></home-footer>
       </div>
</template>
<script>
import HomeHeader from './components/Header'
import HomeFooter from './components/Footer'
export default {
    name:'App',
    components: {
      HomeHeader,
      HomeFooter
    }
}
</script>

<style lang="less" >
footer{
  text-align: center;
  font-size: 16px;
  border-top: 1px solid #e7eaf1;
  min-height: 40px;
  padding:20px 0;
  color: #a6abb2;
  background: #f7f8fa
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 1s;
}
.fade-enter,
.fade-leave-active {
  opacity: 0;
}

body {
  margin: 0;
  background-color: #000;
  font-family: system,
  "webfont", "PingFang-SC", "Segoe UI", "Microsoft YaHei", "wenquanyi micro hei","Hiragino Sans GB", "Hiragino Sans GB W3", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans",
  "Helvetica Neue", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale; 
  background-color: #f7f8fa;
  color: #656565;
}
a {
  color: #000;
  text-decoration: none;
}
ul {
  margin: 0;
  padding-left: 0;
  li {
    list-style: none;
    text-align: left;
  }
}
p {
  margin: 0;
}
h1,h2,h3,h4,h5,h6{
  margin: 0;
}

</style>
