<template>
    <nav class="bl-nav">
      <ul class="bl-nav-nav">
        <li>首页</li>
      <li>Github</li>
      <li>掘金</li>
      <li>归档</li>
      </ul>
      <div class="bl-nav-title">
          <h1>Stay hungry, stay foolish</h1>
          <h3>以认真的态度做完美的事情</h3>
      </div>
    </nav>
</template>
<style lang="less" scoped>
.bl-nav{
    background: linear-gradient(100deg,#3cd5ff,#007aff);
  &-title{
      text-align: center;
      padding-bottom: 120px;
      padding-top:150px;
      h1{
          font-size: 42px;
          font-weight: 700;
          margin-bottom: 20px;
          color: #ffffff;
          text-shadow: 0 1px 1px rgba(0,0,0,.2);
      }
      h3{
          font-size: 32px;
          color: #ffffff;
          text-shadow: 0 1px 1px rgba(0,0,0,.2);
      }
  }
}
.bl-nav-nav{
    width: 100%;
    display: flex;
    justify-content: space-between;
    li{
        width:100px;
        font-size: 18px;
    }
}
</style>
