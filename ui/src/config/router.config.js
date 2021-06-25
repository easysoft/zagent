// eslint-disable-next-line
import { UserLayout, BasicLayout } from '@/layouts'
import { dashboard, plan } from '@/core/icons'

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
    redirect: '/test/plan/list',
    children: [
      {
        path: '/platform/dashboard',
        name: 'platform-dashboard',
        component: () => import('@/views/platform/Dashboard'),
        meta: { title: 'menu.platform.dashboard', icon: dashboard, keepAlive: true }
      },
      {
        path: '/test/plan',
        name: 'plan',
        redirect: '/test/plan/list',
        component: RouteView,
        meta: { title: 'menu.plan', icon: plan, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/test/plan/list',
            name: 'plan-list',
            component: () => import('@/views/test/plan/List'),
            meta: { title: 'menu.plan.list', keepAlive: true }
          },
          {
            path: '/test/plan/:id/edit',
            name: 'plan-edit',
            component: () => import('@/views/test/plan/Edit'),
            meta: { title: 'menu.plan.edit', keepAlive: true }
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
