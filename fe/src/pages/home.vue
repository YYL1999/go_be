<template>
  <div class="bolg">
    <main class="bl-main">
      <section v-for="(item,index) in list" :key="index">
        <bl-list :item="item"></bl-list>
      </section>
      <div class="bl-fy">
        <div class="bl-fy-main">
          <section>
            <div class="avatar">
              <img src="https://jingwei.link/assets/img/profile.png">
            </div>
            <h3>虚竹</h3>
            <p>Stay hungry, stay foolish</p>
            <ul class="bl-fy-main-a">
              <li>a</li>
              <li>a</li>
              <li>a</li>
            </ul>
          </section>
          <section class="tag-card">
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>

            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
          </section>
        </div>
        <div class="bl-fy-main">
          <h3>友情链接</h3>
          <section class="tag-card">
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
            <a href="baidu.com">百度</a>
          </section>
        </div>
        <div class="search-card">
          <input class="search-input" type="text" placeholder="Search···">
        </div>
      </div>
    </main>
  </div>
</template>
<script>
import { mapActions,mapState } from 'vuex'
import {getAllArticle} from '../api/article'
import axios from 'axios'
export default {
  data() {
    return {
      list: [
        {
          title: "总共话",
          message: "xsxsxs"
        },
        {
          title: "总共话",
          message: "xsxsxs"
        },
        {
          title: "总共话",
          message: "xsxsxs"
        },
        {
          title: "总共话",
          message: "xsxsxs"
        },
        {
          title: "总共话",
          message: "xsxsxs"
        },
         {
          title: "总共话",
          message: "xsxsxs"
        },
         {
          title: "总共话",
          message: "xsxsxs"
        },
         {
          title: "总共话",
          message: "xsxsxs"
        }
      ]
    };
  },
  computed: {
    // ...mapState(["list"])
  },
  methods : {
    getArticle() {
       axios.get('http://0.0.0.0:8088/getallarticle'
       ).then((response)=> {
         return response.data
       }).then((data) => {
        this.list = data.msg
        console.log(data)
       })
       
    }
  },
  mounted () {
    this.getArticle()
  }
};
</script>
<style lang="less" scoped>
.bolg {
  background-color: #f7f8fa;
  padding-top: 100px;
  overflow: hidden;
}
.bl-main {
  width: 40%;
  margin: 40px auto 60px;
  position: relative;
}
.bl-fy {
  position: absolute;
  top: 0;
  left: 700px;
  width: 240px;
  &-main {
    padding: 30px 20px;
    background-color: #fff;
    border: 1px solid #e7eaf1;
    border-radius: 3px;
    box-sizing: border-box;
    box-shadow: 0 1px 3px rgba(0, 37, 55, 0.06);
    section {
      padding: 30px 0;
      border-bottom: 1px solid #e7eaf1;
    }
    .tag-card {
      border-bottom: none;
      padding-bottom: 0;
      font-size: 14px;
      a {
        display: inline-block;
        width: auto;
        height: 22px;
        padding: 0 8px;
        color: #a6abb2;
        line-height: 22px;
        box-sizing: border-box;
        box-shadow: 0 0 0.6px #d5dae1;
        border-radius: 2px;
        background-color: #fafafa;
        cursor: pointer;
        margin-right: 2px;
        transition: 0.2s;
        margin-bottom: 8px;
      }
    }

    .avatar {
      width: 70px;
      height: 70px;
      border-radius: 50%;
      margin: 0 auto;
      overflow: hidden;
      box-shadow: 0 1px 4px rgba(100, 110, 120, 0.53);
      img {
        width: 70px;
      }
    }
    h3 {
      width: 100%;
      font-size: 16px;
      color: #4e4e4e;
      font-weight: 500;
      margin: 12px 0 6px;
      text-align: center;
    }
    p {
      font-size: 14px;
      line-height: 20px;
      text-align: center;
    }
    &-a {
      margin-top: 14px;
      height: 100%;
      display: flex;
      width: 60%;
      margin: 0 auto;
      justify-content: space-around;
      line-height: 24px;
      li {
        position: relative;
        display: inline-block;
        width: 26px;
        height: 26px;
        font-size: 14px;
      }
    }
  }
}
.search-card {
  position: relative;
  margin-top: 16px;
  padding: 0;
  font-size: 21px;
  .search-input {
    border: 1px solid #e7eaf1;
    border-radius: 3px;
    box-sizing: border-box;
    box-shadow: 0 1px 3px rgba(0, 37, 55, 0.06);
    position: relative;
    padding: 12px 46px 12px 18px;
    width: 100%;
    background: #fff;
    outline: 0;
    z-index: 1;
    font-size: 16px;
    color: #4e4e4e;
  }
}
</style>

