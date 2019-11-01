<template>
  <div :class="{isSideShow}">
    <div class="side">
      <router-link
        class="side-link"
        v-for="(val, idx) in sideBarList"
        :key="idx"
        :to="`/manage/${val.path}`"
      >
        <span>{{val.name}}</span>
      </router-link>
    </div>
    <div class="nav">
      <div id="bs-nav-wrap">
        <div v-if="isAlertShow" class="nav-alert-info">告警提示，自动弹出</div>
        <div class="nav-bar">
          <div>
            <div @click="toggleSideEvent" class="nav-icon">
                aa
            </div>
            <div class="bold">hhh</div>
          </div>
          <div @click="toggleHalfEvent">
            <!-- <div class="time" v-if="timerange">
              <span
                v-if="!hover"
                @mouseover="hover = true"
                @mouseleave="hover = false">
                  {{ rangeText[timerange.shortcut] || timerangeText }}
              </span>
              <span
                v-if="hover"
                @mouseover="hover = true"
                @mouseleave="hover = false">
                  {{ timerangeText }}
              </span>
            </div> -->
          </div>
        </div>
      </div>
      <!-- <div class="request-loading" :class="{'request-loading-show':loading}">
        <div class="request-loading-main"></div>
      </div> -->
      <div class="rightcontent" :style="{height: `calc(100vh - ${rightContentHeight}px)`}">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>
<script>
import { SideBar } from '../utils/layout'

import { mapActions, mapGetters, mapState } from 'vuex'
export default {
  data () {
    return {
      hover: false,
      isSideShow: true,
      isAlertShow: false, // 自动展示报警信息
      isHalfInfoShow: false,
      range: null,
      sideBarList: SideBar,
      currentRouter: '',
      breadLists: [],
      timerId: null
    }
  },
  computed: {
    pathSegs () {
      return this.$route.path.split('/')
    },
    rightContentHeight () {
      return this.isHalfInfoShow ? 344 + 48 + 32 : 80
    },
    name1 () {
      return this.sideBarList.reduce((acc, cur) => {
        acc[cur.path] = cur.name
        return acc
      }, {})[this.pathSegs[1]]
    },
    algoName () {
      return this.pathSegs[2] === 'create'
        ? AlgorithmTypeName[this.pathSegs[3]]
        : this.detail
          ? this.detail.props.name
          : ''
    },
    name2 () {
      let name = {
        eventscenter: n => n,
        algorithmscenter: n => {
          let action = {
            create: '新建算法',
            edit: '编辑算法',
            model: '模型配置'
          }
          return action[n] ? `${action[n]}-${this.algoName}` : this.algoName
        }
      }
      return name[this.pathSegs[1]]
        ? name[this.pathSegs[1]](this.pathSegs[2])
        : ''
    },
    // TODO bread item color
    breadList () {
      let res = [
        {
          text: this.name1
        }
      ]
      if (this.pathSegs.length > 2) {
        res[0].link = this.pathSegs[1]
        res.push({
          text: this.name2
        })
      }
      return res
    }
  },
  watch: {
    '$route.path': function (newVal, oldVal) {
      this.currentRouter = newVal.split('/')[1]
    },
    timerange () {
      this.range = {...this.timerange}
      /** 设置定时器 */
      if (this.timerId !== null) {
        clearInterval(this.timerId)
      }
      if (this.range.shortcut) {
        this.timerId = setInterval(() => {
          this.updateTimerange({shortcut: this.range.shortcut})
        }, 1000)
      }
    }
  },
  methods: {
    ...mapActions('timerange', [
      'updateTimerange'
    ]),
    toggleSideEvent () {
      this.isSideShow = !this.isSideShow
    },
    toggleHalfEvent () {
      this.isHalfInfoShow = !this.isHalfInfoShow
      if (this.isHalfInfoShow) {
        /** 如果打开时间选取控件, 则取消定时器 */
        if (this.timerId !== null) {
          clearInterval(this.timerId)
        }
      }
    },
    /**
     * 重置面包屑导航的方法
     *  @param {*} title 及路径中需要匹配的参数，如果为undefined，说明不需要匹配，为false说明不需要面包屑导航（如设置）
     */
    resetLinks (params, _name) {
      if (params === false) return (this.breadLists = [])
      /** 根据当前的路由来进行判断处理 */
      let name = _name || this.$route.name
      if (name === 'cli') {
        this.breadLists = []
      } else {
        this.breadLists = getBreadList({ name, params })
      }
    },
    /**
     * 返回父级路由
     */
    backToParent (_name) {
      let name = _name || this.$route.name
      let parentName = getParentName(name)
      this.$router.replace({ name: parentName || 'eventscenter' })
    }
  }
}
</script>

