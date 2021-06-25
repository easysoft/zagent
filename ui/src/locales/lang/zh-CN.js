import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

const locale = {
  'message': '-',
  'menu.home': '主页',

  'menu.platform': '平台',
  'menu.platform.dashboard': '仪表盘',
  'menu.settings': '设置',
  'menu.settings.sys': '系统设置',
  'menu.settings.account': '账号设置',

  'menu.task': '任务',
  'menu.task.list': '任务列表',
  'menu.task.edit': '任务编辑',

  'common.login': '登录',
  'common.logout': '登出',
  'form.view': '查看',
  'form.create': '新建',
  'form.edit': '编辑',
  'form.list': '列表',
  'form.design': '设计',
  'form.maintain': '维护',
  'form.remove': '删除',
  'form.disable': '禁用',
  'form.enable': '启用',
  'form.back': '返回',
  'form.save': '保存',
  'form.submit': '提交',
  'form.send': '发送',
  'form.reset': '重置',
  'form.cancel': '取消',
  'form.search': '查询',
  'form.collapse': '收缩',
  'form.expand': '展开',

  'form.pls.select': '请选择',
  'form.ok': '确认',
  'form.confirm.to.remove': '确认删除？',

  'form.all': '所有',
  'form.no': '序号',
  'form.code': '编码',
  'form.name': '名称',
  'form.placeholder': '占位符',
  'form.path': '路径',
  'form.content': '内容',
  'form.status': '状态',
  'form.desc': '描述',
  'form.createdBy': '创建人',
  'form.createdAt': '创建时间',
  'form.updatedAt': '更新时间',
  'form.is.default': '是否默认',
  'form.opt': '操作',
  'form.opt.log': '操作日志',

  'status.enable': '启用',
  'status.disable': '禁用',

  'valid.required.code': '请输入编码',
  'valid.required.name': '请输入名称',
  'valid.required.project': '请选择项目',
  'valid.required.path': '请输入路径',
  'valid.required.content': '请输入内容',
  'valid.select.dict': '请选择词典。',
  'valid.slot.type': '请选择类型。',
  'valid.project.path': '不是一个合法的项目路径',

  'common.status': '状态',
  'common.info': '消息',
  'common.tips': '提示',
  'common.confirm': '确认',
  'common.notify': '通知',
  'common.create': '新建',
  'common.back': '返回',

  'msg.warn': '提醒',
  'msg.confirm.to.logout': '确认退出？',
  'msg.forbidden': '确认退出？',
  'msg.unauthorized': '未授权的',
  'msg.auth.fail': '授权失败',

  'app.setting.pagestyle': '页面演示设置',
  'app.setting.pagestyle.light': '淡色',
  'app.setting.pagestyle.dark': '深色',
  'app.setting.pagestyle.realdark': '深黑',
  'app.setting.themecolor': '主题色彩',
  'app.setting.navigationmode': '导航模式',
  'app.setting.content-width': '内容宽度',
  'app.setting.fixedheader': '固定头部',
  'app.setting.fixedsidebar': '固定菜单栏',
  'app.setting.sidemenu': '左侧菜单栏布局',
  'app.setting.topmenu': '顶部菜单布局',
  'app.setting.content-width.fixed': '固定',
  'app.setting.content-width.fluid': '流式',
  'app.setting.othersettings': '其他设置',
  'app.setting.weakmode': '弱模式',
  'app.setting.copy': '复制设置',
  'app.setting.loading': '加载主题',
  'app.setting.copyinfo': '拷贝成功，请替换src/models/setting.js里的默认设置。',
  'app.setting.production.hint': '设置面板仅显示在开发模式中，请手工进行编辑。'
}

export default {
  ...components,
  ...locale
}
