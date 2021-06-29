export function getBuildTypes (that) {
  const m = new Map()

  m.set('selenium', that.$t('form.selenium.test'))
  m.set('appium', that.$t('form.appium.test'))

  return m
}

export function getOsCategories (that) {
  const m = new Map()

  m.set('windows', that.$t('form.os.category.windows'))
  m.set('linux', that.$t('form.os.category.linux'))
  m.set('mac', that.$t('form.os.category.mac'))

  return m
}

export function getOsTypes (that) {
  // return { 'windows': getWindowsMap(that), 'linux': getLinuxMap(that), 'mac': getMacMap(that) }

  const m = new Map()

  m.set('win10', that.$t('form.os.type.win10'))
  m.set('win7', that.$t('form.os.type.win7'))
  m.set('winxp', that.$t('form.os.type.winxp'))

  m.set('ubuntu', that.$t('form.os.type.ubuntu'))
  m.set('centos', that.$t('form.os.type.centos'))
  m.set('debian', that.$t('form.os.type.debian'))

  m.set('mac', that.$t('form.os.type.mac'))

  return m
}

/* function getWindowsMap (that) {
  const m = new Map()

  m.set('win10', that.$t('form.os.type.win10'))
  m.set('win7', that.$t('form.os.type.win7'))
  m.set('winxp', that.$t('form.os.type.winxp'))

  return m
}
function getLinuxMap (that) {
  const m = new Map()

  m.set('ubuntu', that.$t('form.os.type.ubuntu'))
  m.set('centos', that.$t('form.os.type.centos'))
  m.set('debian', that.$t('form.os.type.debian'))

  return m
}
function getMacMap (that) {
  const m = new Map()

  m.set('mac', that.$t('form.os.type.mac'))

  return m
} */

export function getOsVersion (that) {
  return []
}

export function getOsLangs (that) {
  const m = new Map()

  m.set('en_us', that.$t('form.os.lang.en_us'))
  m.set('zh_cn', that.$t('form.os.lang.zh_cn'))

  return m
}
