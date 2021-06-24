<template>
  <div>
    <div v-if="msg.type!=='question'" class="message a" :class="{'welcome':msg.type==='welcome'}">
      <div class="avatar">
        <a-avatar slot="avatar" icon="android" class="my-avatar-icon" :style="{ fontSize: '23px' }" />
      </div>
      <div class="box a" :class="{'welcome':msg.type==='welcome'}">
        <div class="content">
          <span v-if="msg.type==='welcome'">{{ $t('msg.testing.welcome') }}</span>
          <span v-else-if="msg.type==='pardon'">{{ $t('msg.testing.pardon') }}</span>
          <span v-else-if="msg.type==='answer'">{{ msg.content }}</span>
        </div>
        <div v-if="msg.type!=='welcome'" class="action">
          <a @click="view('result', msg.key)">{{ $t('common.view.result') }}</a>
          &nbsp;&nbsp;&nbsp;  | &nbsp;&nbsp;&nbsp;
          <a @click="view('json', msg.key)">{{ $t('common.view.json') }}</a>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <a v-if="viewMode!==''" @click="view('', 0)">
            |
            {{ $t('common.view.nothing') }}
          </a>
        </div>
      </div>
    </div>

    <div v-if="msg.type==='question'" class="message q">
      <div class="avatar">
        <a-avatar slot="avatar" icon="user" class="my-avatar-icon" :style="{ fontSize: '23px' }" />
      </div>
      <div class="box q">
        <div class="content">
          <span v-if="msg.type!=='welcome'">{{ msg.content }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ChatMessage',
  props: {
    msg: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  data () {
    return {
      typeMap: {},
      viewMode: ''
    }
  },
  watch: {
  },
  mounted () {
  },
  created () {
    this.typeMap = {
      true: {
        type: 'processing',
        text: this.$t('status.enable')
      },
      false: {
        type: 'default',
        text: this.$t('status.disable')
      }
    }
  },
  filters: {
    typeFilter (type, typeMap) {
      return typeMap[type].type
    }
  },
  methods: {
    view (mode, key) {
      this.viewMode = mode
      this.$emit('view', { mode: mode, key: key })
    }
  }
}
</script>

<style lang="less" scoped>

.message {
  position: relative;
  display: flex;
  margin-bottom: 10px;
  padding: 0 20px;
  width: calc(100% - 20px);

  .avatar {
    width: 30px;
    padding: 15px 0;
  }
  &.q {
    flex-direction:row-reverse;
    .avatar {
      padding: 10px 0;
    }
  }
  &.welcome {
    .avatar {
      padding: 6px 0;
    }
  }
  .box{
    position: relative;
    min-width: 220px;
    background: #f2f4f5;
    top:0px;
    -moz-border-radius: 12px;
    -webkit-border-radius: 12px;
    border-radius: 12px;
    padding: 10px;

    &.a {
      height: 60px;
      left: 28px;
      &:before {
        position: absolute;
        content: "";
        width: 0;
        height: 0;
        right: 100%;
        top: 18px;
        border-top: 13px solid transparent;
        border-bottom: 13px solid transparent;
        border-right: 26px solid #f2f4f5;
      }

      &.welcome {
        height: 45px;
        &:before {
          top: 10px;
        }
      }
    }
    &.q {
      height: 45px;
      right: 25px;
      &:after {
        position: absolute;
        content: "";
        width: 0;
        height: 0;
        right: -22px;
        top: 10px;
        border-top: 13px solid transparent;
        border-bottom: 13px solid transparent;
        border-left: 26px solid #f2f4f5;
      }
    }

    .content {
      color: rgba(0, 0, 0, 0.45);
      text-align: left;
    }
    .action {
      font-size: 12px;
      line-height: 23px;
    }
  }
}

</style>
