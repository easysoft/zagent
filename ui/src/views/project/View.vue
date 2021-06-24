<template>
  <page-header-wrapper
    :title="model.name"
  >
    <template v-slot:content>
      <a-descriptions size="small" :column="isMobile ? 1 : 2">
        <a-descriptions-item :label="$t('form.createdBy')">{{ model.createdBy }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.path')">{{ model.path }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.createdAt')">{{ model.createdAt | moment }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.updatedAt')">{{ model.updtedAt | moment }}</a-descriptions-item>

        <a-descriptions-item :label="$t('form.desc')"></a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- actions -->
    <template v-slot:extra>
      <a-button-group style="margin-right: 4px;">
        <a-button @click="compile()">{{ $t('common.compile') }}</a-button>

        <a-button
          @click="training()"
          :disabled="model.trainingStatus === 'start_training'">
            {{ $t('common.training') }}
        </a-button>

        <template v-if="model.serviceStatus !== 'start_service'">
          <a-button @click="startService()" :disabled="model.trainingStatus !== 'end_training'">
            {{ $t('common.start_service') }}
          </a-button>
        </template>
        <a-button @click="endService()" v-if="model.serviceStatus === 'start_service'">{{ $t('common.stop_service') }}</a-button>

      </a-button-group>
      <a-button @click="back()" type="primary">{{ $t('common.back') }}</a-button>
    </template>

    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="9" :sm="9"></a-col>
        <a-col :xs="12" :sm="12">
          <div class="text">{{ $t('common.status') }}</div>
          <div class="heading">
            <a-badge
              v-if="model.trainingStatus === 'start_training'"
              :status="'processing'"
              :text="$t('status.start.training')" />

            <a-badge
              v-if="model.trainingStatus !== 'start_training'"
              :status="model.serviceStatus | statusTypeFilter(statusMap)"
              :text="model.serviceStatus | statusFilter(statusMap)" />
          </div>
        </a-col>
        <a-col :xs="3" :sm="3"></a-col>
      </a-row>
    </template>

    <a-card
      style="margin-top: 24px"
      :title="$t('form.opt.log')"
      :bordered="false">

      <div>
        <a-table
          rowKey="id"
          :columns="historyColumns"
          :dataSource="model.histories"
          :pagination="false"
        >
          <span slot="serial" slot-scope="text, record, index">
            {{ index + 1 }}
          </span>
          <span slot="action" slot-scope="text, record">
            {{ $t('common.' + record.action) }}
          </span>
          <span slot="createdTime" slot-scope="text, record">
            {{ record.createdTime | moment }}
          </span>
        </a-table>
      </div>
    </a-card>

    <br />
    <div>
      <a-form layout="inline">
        <a-form-item><a-input id="input" type="text" v-model="inputModel" /></a-form-item>
        <a-form-item><a-button id="sendBtn" @click="sendWs">Send</a-button></a-form-item>
      </a-form>
      <div><pre id="output">{{ outputModel }}</pre></div>
    </div>

  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { getProject, compileProject, trainingProject, startService, endService } from '@/api/manage'

import { baseMixin } from '@/store/app-mixin'

export default {
  name: 'ProjectEdit',
  mixins: [baseMixin],
  statusMap: {},
  props: {
    id: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.id)
      }
    }
  },
  data () {
    return {
      wsConn: null,
      room1: null,
      inputModel: 'websocket request',
      outputModel: '',

      model: {},
      moment,

      historyColumns: [
        {
          title: this.$t('form.no'),
          scopedSlots: { customRender: 'serial' }
        },
        {
          title: '操作类型',
          scopedSlots: { customRender: 'action' }
        },
        {
          title: '操作人',
          dataIndex: 'userName',
          key: 'userName'
        },
        {
          title: '操作时间',
          scopedSlots: { customRender: 'createdTime' }
        }
      ],

      optTabList: [
        {
          key: '1',
          tab: '测试'
        },
        {
          key: '2',
          tab: '日志'
        }
      ],
      optActiveTabKey: '1'
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  mounted () {
    this.loadData()
  },
  created () {
    const that = this
    this.$global.EventBus.$on(this.$global.wsEventName, (json) => {
      console.log('EventBus in page', json)
      that.outputModel += json.room + ': ' + json.msg + '\n'

      const msg = JSON.parse(json.msg)
      if (msg.action === 'end_training' && msg.projectId === that.model.id) {
        that.model.trainingStatus = 'end_training'
      }
    })

    this.statusMap = {
      '': {
        type: 'default',
        text: ''
      },
      start_service: {
        type: 'success',
        text: this.$t('status.start.service')
      },
      stop_service: {
        type: 'default',
        text: this.$t('status.stop.service')
      }
    }
  },
  filters: {
    statusFilter (status, statusMap) {
      if (!status) status = ''
      return statusMap[status].text
    },
    statusTypeFilter (status, statusMap) {
      if (!status) status = ''
      return statusMap[status].type
    }
  },
  methods: {
    sendWs () {
      console.log('sendWs', this.inputModel, this.$global.ws)
      this.$global.ws.room(this.$global.wsDefaultRoom).emit('OnChat', this.inputModel)

      this.outputModel += 'me: ' + this.inputModel + '\n'
    },

    loadData () {
      if (!this.id) {
        return
      }
      if (this.id) {
        getProject(this.id).then(json => {
          this.model = json.data
        })
      } else {
        this.reset()
      }
    },
    compile () {
      console.log('compile')
      compileProject(this.model).then(json => {
        console.log('compile', json)

        if (json.code === 200) {
          const that = this
          this.$notification['success']({
            message: that.$t('common.tips'),
            description: that.$t('msg.compile.success'),
            // placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
        }
      })
    },
    training () {
      console.log('training')

      trainingProject(this.model).then(json => {
        console.log('training', json)
        if (json.code === 200) {
          this.model.trainingStatus = json.data.trainingStatus
          const that = this
          this.$notification['success']({
            message: that.$t('common.tips'),
            description: that.$t('msg.training.start'),
            // placement: 'bottomRight',
            duration: 8
          })
        }
      })
    },
    startService () {
      console.log('startService')
      startService(this.model).then(json => {
        console.log('startService', json)
        if (json.code === 200) {
          this.$notification['success']({
            message: this.$root.$t('common.tips'),
            description: this.$root.$t('msg.service.start'),
            // placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
        }
      })
    },
    endService () {
      console.log('endService')
      endService(this.model).then(json => {
        console.log('endService', json)
        if (json.code === 200) {
          this.$notification['success']({
            message: this.$root.$t('common.tips'),
            description: this.$root.$t('msg.service.stop'),
            // placement: 'bottomRight',
            duration: 8
          })

          this.loadData()
        }
      })
    },
    back () {
      this.$router.push('/project/list')
    }
  }
}
</script>

<style lang="less" scoped>

.detail-layout {
  margin-left: 44px;
}
.text {
  color: rgba(0, 0, 0, .45);
}

.heading {
  color: rgba(0, 0, 0, .85);
  font-size: 20px;
}

</style>
