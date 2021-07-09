// eslint-disable-next-line
import { UserLayout, BasicLayout } from '@/layouts'
import { dashboard, task } from '@/core/icons'

const RouteView = {
  name: 'RouteView',
  render: (h) => h('router-view')
}

export const asyncRouterMap = [

  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/test/task/list',
    children: [
      {
        path: '/platform/dashboard',
        name: 'platform-dashboard',
        component: () => import('@/views/platform/Dashboard'),
        meta: { title: 'menu.platform.dashboard', icon: dashboard, keepAlive: true }
      },
      {
        path: '/test/task',
        name: 'task',
        redirect: '/test/task/list',
        component: RouteView,
        meta: { title: 'menu.task', icon: task, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/test/task/list',
            name: 'task-list',
            component: () => import('@/views/test/task/List'),
            meta: { title: 'menu.task.list', keepAlive: true }
          },
          {
            path: '/test/task/:id/edit',
            name: 'task-edit',
            component: () => import('@/views/test/task/Edit'),
            meta: { title: 'menu.task.edit', keepAlive: true }
          },
          {
            path: '/test/task/:id/view',
            name: 'task-view',
            component: () => import('@/views/test/task/View'),
            meta: { title: 'menu.task.view', keepAlive: true }
          }
        ]
      }
    ]
  },
  {
    path: '*', redirect: '/', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  }

]
