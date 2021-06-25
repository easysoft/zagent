import request from '@/utils/request'

const testPrefix = '/v1/test'

const api = {
  profile: `${testPrefix}/profile`,
  plans: `${testPrefix}/plans`,
  tasks: `${testPrefix}/tasks`,

  user: `${testPrefix}/user`,
  role: `${testPrefix}/role`,
  service: `${testPrefix}/service`,
  permission: `${testPrefix}/permission`,
  permissionNoPager: `${testPrefix}/permission/no-pager`,
  orgTree: `${testPrefix}/org/tree`
}

const WebSocketPath = 'api/v1/ws'
export const WebSocketBaseDev = 'ws://127.0.0.1:8085/'
export function GetWebSocketApi () {
  const isProd = process.env.NODE_ENV === 'production'

  let wsUri = ''
  if (!isProd) {
    wsUri = WebSocketBaseDev
  } else {
    const loc = window.location

    if (loc.protocol === 'https:') {
      wsUri = 'wss:'
    } else {
      wsUri = 'ws:'
    }
    wsUri += '//' + loc.host
    wsUri += loc.pathname
  }

  return wsUri + WebSocketPath
}

export function requestSuccess (code) {
  return code === 200
}

export function getProfile (parameter) {
  return request({
    url: api.profile,
    method: 'get',
    data: parameter
  })
}

// 计划
export function listPlan (params) {
  return request({
    url: api.plans,
    method: 'get',
    params: params
  })
}
export function getPlan (id) {
  return request({
    url: api.plans + '/' + id,
    method: 'get',
    params: {}
  })
}
export function savePlan (model) {
  return request({
    url: !model.id ? api.plans : api.plans + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disablePlan (model) {
  return request({
    url: api.plans + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removePlan (model) {
  return request({
    url: api.plans + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

// 任务
export function listTask (params) {
  return request({
    url: api.tasks,
    method: 'get',
    params: params
  })
}
export function getTask (id, withIntents) {
  return request({
    url: api.tasks + '/' + id,
    method: 'get',
    params: { withIntents: withIntents }
  })
}
export function saveTask (model) {
  return request({
    url: !model.id ? api.tasks : api.tasks + '/' + model.id,
    method: !model.id ? 'post' : 'put',
    data: model
  })
}
export function disableTask (model) {
  return request({
    url: api.tasks + '/' + model.id + '/disable',
    method: 'post',
    params: {}
  })
}
export function removeTask (model) {
  return request({
    url: api.tasks + '/' + model.id,
    method: 'delete',
    params: {}
  })
}

export function getUserList (parameter) {
  return request({
    url: api.user,
    method: 'get',
    params: parameter
  })
}

export function getRoleList (parameter) {
  return request({
    url: api.role,
    method: 'get',
    params: parameter
  })
}

export function getServiceList (parameter) {
  return request({
    url: api.service,
    method: 'get',
    params: parameter
  })
}

export function getPermissions (parameter) {
  return request({
    url: api.permissionNoPager,
    method: 'get',
    params: parameter
  })
}

// id == 0 add     post
// id != 0 update  put
export function saveService (parameter) {
  return request({
    url: api.service,
    method: parameter.id === 0 ? 'post' : 'put',
    data: parameter
  })
}
