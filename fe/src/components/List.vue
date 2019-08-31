<template>
  <div class="bl-list" @click="handleClick">
        <h1>{{item.Title}}</h1> 
        <p>{{item.Content}}</p>
  </div>
</template>
<script>
export default {
  props: {
    item: {
      type: Object
    }
  },
  methods: {
    handleClick () {
      this.$router.push(`article/${this.item.Id}`)
    }
  }
};
</script>

<style lang="less" scoped>
.bl-list{
    width: 100%;
    margin-bottom: 16px;
    padding:30px 20px;
    background: #ffffff;
    border: 1px solid #e7eaf1;
    border-radius: 10px;
    box-sizing: border-box;
    box-shadow: 0 1px 3px rgba(0, 37, 55, .06);
    h1{
        font-size:20px;
    }
    p{
        font-size: 16px;
        padding-top: 5px;
        height: 24px;
        overflow: hidden;
    }
}
</style>
