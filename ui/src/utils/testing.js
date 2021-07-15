import { buildProgressEnd, buildProgressExec, buildProgressPrepareRes, buildProgressStart }
  from '@/utils/const'
import { WebBaseDev } from '@/api/manage'

export function getTaskProgressMap (histories, buildProgress) {
  const taskProgressMap = {}
  histories.forEach((item, index) => {
    let key = ''
    if (buildProgressStart.indexOf(item.progress) > -1) {
      key = 'start'
    } else if (buildProgressPrepareRes.indexOf(item.progress) > -1) {
      key = 'res'
    } else if (buildProgressExec.indexOf(item.progress) > -1) {
      key = 'exec'
    } else if (buildProgressEnd.indexOf(item.progress) > -1) {
      key = 'end'
    }

    taskProgressMap[key] = {}
    taskProgressMap[key].status = buildProgress[item.progress]
    taskProgressMap[key].time = item.createdAt
  })

  return taskProgressMap
}

export function getTaskBuildHistories (buildHistories, that) {
  const ret = {}

  const buildProgress = getBuildProgress(that)
  const buildStatus = getBuildStatus(that)

  buildHistories.forEach((item, index) => {
    if (!(item.queueId in ret)) {
      ret[item.queueId] = []
    }

    ret[item.queueId].push(
      {
        key: item.id,
        progress: buildProgress[item.progress],
        status: buildStatus[item.status],
        time: item.createdAt,
        resultUrl: WebBaseDev + item.resultPath,
        vncUrl: 'http://' + item.nodeIp + ':' + item.vncPort // TODO: use novnc address
      }
    )
  })

  return ret
}

export function getBuildProgress (that) {
  return {
    'created': that.$t('build.progress.created'),
    'pending_res': that.$t('build.progress.pending_res'),
    'launch_vm': that.$t('build.progress.launch_vm'),
    'create_vm_fail': that.$t('build.progress.create_vm_fail'),
    'perform_request_fail': that.$t('build.progress.perform_request_fail'),
    'appium_service_fail': that.$t('build.progress.appium_service_fail'),
    'running': that.$t('build.progress.running'),
    'timeout': that.$t('build.progress.timeout'),
    'completed': that.$t('build.progress.completed'),
    'cancel': that.$t('build.progress.cancel')
  }
}
export function getBuildStatus (that) {
  return {
    'created': that.$t('build.status.created'),
    'pass': that.$t('build.status.pass'),
    'fail': that.$t('build.status.fail')
  }
}
export function getVmStatus (that) {
  return {
    'created': that.$t('vm.status.created'),
    'launch': that.$t('vm.status.launch'),
    'vm_fail_create': that.$t('vm.status.vm_fail_create'),

    'running': that.$t('vm.status.running'),
    'shutoff': that.$t('vm.status.shutoff'),
    'destroy': that.$t('vm.status.destroy'),
    'vm_fail_destroy': that.$t('vm.status.vm_fail_destroy'),

    'busy': that.$t('vm.status.busy'),
    'ready': that.$t('vm.status.ready'),

    'unknown': that.$t('vm.status.unknown')
  }
}

export function getBuildTypes (that) {
  return {
    'selenium': that.$t('form.selenium'),
    'appium': that.$t('form.appium') }
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
