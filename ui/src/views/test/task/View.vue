<template>
  <page-header-wrapper
    :title="model.id + '-' + model.name"
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    @tabChange="handleTabChange"
  >
    <!-- summury -->
    <template v-slot:content>
      <a-descriptions size="small" :column="2">
        <a-descriptions-item :label="$t('form.test.type')">{{ buildTypes[model.buildType] }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.test.env')">
          {{ model.environments ? model.environments.length: 0 }}
        </a-descriptions-item>

        <a-descriptions-item :label="$t('form.createdAt')">{{ model.createdAt | moment }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.pendingAt')">
          {{ model.pendingTime ? $options.filters.moment(model.pendingTime) : '' }}
        </a-descriptions-item>

        <a-descriptions-item :label="$t('form.startAt')">
          {{ model.startTime ? $options.filters.moment(model.startTime) : '' }}
        </a-descriptions-item>
        <a-descriptions-item :label="$t('form.completeAt')">
          {{ model.resultTime ? $options.filters.moment(model.resultTime) : '' }}
        </a-descriptions-item>

        <a-descriptions-item :label="$t('form.desc')">{{ model.desc }}</a-descriptions-item>
      </a-descriptions>
    </template>
    <template v-slot:extra>
      <a-button @click="back()" type="primary" >{{ $t('common.back') }}</a-button>
    </template>
    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="12" :sm="12">
          <div class="text">{{ $t('form.progress') }}</div>
          <div class="heading highlight-blue">{{ buildProgress[model.progress] }}</div>
        </a-col>
        <a-col :xs="12" :sm="12">
          <div class="text">{{ $t('form.status') }}</div>
          <div class="heading highlight-blue">{{ buildStatus[model.status] }}</div>
        </a-col>
      </a-row>
    </template>

    <!-- progress -->
    <a-card v-if="tabActiveKey=='detail'" :bordered="false" :title="$t('form.progress')">

      <a-steps :direction="'horizontal'" :current="currStep" progressDot>
        <a-step>
          <template v-slot:title>
            <span>{{ $t('build.progress.start') }}</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              {{taskProgressMap['start']['status']}} <br/>
              {{taskProgressMap['start']['time'] ? $options.filters.moment(taskProgressMap['start']['time'], 'MM-DD HH:mm:ss') : '' }}
            </div>
          </template>
        </a-step>
        <a-step>
          <template v-slot:title>
            <span>{{ $t('build.progress.res') }}</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              {{taskProgressMap['res']['status']}} <br/>
              {{taskProgressMap['res']['time'] ? $options.filters.moment(taskProgressMap['res']['time'], 'MM-DD HH:mm:ss') : '' }}
            </div>
          </template>
        </a-step>
        <a-step>
          <template v-slot:title>
            <span>{{ $t('build.progress.exec') }}</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              {{taskProgressMap['exec']['status']}} <br/>
              {{taskProgressMap['exec']['time'] ? $options.filters.moment(taskProgressMap['exec']['time'], 'MM-DD HH:mm:ss') : '' }}
            </div>
          </template>
        </a-step>
        <a-step>
          <template v-slot:title>
            <span>{{ $t('build.progress.end') }}</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              {{taskProgressMap['end']['status']}} <br/>
              {{taskProgressMap['end']['time'] ? $options.filters.moment(taskProgressMap['end']['time'], 'MM-DD HH:mm:ss') : '' }}
            </div>
          </template>
        </a-step>
      </a-steps>

    </a-card>

    <!-- operations -->
    <a-card
      v-if="tabActiveKey=='detail'"
      style="margin-top: 24px"
      :bordered="false"
      :tabList="taskQueueTabs"
      :activeTabKey="operationActiveTabKey"
      @tabChange="(key) => {this.operationActiveTabKey = key}"
    >
      <a-table
        :columns="operationColumns"
        :dataSource="taskBuildHistories[operationActiveTabKey]"
        :pagination="false"
      >
        <span slot="time" slot-scope="text, record">
          {{ record.createdAt | moment }}
        </span>
        <span slot="action" slot-scope="text, record">
          <span v-if="record.resultUrl">
            <a :href="record.resultUrl">{{ $t('form.result.down') }}</a>
          </span>
          <span v-if="record.resultUrl && record.vncUrl"> | </span>
          <span v-if="record.vncUrl">
            <a :href="record.vncUrl">{{ $t('form.vnc.url') }}</a>
          </span>
        </span>
      </a-table>
    </a-card>

    <a-card
      v-if="tabActiveKey=='request'"
      style="margin-top: 24px"
      :bordered="false"
    >
      <pre>{{ JSON.stringify(modelJson, null, 4) }}</pre>
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
import { baseMixin } from '@/store/app-mixin'
import {
  getBuildTypes,
  getOsCategories,
  getOsTypes,
  getOsLangs,
  getBuildProgress,
  getBuildStatus,
  getVmStatus, getTaskProgressMap, getTaskBuildHistories
} from '@/utils/testing'
import { getTask } from '@/api/manage'
import { clone, getBuildStep } from '@/utils/util'

export default {
  name: 'TaskView',
  mixins: [baseMixin],
  props: {
    id: {
      type: Number,
      default: function () {
        return parseInt(this.$route.params.id)
      }
    }
  },
  computed: {
  },
  data () {
    return {
      model: { environments: [], queues: [], histories: [] },
      modelJson: {},

      buildProgress: {},
      buildStatus: {},
      vmStatus: {},

      buildTypes: {},
      osCategories: {},
      osTypes: {},
      osLangs: {},

      tabList: [],
      tabActiveKey: 'detail',

      taskQueueTabs: [],
      operationActiveTabKey: '1',
      operationColumns: [],

      currStep: 0,
      taskProgressMap: { start: {}, res: {}, exec: {}, end: {} },
      taskBuildHistories: {},

      wsConn: null,
      room1: null,
      inputModel: 'websocket request',
      outputModel: ''
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  filters: {
  },
  created () {
    const that = this
    this.$global.EventBus.$on(this.$global.wsEventName, (json) => {
      console.log('wsEvent', json)
      that.outputModel += json.room + ': ' + json.msg + '\n'

      const msg = JSON.parse(json.msg)
      if (msg.action === 'task_update' && msg.taskId === that.model.id) {
        console.log('task_update', json)
        this.loadData()
      }
    })

    this.buildProgress = getBuildProgress(this)
    this.buildStatus = getBuildStatus(this)
    this.vmStatus = getVmStatus(this)

    this.buildTypes = getBuildTypes(this)
    this.osCategories = getOsCategories(this)
    this.osTypes = getOsTypes(this)
    this.osLangs = getOsLangs(this)

    this.tabList = [
      { key: 'detail', tab: this.$t('common.detail') },
      { key: 'request', tab: this.$t('common.request') }
    ]

    this.operationColumns = [
      {
        title: this.$t('form.no'),
        dataIndex: 'key',
        key: 'key'
      },
      {
        title: this.$t('form.type'),
        dataIndex: 'type',
        key: 'type'
      },
      {
        title: this.$t('form.progress'),
        dataIndex: 'progress',
        key: 'progress'
      },
      {
        title: this.$t('form.status'),
        dataIndex: 'status',
        key: 'status'
      },
      {
        title: this.$t('form.time'),
        dataIndex: 'time',
        key: 'time',
        scopedSlots: { customRender: 'time' }
      },
      {
        title: this.$t('form.opt'),
        dataIndex: 'action',
        key: 'action',
        width: '200px',
        scopedSlots: { customRender: 'action' }
      }
    ]
  },
  mounted () {
    this.loadData()
  },
  methods: {
    loadData () {
      if (!this.id) return
      if (this.id) {
        this.getModel().then(json => {
          this.model = json.data.task
          this.modelJson = this.convertJson(this.model)

          this.taskQueueTabs = []
          this.model.queues.forEach((queue, index) => {
            const name = this.osTypes[queue.osType] + ' ' + this.osLangs[queue.osLang]
            this.taskQueueTabs.push({ key: queue.id + '', tab: name })
          })

          this.currStep = getBuildStep(this.model.progress)
          this.taskProgressMap = getTaskProgressMap(this.model.histories, this.buildProgress)
          this.taskBuildHistories = getTaskBuildHistories(json.data.buildHistories, this)

          console.log('this.taskProgressMap', this.taskProgressMap, this.taskBuildHistories)
        })
      } else {
        this.reset()
      }
    },
    getModel () {
      return getTask(this.id)
    },
    back () {
      this.$router.push('/task/list')
    },

    convertJson (task) {
      const json = clone(task)
      const arr = ['queues', 'id', 'createdAt', 'updatedAt', 'progress', 'status',
        'startTime', 'pendingTime', 'resultTime', 'userName', 'userId', 'groupId',
        'priority', 'keepResultFiles', 'histories']
      arr.forEach((item, index) => {
        json[item] = undefined
      })

      if (json.scriptUrl === '') {
        json.scriptUrl = json.scmAddress
        json.scmAddress = undefined
      }

      if (json.environments) {
        json.environments.forEach((item, index) => {
          const arr = ['id', 'createdAt', 'updatedAt', 'taskId', 'deletedAt', 'deleted', 'disabled']
          arr.forEach((i, index) => {
            item[i] = undefined
          })
        })
      }

      return json
    },
    handleTabChange (key) {
      console.log(key)
      this.tabActiveKey = key
    },

    sendWs () {
      console.log('sendWs', this.inputModel, this.$global.ws)
      this.$global.ws.room(this.$global.wsDefaultRoom).emit('OnChat', this.inputModel)

      this.outputModel += 'me: ' + this.inputModel + '\n'
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

.mobile {
  .detail-layout {
    margin-left: unset;
  }
  .text {

  }
  .status-list {
    text-align: left;
  }
}
</style>
