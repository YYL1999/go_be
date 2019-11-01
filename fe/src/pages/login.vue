<template>
  <div class="login">
    <div class="login-main">
      <input type="text"  placeholder="请输入用户名" v-model="name" class="login-input">
      <input type="password"  placeholder="请输入密码" v-model="password" class="login-input">
      <button @click="handleClick" class="login-button">登录</button>
    </div>
  </div>
</template>
<script>
export default {
    data() {
        return {
            name:null,
            password:null
        }
    },
    methods: {
        handleClick () {
            console.log(this.name,this.password)
        }
    }
}
</script>

<style lang="less" scoped>
.login {
  width: 100%;
  height: calc(100vh - 300px);
  background: #fafafa;
  display: flex;
  justify-content: center;
  align-items: center;
  &-main{
      width: 500px;
      height: 500px;

  }
}
.login-input{
    -webkit-appearance: none;
    background-color: #fff;
    background-image: none;
    border-radius: 4px;
    border: 1px solid #dcdfe6;
    box-sizing: border-box;
    color: #606266;
    display: inline-block;
    font-size: inherit;
    height: 40px;
    line-height: 40px;
    font-size: 22px;
    outline: none;
    padding: 0 15px;
    transition: border-color .2s cubic-bezier(.645,.045,.355,1);
    width: 100%;
}
.login-button{
    display: inline-block;
    line-height: 1;
    white-space: nowrap;
    cursor: pointer;
    background: #fff;
    border: 1px solid #dcdfe6;
    -webkit-appearance: none;
    text-align: center;
    box-sizing: border-box;
    outline: none;
    margin: 0;
    transition: .1s;
    font-weight: 500;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    padding: 12px 20px;
    font-size:40px;
    border-radius: 4px;
    color: #fff;
    background-color: #409eff;
    border-color: #409eff;
}
</style>

