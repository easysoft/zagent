<template>
  <div>
    <page-header-wrapper>
      <a-card :bordered="false">
        <div class="table-page-search-wrapper">
          <a-form layout="inline">
            <a-row :gutter="48">
              <a-col :md="8" :sm="24">
                <a-form-item :label="$t('form.name')">
                  <a-input v-model="queryParam.keywords" placeholder=""/>
                </a-form-item>
              </a-col>
              <a-col :md="8" :sm="24">
                <a-form-item :label="$t('form.status')">
                  <a-select v-model="queryParam.status">
                    <a-select-option value="">{{ $t('form.all') }}</a-select-option>
                    <a-select-option value="true">{{ $t('form.enable') }}</a-select-option>
                    <a-select-option value="false">{{ $t('form.disable') }}</a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :md="!advanced && 8 || 24" :sm="24">
                <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                  <a-button type="primary" @click="$refs.table.refresh(true)">{{ $t('form.search') }}</a-button>
                  <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">{{ $t('form.reset') }}</a-button>
                </span>
              </a-col>
            </a-row>
          </a-form>
        </div>

        <div class="table-operator">
          <a-button type="primary" icon="plus" @click="create">{{ $t('form.create') }}</a-button>
        </div>

        <s-table
          ref="table"
          size="default"
          rowKey="id"
          :columns="columns"
          :data="loadData"
          :alert="true"
          showPagination="auto"
        >
          <span slot="serial" slot-scope="text, record, index">
            {{ index + 1 }}
          </span>

          <span slot="name" slot-scope="text">
            <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
          </span>

          <span slot="path" slot-scope="text">
            <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
          </span>

          <span slot="status" slot-scope="text, record">
            <a-badge :status="!record.disabled | statusTypeFilter(statusMap)" :text="!record.disabled | statusFilter(statusMap)" />
          </span>

          <span slot="action" slot-scope="text, record">
            <template>
              <a @click="view(record)">{{ $t('form.view') }}</a>
              <a-divider type="vertical" />

              <a @click="test(record)">{{ $t('form.test') }}</a>
              <a-divider type="vertical" />

              <a @click="edit(record)">{{ $t('form.edit') }}</a>
              <a-divider type="vertical" />

              <a v-if="!record.disabled" @click="disable(record)">{{ $t('form.disable') }}</a>
              <a v-if="record.disabled" @click="disable(record)">{{ $t('form.enable') }}</a>
              <a-divider type="vertical" />

              <a-popconfirm
                :title="$t('form.confirm.to.remove')"
                :okText="$t('form.ok')"
                :cancelText="$t('form.cancel')"
                @confirm="confirmRemove(record)"
                @cancel="cancelRemove"
              >
                <a href="#">{{ $t('form.remove') }}</a>
              </a-popconfirm>
            </template>
          </span>
        </s-table>
      </a-card>
    </page-header-wrapper>
  </div>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { listPlan, disablePlan, removePlan } from '@/api/manage'

export default {
  name: 'PlanList',
  components: {
    STable,
    Ellipsis
  },
  columns: [],
  statusMap: {},
  data () {
    return {
      isProd: process.env.NODE_ENV === 'production',

      visible: false,
      confirmLoading: false,
      mdl: null,
      advanced: false,
      queryParam: { status: '' },
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        return listPlan(requestParameters)
          .then(res => {
            return res
          })
      },
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    statusFilter (status, statusMap) {
      return statusMap[status].text
    },
    statusTypeFilter (status, statusMap) {
      return statusMap[status].type
    }
  },
  created () {
    const that = this
    this.$global.EventBus.$on(this.$global.wsEventName, (v) => {
      console.log('EventBus in page', v)

      that.outputModel += v.room + ': ' + v.msg + '\n'
    })

    this.columns = [
      {
        title: this.$t('form.no'),
        scopedSlots: { customRender: 'serial' }
      },
      {
        title: this.$t('form.name'),
        dataIndex: 'name'
      },
      {
        title: this.$t('form.path'),
        dataIndex: 'path'
      },
      {
        title: this.$t('form.status'),
        dataIndex: 'status',
        scopedSlots: { customRender: 'status' }
      },
      // {
      //   title: this.$t('form.is.default'),
      //   dataIndex: 'default',
      //   scopedSlots: { customRender: 'default' }
      // },
      {
        title: this.$t('form.opt'),
        dataIndex: 'action',
        width: '260px',
        scopedSlots: { customRender: 'action' }
      }
    ]

    this.statusMap = {
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
  computed: {

  },
  methods: {
    create () {
      this.mdl = null
      this.visible = true

      this.$router.push('/test/plan/0/edit')
    },
    test (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/test/plan/' + record.id + '/test')
    },
    view (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/test/plan/' + record.id + '/view')
    },
    edit (record) {
      this.visible = true
      this.mdl = { ...record }

      this.$router.push('/test/plan/' + record.id + '/edit')
    },

    disable (record) {
      disablePlan(record).then(json => {
        console.log('disablePlan', json)
        this.$refs.table.refresh(false)
      })
    },
    confirmRemove (record) {
      removePlan(record).then(json => {
        console.log('removePlan', json)
        this.$refs.table.refresh(false)
      })
    },
    cancelRemove (e) {
      console.log(e)
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
