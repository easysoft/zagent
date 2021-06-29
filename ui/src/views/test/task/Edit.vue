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
        <a-form-model ref="form" :model="model" :rules="rules">
          <a-form-model-item :label="$t('form.name')" prop="name" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-input v-model="model.name" />
          </a-form-model-item>

          <a-form-model-item :label="$t('form.test.type')" prop="buildType" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-select v-model="model.buildType">
              <a-select-option v-for="(value, key) in buildTypes" :value="value[0]" :key="key">
                {{ value[1] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.exec.cmd')" prop="BuildCommands" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-textarea v-model="model.buildCommands" />
            <span>{{ $t('form.exec.cmd.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.env.var')" prop="envVars" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-textarea v-model="model.envVars" />
            <span>{{ $t('form.env.var.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.result.files')" prop="resultFiles" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-textarea v-model="model.resultFiles" />
            <span>{{ $t('form.result.files.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.desc')" prop="desc" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-textarea v-model="model.desc" />
          </a-form-model-item>

<!--          <a-form-model-item :label="$t('form.group')" prop="groupId" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-input-number v-model="model.groupId" />
            <span>  {{ $t('form.group.tips') }}</span>
          </a-form-model-item>-->

          <a-form-model-item :label="$t('form.test.env')" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <div class="environments">
              <a-row :gutter="cols" class="title">
                <a-col :offset="1" :span="col">{{ $t('form.os.category') }}</a-col>
                <a-col :span="col">{{ $t('form.os.type') }}</a-col>
                <a-col :span="col">{{ $t('form.os.lang') }}</a-col>
                <a-col :span="col-1">{{ $t('form.opt') }}</a-col>
              </a-row>
              <a-row v-if="!model.environments || model.environments.length == 0" :gutter="cols">
                <a-col :offset="col * 3 + 1" :span="col-1">
                  <a class="edit">
                    <a @click="addEnv(0)" class="edit">{{ $t('form.add') }}</a>
                  </a>
                </a-col>
              </a-row>
              <a-row v-for="(item, index)  in model.environments" :key="index" :gutter="cols">
                <a-col :offset="1" :span="col">
                  <span>{{ osCategories[item.osCategory] }}</span>
                </a-col>
                <a-col :span="col">
                  <span>{{ osTypes[item.osType] }}</span>
                </a-col>
                <a-col :span="col">
                  <span>{{ osLangs[item.osLang] }}</span>
                </a-col>

                <a-col :span="col-1">
                  <a class="edit">
                    <a @click="addEnv(index)" class="edit"><a-icon type="file-add" /></a> &nbsp;
                    <a @click="editEnv(index)" class="edit"><a-icon type="edit" /> </a> &nbsp;
                    <a @click="removeEnv(index)" class="edit"><a-icon type="delete" /></a> &nbsp;
                  </a>
                </a-col>
              </a-row>
            </div>
          </a-form-model-item>

          <a-form-item :wrapperCol="wrapperFull" style="text-align: center">
            <a-button @click="save()" htmlType="submit" type="primary">{{ $t('form.save') }}</a-button>
            <a-button @click="reset()" style="margin-left: 8px">{{ $t('form.reset') }}</a-button>
          </a-form-item>
        </a-form-model>
      </a-card>
    </page-header-wrapper>

    <a-modal
      :title="$t('form.edit.env')"
      :width="600"
      :visible="editEnvVisible"
      :okText="$t('form.save')"
      :cancelText="$t('form.cancel')"
      @ok="saveEnv"
      @cancel="cancelEnv">
      <div>
        <a-form-model ref="editEnvForm" :model="environment" :rules="rules">
          <a-form-model-item :label="$t('form.os.category')" prop="osCategory" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-select v-model="environment.osCategory" @change="envChanged()">
              <a-select-option v-for="(value, key) in envData.categories" :value="value" :key="key">
                {{ osCategories[value] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.os.type')" prop="osType" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-select v-model="environment.osType" @change="envChanged()">
              <a-select-option v-for="(value, key) in envData.types" :value="value" :key="key">
                {{  osTypes[value] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.os.lang')" prop="osLang" :labelCol="labelCol" :wrapperCol="wrapperCol">
            <a-select v-model="environment.osLang">
              <a-select-option v-for="(value, key) in envData.langs" :value="value" :key="key">
                {{ osLangs[value] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-form-model>
      </div>
    </a-modal>

  </div>
</template>

<script>
import { labelCol, wrapperCol, wrapperFull } from '@/utils/const'
import { getBuildTypes, getOsCategories, getOsTypes, getOsLangs } from '@/utils/testing'
import { requestSuccess, getTask, saveTask, getTestEnvs } from '@/api/manage'

export default {
  name: 'TaskEdit',
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
      labelCol: labelCol,
      wrapperCol: wrapperCol,
      wrapperFull: wrapperFull,

      model: { environments: [] },
      envData: {},
      environment: { osLang: 'zh_cn' },
      environmentIndex: -1,
      isInsert: false,

      editEnvVisible: false,
      buildTypes: {},
      osCategories: {},
      osTypes: {},
      osLangs: {},
      rules: {
        name: [{ required: true, message: this.$t('valid.required.name'), trigger: 'blur' }],
        buildType: [{ required: true, message: this.$t('valid.required.buildType'), trigger: 'blur' }],
        osCategory: [{ required: true, message: this.$t('valid.required.osCategory'), trigger: 'blur' }],
        osType: [{ required: true, message: this.$t('valid.required.osType'), trigger: 'blur' }],
        osLang: [{ required: true, message: this.$t('valid.required.osLang'), trigger: 'blur' }]
      },

      cols: 24,
      col: 6
    }
  },
  watch: {
    id: function () {
      console.log('watch id', this.id)
      this.loadData()
    }
  },
  created () {
    this.buildTypes = getBuildTypes(this)
    this.osCategories = getOsCategories(this)
    this.osTypes = getOsTypes(this)
    this.osLangs = getOsLangs(this)
  },
  mounted () {
    this.loadData()
  },
  methods: {
    loadTestEnvs () {
      getTestEnvs(this.environment).then(json => {
        this.envData = json.data
      })
    },
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
    save (e) {
      console.log(this.model)
      this.$refs.form.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        saveTask(this.model).then(json => {
          console.log('saveTask', json)
          if (requestSuccess(json.code)) {
            this.$router.push('/task/list')
          }
        })
      })
    },
    reset () {
      this.$refs.form.resetFields()
    },
    back () {
      this.$router.push('/task/list')
    },

    addEnv (index) {
      console.log('addEnv', index)
      this.environment = { osLang: 'zh_cn' }
      this.environmentIndex = index
      this.isInsert = true

      this.loadTestEnvs()
      this.editEnvVisible = true
    },
    editEnv (index) {
      console.log('editEnv', index)
      this.environment = this.model.environments[index]
      this.environmentIndex = index
      this.isInsert = false

      this.loadTestEnvs()
      this.editEnvVisible = true
    },
    removeEnv (index) {
      console.log('removeEnv', index)
      this.isInsert = false
      this.model.environments.splice(index, 1)
    },
    saveEnv () {
      console.log('saveEnv')

      this.$refs.editEnvForm.validate(valid => {
        if (!valid) {
          console.log('validate fail', valid)
          return false
        }

        if (this.isInsert) {
          this.model.environments.splice(this.environmentIndex + 1, 0, this.environment)
        } else {
          this.model.environments[this.environmentIndex] = this.environment
        }

        this.editEnvVisible = false
      })
    },
    cancelEnv () {
      console.log('cancelEnv')
      this.editEnvVisible = false
    },
    envChanged () {
      this.loadTestEnvs()
    }
  }
}
</script>

<style lang="less" scoped>
.environments {
  margin-bottom: 10px;

  .title {
    font-weight: bolder;
    margin-bottom: 5px;
    padding-bottom: 5px;
    border-bottom: 1px solid #e9f2fb;
  }

  .edit {
    line-height: 32px;
  }
}
</style>