<style lang="less" scoped>
.bold {
  font-weight: bold;
  font-size: 22px;
}

.nav {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  position: relative;
  // min-width: 1200px;
  .nav-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 48px;
    background-color: black;
    color: #fff;
    .nav-icon {
      cursor: pointer;
      fill: #fff;
      margin: 10px;
    }
    & > div {
      display: flex;
      align-items: center;
    }
    .time {
      cursor: pointer;
      margin-right: 10px;
    }
  }
  .nav-alert-info {
    width: 100%;
    height: 48px;
    background-color: #da1e28;
    text-align: center;
  }
  .nav-half {
    .nav-half-info {
      &.isAlertShow {
        height: calc(50vh - 96px);
      }
    }
  }
}
.breadcrumb {
  padding-left: 80px;
  display: flex;
  align-items: center;
  height: 32px;
  background-color: gray;
}
.rightcontent {
  overflow: auto;
}
.side {
  width: 160px;
  background-color: #fff;
  height: 100vh;
  font-size: 22px;
  position: absolute;
  left: -160px;
  color: #bababa;
  top: 0;
  .brand {
    display: flex;
    align-items: flex-end;
    margin-bottom: 32px;
    padding-top:10px;
    img {
      height: 28px;
    }
  }
  .side-link {
    display: flex;
    align-items: center;
  }
  .icon {
    // width: 20px;
    height: 30px;
    margin-right: 8px;
  }
  a {
    display: block;
    box-sizing: border-box;
    padding-left: 32px;
    height: 48px;
    line-height: 48px;
    &:hover {
      cursor: pointer;
    }
    position: relative;
  }
  .router-link-active {
    background: gray;
    &::before {
      content: '';
      display: block;
      width: 3px;
      position: absolute;
      height: 48px;
      top: 0;
      left: 0;
      background-color: blue;
    }
  }
}
.daterange-wrapper {
  padding: 20px 0;
  background: gray;
  position: relative;
  height: 100%;
  width: 100%;
  .bx--btn--primary {
    position: absolute;
    bottom: 10px;
    right: 10px;
  }
}
.isSideShow {
  .nav {
    width: calc(100vw - 160px);
    margin-left: 160px;
    height: 100vh;
    overflow: hidden;
  }
  .side {
    left: 0;
  }
}
.isAlertShow {
  &.nav-alert-info {
    display: none;
  }
}
</style>
<style lang="scss">
.brand {
  img {
    width: 100% !important;
  }
}
.breadcrumb {
  .bx--link {
    color: #0062ff;
  }
  .bx--link--current {
    color: rgba(29, 29, 29, 1) !important;
    cursor: not-allowed;
    pointer-events: none;
    -ms-touch-action: none;
    touch-action: none;
  }
  .bx--link:focus {
    box-shadow: none;
  }
}
.request-loading {
  position: absolute;
  top: 0px;
  z-index: 99;
  opacity: 1;
  animation: fade 1s linear;
  background-repeat: no-repeat;
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
  width: 100%;
  height: 48px;
  .request-loading-main {
    animation: loading 1s linear infinite;
    background: #149bdf;
    width: 100%;
    height: 100%;
    background-image: -webkit-linear-gradient(
      left,
      #3752ff,
      #2a39bf,
      #1b267f,
      #3752ff
    );
    background-size: 33.3% 10px;
  }
}
@keyframes loading {
  0% {
    background-image: -webkit-linear-gradient(
      left,
      #3752ff,
      #2a39bf,
      #1b267f,
      #3752ff
    );
    background-position: 0 0;
  }
  33% {
    background-image: -webkit-linear-gradient(
      left,
      #1b267f,
      #3752ff,
      #2a39bf,
      #1b267f
    );
    background-position: 33% 0;
  }
  66% {
    background-image: -webkit-linear-gradient(
      left,
      #2a39bf,
      #1b267f,
      #3752ff,
      #2a39bf
    );
    background-position: 66% 0;
  }
  100% {
    background-image: -webkit-linear-gradient(
      left,
      #3752ff,
      #2a39bf,
      #1b267f,
      #3752ff
    );
    background-position: 100% 0;
  }
}

.request-loading-show {
  opacity: 0;
  animation: out 1s linear;
  z-index: -1;
}
@keyframes fade {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 1;
  }
}
@keyframes out {
  0% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 0;
  }
}
</style>