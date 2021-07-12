
const TEST_HISTORIES = 'Test-Histories'

const colsFull = 24
const wrapperColFull = { lg: { span: 16 }, sm: { span: 16 } }
const labelColFull = { lg: { span: 4 }, sm: { span: 4 } }

const colsHalf = 12
const wrapperColHalf = { lg: { span: 12 }, sm: { span: 12 } }
const labelColHalf = { lg: { span: 8 }, sm: { span: 8 } }
const labelColHalf2 = { lg: { span: 4 }, sm: { span: 4 } }

const noLabel = { offset: 4 }

const buildProgressInit = ['created']
const buildProgressPrepareRes = ['pending_res', 'launch_vm', 'create_vm_fail']
const buildProgressExec = ['perform_request_fail', 'appium_service_fail', 'running']
const buildProgressComplete = ['timeout', 'completed', 'cancel']

export { colsFull, colsHalf, labelColFull, wrapperColFull, labelColHalf, labelColHalf2, wrapperColHalf,
    noLabel, TEST_HISTORIES, buildProgressInit, buildProgressPrepareRes, buildProgressExec, buildProgressComplete }
