<template>
  <div>
    <page-header-wrapper content="">
      <div class="toolbar-edit">
        <div class="left"></div>
        <div class="right">
          <a-button @click="back()" type="primary">{{ $t('common.back') }}</a-button>
        </div>
      </div>
      <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
        VIEW
      </a-card>
    </page-header-wrapper>

  </div>
</template>

<script>
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
      osLangs: {}
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
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
    }
  }
}
</script>

<style lang="less" scoped>

</style>
