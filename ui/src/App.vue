<template>
  <a-config-provider :locale="locale">
    <div id="app">
      <router-view/>
    </div>
  </a-config-provider>
</template>

<script>
import { domTitle, setDocumentTitle } from '@/utils/domUtil'
import { i18nRender } from '@/locales'
import * as neffos from 'neffos.js'
import { GetWebSocketApi } from '@/api/manage'

export default {
  data () {
    return {

    }
  },
  computed: {
    locale () {
      // 只是为了切换语言时，更新标题
      const { title } = this.$route.meta
      title && (setDocumentTitle(`${i18nRender(title)} - ${domTitle}`))

      return this.$i18n.getLocaleMessage(this.$store.getters.lang).antLocale
    }
  },
  created () {
    this.intiWebSocket()
  },
  methods: {
    async intiWebSocket () {
      const that = this

      try {
        const conn = await neffos.dial(GetWebSocketApi(), {
          default: {
            _OnNamespaceConnected: (nsConn, msg) => {
              if (nsConn.conn.wasReconnected()) {
                console.log('re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
              }

              console.log('connected to namespace: ' + msg.Namespace)
              that.$global.setWs(nsConn)
              that.$global.ws.joinRoom(that.$global.wsDefaultRoom)
            },
            _OnNamespaceDisconnect: (nsConn, msg) => {
              console.log('disconnected from namespace: ' + msg.Namespace)
            },
            OnVisit: (nsConn, msg) => {
              console.log('OnVisit', msg)
            },
            // implement in webpage
            OnChat: (nsConn, msg) => {
              console.log('OnChat in app', msg)
              console.log(msg.Room + ': response ' + msg.Body)

              that.$global.EventBus.$emit(that.$global.wsEventName, { 'room': msg.Room, msg: msg.Body })
            }
          }
        })
        await conn.connect('default')
      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>
