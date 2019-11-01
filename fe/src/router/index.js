import Vue from 'vue'
import Router from 'vue-router'

function loadView(page) {
    return () => import(`../pages/${page}.vue`)
}

Vue.use(Router)

export default new Router ({
    linkActiveClass:'Active',
    mode:"history",
    routes:[
        {
            path:'/',
            component: loadView('home')
        },
        {
             path:'/article/:id',
             name:"article",
             component:loadView('article')
        },
        {
             path:'/login',
             name:"login",
             component:loadView('login')
        },
        {
            path:'/manage',
            component:loadView('manage/index'),
            children:[
                {
                    path:'/manage/upload',
                    name:"upload",
                    component:loadView('manage/upload')
                }
            ]
        },
        {
            path:'/notfound',
            name:'NotFound',
            component: loadView("404")
        },
        {
            path:'*',
            redirect:'/notfound'
        }
    ]
})