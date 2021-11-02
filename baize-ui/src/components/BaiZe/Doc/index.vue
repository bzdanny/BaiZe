<template>
  <div>
    <svg-icon icon-class="question" @click="goto"/>
  </div>
</template>

<script>
export default {
  name: 'BaiZeDoc',
  data() {
    return {
      url: 'http://www.baidu.com'
    }
  },
  methods: {
    goto() {
      window.open(this.url)
    }
  }
}
</script>
