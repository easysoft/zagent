<template>
  <global-footer class="footer custom-render">
    <template v-slot:links></template>
    <template v-slot:copyright>
      © 2021 <a href="https://github.com/easysoft/ztest" target="_blank">ztest</a>
    </template>
  </global-footer>
</template>

<script>
import { GlobalFooter } from '@ant-design-vue/pro-layout'
export default {
  name: 'ProGlobalFooter',
  components: {
    GlobalFooter
  }
}
</script>
