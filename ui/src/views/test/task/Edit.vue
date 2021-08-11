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
          <a-form-model-item :label="$t('form.name')" prop="name" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-input v-model="model.name" />
          </a-form-model-item>

          <a-form-model-item :label="$t('form.test.type')" prop="buildType" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="model.buildType">
              <a-select-option v-for="(value, key) in buildTypes" :value="key" :key="key">
                {{ value }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
          <a-row v-if="model.buildType=='selenium'" :gutter="colsFull">
            <a-col :span="colsHalf">
              <a-form-model-item :label="$t('form.driver.type')" prop="seleniumDriverType" :labelCol="labelColHalf" :wrapperCol="wrapperColHalf">
                <a-input v-model="model.browserType" />
                <span>{{ $t('form.driver.type.tips') }}</span>
              </a-form-model-item>
            </a-col>

            <a-col :span="colsHalf">
              <a-form-model-item :label="$t('form.driver.version')" prop="seleniumDriverVersion" :labelCol="labelColHalf2" :wrapperCol="wrapperColHalf">
                <a-input v-model="model.browserVersion" />
                <span>{{ $t('form.driver.version.tips') }}</span>
              </a-form-model-item>
            </a-col>
          </a-row>

          <a-form-model-item :label="$t('form.test.code')" prop="scriptUrl" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-textarea v-model="model.scriptUrl" @change="scriptUrlChanged" />
            <span>{{ $t('form.test.code.tips') }}</span>
          </a-form-model-item>

          <a-row v-if="isScm" :gutter="colsFull">
            <a-col :span="colsHalf">
              <a-form-model-item :label="$t('form.scm.account')" prop="scmAccount" :labelCol="labelColHalf" :wrapperCol="wrapperColHalf">
                <a-input v-model="model.scmAccount" />
              </a-form-model-item>
            </a-col>

            <a-col :span="colsHalf">
              <a-form-model-item :label="$t('form.scm.password')" prop="scmPassword" :labelCol="labelColHalf2" :wrapperCol="wrapperColHalf">
                <a-input v-model="model.scmPassword" />
              </a-form-model-item>
            </a-col>
          </a-row>

          <a-form-model-item :label="$t('form.exec.cmd')" prop="buildCommands" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-textarea v-model="model.buildCommands" />
            <span>{{ $t('form.exec.cmd.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.result.files')" prop="resultFiles" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-textarea v-model="model.resultFiles" />
            <span>{{ $t('form.result.files.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.env.var')" prop="envVars" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-textarea v-model="model.envVars" />
            <span>{{ $t('form.env.var.tips') }}</span>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.desc')" prop="desc" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-textarea v-model="model.desc" />
          </a-form-model-item>

<!--          <a-form-model-item :label="$t('form.group')" prop="groupId" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-input-number v-model="model.groupId" />
            <span>  {{ $t('form.group.tips') }}</span>
          </a-form-model-item>-->

          <a-form-model-item :label="$t('form.test.env')" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
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
              <a-row v-for="(item, index) in model.environments" :key="index" :gutter="cols">
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

          <a-form-item :wrapperCol="wrapperColFull" style="text-align: center">
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
          <a-form-model-item :label="$t('form.os.category')" prop="osCategory" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="environment.osCategory" @change="envChanged()">
              <a-select-option v-for="(value, key) in envData.categories" :value="value" :key="key">
                {{ osCategories[value] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.os.type')" prop="osType" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
            <a-select v-model="environment.osType" @change="envChanged()">
              <a-select-option v-for="(value, key) in envData.types" :value="value" :key="key">
                {{ osTypes[value] }}
              </a-select-option>
            </a-select>
          </a-form-model-item>

          <a-form-model-item :label="$t('form.os.lang')" prop="osLang" :labelCol="labelColFull" :wrapperCol="wrapperColFull">
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
import {
  wrapperColFull,
  colsFull,
  colsHalf,
  labelColHalf,
  labelColHalf2,
  wrapperColHalf,
  labelColFull
} from '@/utils/const'
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
  computed: {
  },
  data () {
    return {
      colsFull: colsFull,
      colsHalf: colsHalf,

      labelColFull: labelColFull,
      wrapperColFull: wrapperColFull,

      labelColHalf: labelColHalf,
      labelColHalf2: labelColHalf2,
      wrapperColHalf: wrapperColHalf,

      model: {},
      envData: {},
      environment: {},
      environmentIndex: -1,
      isScm: false,
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
        osLang: [{ required: true, message: this.$t('valid.required.osLang'), trigger: 'blur' }],
        scriptUrl: [{ required: true, message: this.$t('valid.required.scriptUrl'), trigger: 'blur' }],
        buildCommands: [{ required: true, message: this.$t('valid.required.buildCommands'), trigger: 'blur' }],
        resultFiles: [{ required: true, message: this.$t('valid.required.resultFiles'), trigger: 'blur' }]
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
        const a = 1

        if (a === 1) {
          this.model = {
            'name': 'test',
            'buildType': 'selenium',
            'browserType': 'chrome',
            'browserVersion': '92',
            'envVars': 'abc=123',
            'scriptUrl': 'https://gitee.com/ngtesting/ci_test_selenium.git',
            'buildCommands': 'mvn clean test -Dtestng.suite=target/test-classes/baidu-test.xml',
            'resultFiles': 'target/surefire-reports',
            'environments': [ { 'osCategory': 'windows', 'osType': 'win10', 'osLang': 'zh_cn' } ]
          }
        } else {
          this.model = {
            'name': 'test',
            'buildType': 'testng',
            'envVars': 'abc=123',
            'scriptUrl': 'https://gitee.com/ngtesting/ci_test_testng.git',
            'buildCommands': 'docker run -i --rm --name testng-in-docker -v "$(pwd)":/usr/src/mymaven -v ~/.m2:/root/.m2 -w /usr/src/mymaven maven mvn clean package',
            'resultFiles': 'target/surefire-reports',
            'environments': [ { 'osCategory': 'linux', 'osType': 'ubuntu', 'osVersion': '20', 'osLang': 'zh_cn' } ]
          }
        }

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
      this.environment = {}
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
    },

    scriptUrlChanged () {
      if (this.model.scriptUrl.indexOf('.zip') > -1) {
        this.isScm = false
      } else {
        this.isScm = true
      }
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
