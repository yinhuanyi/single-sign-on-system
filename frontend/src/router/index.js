import { createRouter, createWebHashHistory } from 'vue-router'
import layout from '@/layout'

const privateRoutes = [
  {
    path: '/user',
    component: layout,
    meta: { title: 'user', icon: 'personnel' },
    redirect: '/user/manage',
    children: [
      {
        path: '/user/manage',
        component: () => import('@/views/user-manage'),
        meta: {
          title: 'userManage',
          icon: 'personnel-manage'
        }
      },
      {
        path: '/user/role',
        component: () => import('@/views/role-list'),
        meta: {
          title: 'roleList',
          icon: 'role'
        }
      },
      {
        path: '/user/permission',
        component: () => import('@/views/permission-list'),
        meta: {
          title: 'permissionList',
          icon: 'permission'
        }
      },
      {
        path: '/user/info/:id', // 这里是动态路径，在组件的js中可以通过route获取到id的值
        name: 'userInfo',
        component: () => import('@/views/user-info'),
        meta: {
          title: 'userInfo'
        }
      },
      {
        path: '/user/import',
        name: 'import',
        component: () => import('@/views/import'),
        meta: {
          title: 'excelImport'
        }
      }
    ]
  },
  {
    path: '/article',
    component: layout,
    redirect: '/article/ranking',
    meta: {
      title: 'article',
      icon: 'article'
    },
    children: [
      {
        path: '/article/ranking',
        component: () => import('@/views/article-ranking/index'),
        name: 'articleRanking',
        meta: {
          title: 'articleRanking',
          icon: 'article-ranking'
        }
      },
      {
        path: '/article/:id',
        component: () => import('@/views/article-detail/index'),
        name: 'articleDetail',
        meta: {
          title: 'articleDetail'
        }
      },
      {
        path: '/article/create',
        component: () => import('@/views/article-create/index'),
        name: 'articleCreate',
        meta: {
          title: 'articleCreate',
          icon: 'article-create'
        }
      },
      {
        path: '/article/editor/:id',
        component: () => import('@/views/article-create/index'),
        name: 'articleEditor',
        meta: {
          title: 'articleEditor'
        }
      }
    ]
  }
]

const publicRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login')
  },
  {
    path: '/',
    component: layout,
    redirect: '/profile',
    children: [
      {
        path: '/profile',
        name: 'profile',
        component: () => import('@/views/profile'),
        // 这里将个人中心作为菜单栏的第一个菜单
        meta: { title: 'profile', icon: 'el-icon-user' }
      },
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('@/views/dashboard'),
        meta: { title: 'dashboard', icon: 'el-icon-user' }
      },
      {
        path: '/404',
        name: '404',
        component: () => import('@/views/error-page/404')
      },
      {
        path: '/401',
        name: '401',
        component: () => import('@/views/error-page/401')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  // 这里将共有路由和私有路由合并
  routes: [...publicRoutes, ...privateRoutes]
})

export default router
