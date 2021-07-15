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
  'menu.task.view': '任务查看',

  'form.uint.ge': '个',
  'form.view': '查看',
  'form.add': '添加',
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
  'form.step': '步骤',
  'form.result': '结果',
  'form.result.down': '下载结果',
  'form.vnc.url': '远程桌面',
  'form.time': '时间',
  'form.code': '编码',
  'form.name': '名称',
  'form.placeholder': '占位符',
  'form.type': '类型',
  'form.content': '内容',
  'form.progress': '进度',
  'form.status': '状态',
  'form.desc': '描述',
  'form.createdBy': '创建人',
  'form.createdAt': '创建时间',
  'form.updatedAt': '更新时间',
  'form.pendingAt': '延期时间',
  'form.startAt': '开始时间',
  'form.completeAt': '完成时间',
  'form.is.default': '是否默认',
  'form.opt': '操作',
  'form.opt.log': '操作日志',

  'form.test.env': '测试环境',
  'form.group': '分组ID',
  'form.group.tips': '用于分组显示',
  'form.driver.type': '驱动类型',
  'form.driver.type.tips': '如chrome、firefox、edge和ie。',
  'form.driver.version': '驱动版本',
  'form.driver.version.tips': '主版本，如92。',
  'form.test.code': '测试代码',
  'form.test.code.tips': '输入项目Git仓库或Zip文件的下载地址。',
  'form.scm.account': 'Git账号',
  'form.scm.password': 'Git账号密码',
  'form.exec.cmd': '测试命令',
  'form.exec.cmd.tips': '输入执行测试的命令，可引用环境变量。',
  'form.env.var': '环境变量',
  'form.env.var.tips': '需要传递的环境变量，格式"变量名=取值"，支持多行。',
  'form.result.files': '结果文件',
  'form.result.files.tips': '列出需要打包的测试结果文件，支持多行。',
  'form.test.type': '测试类型',
  'form.selenium': 'Selenium',
  'form.appium': 'Appium',
  'form.selenium.test': 'Selenium测试',
  'form.appium.test': 'Appium测试',

  'form.os.category': '系统分类',
  'form.os.category.windows': 'Windows',
  'form.os.category.linux': 'Linux',
  'form.os.category.mac': 'Mac',

  'form.os.type': '系统类型',
  'form.os.type.win10': 'Win10',
  'form.os.type.win7': 'Win7',
  'form.os.type.winxp': 'WinXP',
  'form.os.type.ubuntu': 'Ubuntu',
  'form.os.type.centos': 'Centos',
  'form.os.type.debian': 'Debian',
  'form.os.type.mac': 'Mac',

  'form.os.lang': '系统语言',
  'form.os.lang.en_us': '美国英语',
  'form.os.lang.zh_cn': '简体中文',

  'form.edit.env': '编辑环境',

  'history.type.task': '任务',
  'history.type.queue': '队列',
  'history.type.build': '构建',
  'history.type.vm': '虚机',

  'build.init': '初始化',
  'build.progress.start': '开始',
  'build.progress.res': '准备',
  'build.progress.exec': '执行',
  'build.progress.end': '结束',

  'build.progress.created': '新建',
  'build.progress.pending_res': '等待资源',
  'build.progress.launch_vm': '启动虚机',
  'build.progress.create_vm_fail': '创建虚机失败',
  'build.progress.perform_request_fail': '执行请求失败',
  'build.progress.appium_service_fail': 'Appium服务失败',
  'build.progress.running': '执行中',
  'build.progress.timeout': '超时',
  'build.progress.completed': '完成',
  'build.progress.cancel': '取消',

  'build.status.created': '新建',
  'build.status.pass': '通过',
  'build.status.fail': '失败',

  'vm.status.created': '新建',
  'vm.status.launch': '启动',
  'vm.status.vm_fail_create': '虚机创建失败',
  'vm.status.running': '运行中',
  'vm.status.shutoff': '关闭',
  'vm.status.destroy': '摧毁',
  'vm.status.vm_fail_destroy': '摧毁失败',
  'vm.status.busy': '繁忙',
  'vm.status.ready': '就绪',
  'vm.status.unknown': '未知',

  'status.enable': '启用',
  'status.disable': '禁用',

  'valid.required.code': '请输入编码',
  'valid.required.name': '请输入名称',
  'valid.required.buildType': '请选择类型',
  'valid.required.osCategory': '请选择系统分类',
  'valid.required.osType': '请选择系统类型',
  'valid.required.osVersion': '请选择系统版本',
  'valid.required.osLang': '请选择系统语言',
  'valid.required.scriptUrl': '请输入测试项目的地址',
  'valid.required.buildCommands': '构建命令不能为空',
  'valid.required.resultFiles': '结果文件不能为空',

  'common.login': '登录',
  'common.logout': '登出',
  'common.detail': '详情',
  'common.request': '请求',
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
