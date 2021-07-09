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
        <a-descriptions-item :label="$t('form.test.env')">{{ model.environments.length }}</a-descriptions-item>

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
      <a-steps :direction="'horizontal'" :current="1" progressDot>
        <a-step>
          <template v-slot:title>
            <span>创建项目</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              曲丽丽<a-icon type="dingding" style="margin-left: 8px;" />
              <div>2016-12-12 12:32</div>
            </div>
          </template>
        </a-step>
        <a-step>
          <template v-slot:title>
            <span>部门初审</span>
          </template>
          <template v-slot:description>
            <div class="antd-pro-pages-profile-advanced-style-stepDescription">
              周毛毛<a-icon type="dingding" style="color: rgb(0, 160, 233); margin-left: 8px;" />
              <div><a>催一下</a></div>
            </div>
          </template>
        </a-step>
        <a-step title="财务复核" />
        <a-step title="完成" />
      </a-steps>
    </a-card>

    <!-- operations -->
    <a-card
      v-if="tabActiveKey=='detail'"
      style="margin-top: 24px"
      :bordered="false"
      :tabList="operationTabList"
      :activeTabKey="operationActiveTabKey"
      @tabChange="(key) => {this.operationActiveTabKey = key}"
    >
      <a-table
        :columns="operationColumns"
        :dataSource="operations[operationActiveTabKey]"
        :pagination="false"
      >
        <span slot="action" slot-scope="text, record">
          <a :href="record.resultUrl">{{ $t('form.result.down') }}</a>
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

  </page-header-wrapper>
</template>

<script>
import { baseMixin } from '@/store/app-mixin'
import {
} from '@/utils/const'
import {
  getBuildTypes,
  getOsCategories,
  getOsTypes,
  getOsLangs,
  getBuildProgress,
  getBuildStatus,
  getVmStatus
} from '@/utils/testing'
import { getTask } from '@/api/manage'
import { clone } from '@/utils/util'

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
      model: { environments: [], queues: [] },
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

      operationTabList: [],
      operationActiveTabKey: '1',
      operations: {},

      operationColumns: []
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
        title: this.$t('form.step'),
        dataIndex: 'step',
        key: 'step'
      },
      {
        title: this.$t('form.result'),
        dataIndex: 'result',
        key: 'result'
      },
      {
        title: this.$t('form.time'),
        dataIndex: 'time',
        key: 'time'
      },
      {
        title: this.$t('form.opt'),
        dataIndex: 'action',
        key: 'action',
        width: '150px',
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
          this.model = json.data
          this.modelJson = this.convertJson(this.model)

          this.model.queues.forEach((queue, index) => {
            const name = this.osTypes[queue.osType] + ' ' + this.osLangs[queue.osLang]
            this.operationTabList.push({ key: queue.id + '', tab: name })

            // queue.buildHistories.forEach((buildHis, index) => {
            // })

            this.operations[queue.id] = [
              {
                key: 1,
                step: '创建虚拟机',
                result: '成功',
                time: '2017-10-03  19:23:12',
                resultUrl: 'http://localhost:8085/down/upload/2021-07-08/testResult-1ba63fd9-cb8c-4dd8-a942-86a74960b469.zip'
              }
            ]
          })

          this.operationTabList.push({ key: '2', tab: 'Win7 简体中文' })
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
        'priority', 'keepResultFiles']
      arr.forEach((item, index) => {
        json[item] = undefined
      })

      if (json.scriptUrl === '') {
        json.scriptUrl = json.scmAddress
        json.scmAddress = undefined
      }

      if (json.environments) {
        json.environments.forEach((item, index) => {
          item.id = undefined
          item.createdAt = undefined
          item.updatedAt = undefined
          item.taskId = undefined
        })
      }

      return json
    },
    handleTabChange (key) {
      console.log(key)
      this.tabActiveKey = key
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
