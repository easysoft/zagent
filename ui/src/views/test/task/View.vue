<template>
  <page-header-wrapper
    :title="model.name"
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    @tabChange="handleTabChange"
  >
    <template v-slot:content>
      <a-descriptions size="small" :column="isMobile ? 1 : 2">
        <a-descriptions-item :label="$t('form.test.type')">{{ buildTypes[model.buildType] }}</a-descriptions-item>
        <a-descriptions-item :label="$t('form.test.env')">{{ model.environments.length }} {{ $t('form.uint.ge') }}</a-descriptions-item>

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

    <!-- actions -->
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

    <a-card :bordered="false" :title="$t('form.progress')">
      <a-steps :direction="isMobile && 'vertical' || 'horizontal'" :current="1" progressDot>
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

    <!-- 操作 -->
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      :tabList="operationTabList"
      :activeTabKey="operationActiveTabKey"
      @tabChange="(key) => {this.operationActiveTabKey = key}"
    >
      <a-table
        v-if="operationActiveTabKey === '1'"
        :columns="operationColumns"
        :dataSource="operation"
        :pagination="false"
      >
        <template
          slot="status"
          slot-scope="status">
          <a-badge :status="status | statusTypeFilter" :text="status | statusFilter"/>
        </template>
      </a-table>

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
      model: {},

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

      operationColumns: [
        {
          title: '操作类型',
          dataIndex: 'type',
          key: 'type'
        },
        {
          title: '操作人',
          dataIndex: 'name',
          key: 'name'
        },
        {
          title: '执行结果',
          dataIndex: 'status',
          key: 'status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '操作时间',
          dataIndex: 'updatedAt',
          key: 'updatedAt'
        },
        {
          title: '备注',
          dataIndex: 'remark',
          key: 'remark'
        }
      ],
      operation: [
        {
          key: 'op1',
          type: '订购关系生效',
          name: '曲丽丽',
          status: 'agree',
          updatedAt: '2017-10-03  19:23:12',
          remark: '-'
        }
      ]
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  filters: {
    statusFilter (status) {
      const statusMap = {
        'agree': '成功',
        'reject': '驳回'
      }
      return statusMap[status]
    },
    statusTypeFilter (type) {
      const statusTypeMap = {
        'agree': 'success',
        'reject': 'error'
      }
      return statusTypeMap[type]
    }
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
      { key: 'rule', tab: this.$t('common.request') }
    ]

    this.operationTabList = [ { key: '1', tab: '操作日志一' } ]
  },
  mounted () {
    this.loadData()
  },
  methods: {
    loadData () {
      if (!this.id) {
        return
      }
      if (this.id) {
        this.getModel().then(json => {
          this.model = json.data
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

    handleTabChange (key) {
      console.log('')
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

.no-data {
  color: rgba(0, 0, 0, .25);
  text-align: center;
  line-height: 64px;
  font-size: 16px;

  i {
    font-size: 24px;
    margin-right: 16px;
    position: relative;
    top: 3px;
  }
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
