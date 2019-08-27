<template>
    <div class="bolg-body">
     <div class="article-main">
       <h2 class="article-subtitile">{{data.Title}}</h2>
       <article class="editor-body">
        {{data.Content}}
       </article>

     </div>
    </div>
</template>
<script>
import axios from 'axios'
export default {
   data () {
       return {
          data:null
       }
   },
   mounted () {
     axios.get('http://0.0.0.0:8088/getarticle/127').then((response) => {
         return response.data
     }).then((data) => {
       this.data = data.msg
       console.log(this.data)
     })
   }
}
</script>
<style lang="less" scoped>
.article-main{
  position: relative;
  width: 100%;
  background: #fff;
  overflow: hidden;
  padding: 40px 0 80px;
  .article-subtitile{
      position: relative;
      width: 720px;
      font-size: 18px;
      color: #a6a6a6;
      font-weight: 400;
      line-height: 28px;
      text-align: center;
      padding: 35px 0;
      margin: 0 auto 40px;
      box-sizing: border-box;
      &::before{
          top:0;
          position: absolute;
          content: '';
          left: 50%;
          margin-left: -13%;
          width: 26%;
          height: 2px;
          background: #eaecee;
      }
      &::after{
          bottom: 0;
          position: absolute;
            content: '';
            left: 50%;
            margin-left: -13%;
            width: 26%;
            height: 2px;
            background-color: #eaecee;
      }
  }
}
.editor-body{
    width: 720px;
    box-sizing: border-box;
    padding: 0 0 32px;
    font-size: 16px;
    line-height: 1.5;
    word-wrap: break-word;
    color:#24292e;
    margin:0 auto;
}
</style>

