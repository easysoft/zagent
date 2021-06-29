export function getBuildTypes (that) {
  return { 'selenium': that.$t('form.selenium.test'), 'appium': that.$t('form.appium.test') }
}

export function getOsCategories (that) {
  return {
    'windows': that.$t('form.os.category.windows'),
    'linux': that.$t('form.os.category.linux'),
    'mac': that.$t('form.os.category.mac')
  }
}

export function getOsTypes (that) {
  return {
    'win10': that.$t('form.os.type.win10'),
    'win7': that.$t('form.os.type.win7'),
    'winxp': that.$t('form.os.type.winxp'),
    'ubuntu': that.$t('form.os.type.ubuntu'),
    'centos': that.$t('form.os.type.centos'),
    'debian': that.$t('form.os.type.debian'),
    'mac': that.$t('form.os.type.mac')
  }
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
  return {
    'en_us': that.$t('form.os.lang.en_us'),
    'zh_cn': that.$t('form.os.lang.zh_cn')
  }
}
