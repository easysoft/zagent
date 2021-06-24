<template>
  <global-footer class="footer custom-render">
    <template v-slot:links></template>
    <template v-slot:copyright>
      © 2021 <a href="https://util.ai" target="_blank">util.ai</a>
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
