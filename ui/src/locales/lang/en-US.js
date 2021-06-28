import antdEnUS from 'ant-design-vue/es/locale-provider/en_US'
import momentEU from 'moment/locale/eu'

const components = {
  antLocale: antdEnUS,
  momentName: 'eu',
  momentLocale: momentEU
}

const locale = {
  'message': '-',
  'menu.home': 'Home',

  'menu.platform': 'Platform',
  'menu.platform.dashboard': 'Dashboard',
  'menu.settings': 'Settings',
  'menu.settings.sys': 'System Settings',
  'menu.settings.account': 'Account Settings',

  'menu.task': 'Task',
  'menu.task.list': 'Task List',
  'menu.task.edit': 'Task Edit',

  'form.view': 'View',
  'form.create': 'Create',
  'form.edit': 'Edit',
  'form.list': 'List',
  'form.design': 'Design',
  'form.maintain': 'Maintain',
  'form.remove': 'Delete',
  'form.disable': 'Disable',
  'form.enable': 'Enable',
  'form.back': 'Back',
  'form.save': 'Save',
  'form.submit': 'Submit',
  'form.send': 'Send',
  'form.reset': 'Reset',
  'form.confirm': 'Confirm',
  'form.cancel': 'Cancel',
  'form.search': 'Search',
  'form.collapse': 'Collapse',
  'form.expand': 'Expand',

  'form.pls.select': 'Please Select',
  'form.ok': 'Ok',
  'form.confirm.to.remove': 'Confirm to delete?',

  'form.all': 'All',
  'form.no': 'NO',
  'form.code': 'Code',
  'form.name': 'Name',
  'form.placeholder': 'Placeholder',
  'form.type': 'Type',
  'form.path': 'Path',
  'form.content': 'Content',
  'form.status': 'Status',
  'form.desc': 'Description',
  'form.createdBy': 'Created By',
  'form.createdAt': 'Created Time',
  'form.updatedAt': 'Updated Time',
  'form.is.default': 'Is Default',
  'form.opt': 'Operation',
  'form.opt.log': 'Operation Log',

  'status.enable': 'Enable',
  'status.disable': 'Disable',

  'valid.required.code': 'Please input code.',
  'valid.required.name': 'Please input name.',
  'valid.required.buildType': 'Please select type.',

  'common.status': '状态',
  'common.login': 'Login',
  'common.logout': 'Logout',
  'common.info': 'Info',
  'common.tips': 'Tips',
  'common.confirm': 'Confirmation',
  'common.notify': 'Notification',
  'common.create': 'Create',
  'common.back': 'Back',

  'msg.warn': 'Warning',
  'msg.confirm.to.logout': 'Do you really log-out.',
  'msg.forbidden': 'Forbidden',
  'msg.unauthorized': 'Unauthorized',
  'msg.auth.fail': 'Authorization verification failed',

  'app.setting.pagestyle': 'Page style setting',
  'app.setting.pagestyle.light': 'Light style',
  'app.setting.pagestyle.dark': 'Dark style',
  'app.setting.pagestyle.realdark': 'RealDark style',
  'app.setting.themecolor': 'Theme Color',
  'app.setting.navigationmode': 'Navigation Mode',
  'app.setting.content-width': 'Content Width',
  'app.setting.fixedheader': 'Fixed Header',
  'app.setting.fixedsidebar': 'Fixed Sidebar',
  'app.setting.sidemenu': 'Side Menu Layout',
  'app.setting.topmenu': 'Top Menu Layout',
  'app.setting.content-width.fixed': 'Fixed',
  'app.setting.content-width.fluid': 'Fluid',
  'app.setting.othersettings': 'Other Settings',
  'app.setting.weakmode': 'Weak Mode',
  'app.setting.copy': 'Copy Setting',
  'app.setting.loading': 'Loading theme',
  'app.setting.copyinfo': 'copy success，please replace defaultSettings in src/models/setting.js',
  'app.setting.production.hint': 'Setting panel shows in development environment only, please manually modify'
}

export default {
  ...components,
  ...locale
}
