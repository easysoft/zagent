// eslint-disable-next-line
import { UserLayout, BasicLayout } from '@/layouts'
import { dashboard, project, task, lookup, synonym, regex, setting } from '@/core/icons'

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
    redirect: '/project/list',
    children: [
      {
        path: '/platform/dashboard',
        name: 'platform-dashboard',
        component: () => import('@/views/platform/Dashboard'),
        meta: { title: 'menu.platform.dashboard', icon: dashboard, keepAlive: true }
      },
      {
        path: '/project',
        name: 'project',
        redirect: '/project/list',
        component: RouteView,
        meta: { title: 'menu.project', icon: project, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/project/list',
            name: 'project-list',
            component: () => import('@/views/project/List'),
            meta: { title: 'menu.project.list', keepAlive: true }
          },
          {
            path: '/project/:id/edit',
            name: 'project-edit',
            component: () => import('@/views/project/Edit'),
            meta: { title: 'menu.project.edit', keepAlive: true }
          },
          {
            path: '/project/:id/view',
            name: 'project-view',
            component: () => import('@/views/project/View'),
            meta: { title: 'menu.project.view', keepAlive: true }
          },
          {
            path: '/project/:id/test',
            name: 'project-view',
            component: () => import('@/views/project/Test'),
            meta: { title: 'menu.project.test', keepAlive: true }
          }
        ]
      },
      {
        path: '/nlu/task',
        name: 'task',
        redirect: '/nlu/task/list',
        component: RouteView,
        meta: { title: 'menu.task', icon: task, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/nlu/task/list',
            name: 'task-list',
            component: () => import('@/views/nlu/task/List'),
            meta: { title: 'menu.task.list', keepAlive: true }
          },
          {
            path: '/nlu/task/:id/edit',
            name: 'task-edit',
            component: () => import('@/views/nlu/task/Edit'),
            meta: { title: 'menu.task.edit', keepAlive: true }
          }
        ]
      },
      {
        path: '/nlu/synonym',
        name: 'synonym',
        redirect: '/nlu/synonym/list',
        component: RouteView,
        meta: { title: 'menu.synonym', icon: synonym, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/nlu/synonym/list',
            name: 'synonym-list',
            component: () => import('@/views/nlu/synonym/List'),
            meta: { title: 'menu.synonym.list', keepAlive: true }
          },
          {
            path: '/nlu/synonym/:id/edit',
            name: 'synonym-edit',
            component: () => import('@/views/nlu/synonym/Edit'),
            meta: { title: 'menu.synonym.edit', keepAlive: true }
          },
          {
            path: '/nlu/synonym/:synonymId/items',
            name: 'synonym-maintain',
            component: () => import('@/views/nlu/synonym/item/List'),
            meta: { title: 'menu.synonym.items', keepAlive: true }
          }
        ]
      },
      {
        path: '/nlu/lookup',
        name: 'lookup',
        redirect: '/nlu/lookup/list',
        component: RouteView,
        meta: { title: 'menu.lookup', icon: lookup, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/nlu/lookup/list',
            name: 'lookup-list',
            component: () => import('@/views/nlu/lookup/List'),
            meta: { title: 'menu.lookup.list', keepAlive: true }
          },
          {
            path: '/nlu/lookup/:id/edit',
            name: 'lookup-edit',
            component: () => import('@/views/nlu/lookup/Edit'),
            meta: { title: 'menu.lookup.edit', keepAlive: true }
          },
          {
            path: '/nlu/lookup/:lookupId/items',
            name: 'lookup-maintain',
            component: () => import('@/views/nlu/lookup/item/List'),
            meta: { title: 'menu.lookup.items', keepAlive: true }
          }
        ]
      },
      {
        path: '/nlu/regex',
        name: 'regex',
        redirect: '/nlu/regex/list',
        component: RouteView,
        meta: { title: 'menu.regex', icon: regex, keepAlive: true },
        hideChildrenInMenu: true,
        children: [
          {
            path: '/nlu/regex/list',
            name: 'regex-list',
            component: () => import('@/views/nlu/regex/List'),
            meta: { title: 'menu.regex.list', keepAlive: true }
          },
          {
            path: '/nlu/regex/:id/edit',
            name: 'regex-edit',
            component: () => import('@/views/nlu/regex/Edit'),
            meta: { title: 'menu.regex.edit', keepAlive: true }
          },
          {
            path: '/nlu/regex/:regexId/items',
            name: 'regex-maintain',
            component: () => import('@/views/nlu/regex/item/List'),
            meta: { title: 'menu.regex.items', keepAlive: true }
          }
        ]
      },
      {
        path: '/settings/sys',
        name: 'settings-sys',
        component: () => import('@/views/settings/sys/Index'),
        meta: { title: 'menu.settings.sys', icon: setting, keepAlive: true }
      },
      {
        path: '/settings/account',
        name: 'settings-account',
        hidden: true,
        component: () => import('@/views/settings/account/Index'),
        meta: { title: 'menu.settings.account', keepAlive: true }
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
