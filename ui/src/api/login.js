import request from '@/utils/request'

const prefix = '/v1/admin'

const userApi = {
  Login: `${prefix}/login`,
  Logout: `${prefix}/logout`,
  ForgePassword: '/auth/forge-password',
  Register: '/auth/register',
  SendSms: '/account/sms',
  SendSmsErr: '/account/sms_err',

  // get my info
  UserInfo: `${prefix}/profile`,
  UserMenu: '/user/nav'
}

export function login (parameter) {
  return request({
    url: userApi.Login,
    method: 'post',
    data: parameter
  })
}

export function getSmsCaptcha (parameter) {
  return request({
    url: userApi.SendSms,
    method: 'post',
    data: parameter
  })
}

export function getInfo () {
  return request({
    url: userApi.UserInfo,
    method: 'get',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

export function getCurrentUserNav () {
  return request({
    url: userApi.UserMenu,
    method: 'get'
  })
}

export function logout () {
  return request({
    url: userApi.Logout,
    method: 'post',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}
