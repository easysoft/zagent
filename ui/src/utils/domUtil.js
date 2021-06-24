import config from '@/config/defaultSettings'

export const setDocumentTitle = function (title) {
  document.title = title
  const ua = navigator.userAgent
  // eslint-disable-next-line
  const regex = /\bMicroMessenger\/([\d\.]+)/
  if (regex.test(ua) && /ip(hone|od|ad)/i.test(ua)) {
    const i = document.createElement('iframe')
    i.src = '/favicon.ico'
    i.style.display = 'none'
    i.onload = function () {
      setTimeout(function () {
        i.remove()
      }, 9)
    }
    document.body.appendChild(i)
  }
}

export const setSelectionRange = function (ctrl, pos) {
  console.log('setSelectionRange', ctrl, pos)

  setTimeout(() => {
    if (ctrl.setSelectionRange) {
      ctrl.focus()
      ctrl.setSelectionRange(pos, pos)
    } else if (ctrl.createTextRange) {
      const range = ctrl.createTextRange()
      range.collapse(true)
      range.moveEnd('character', pos)
      range.moveStart('character', pos)
      range.select()
    }
  }, 100)
}

export const domTitle = config.title
