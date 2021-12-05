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
      url: 'https://www.ibaize.vip'
    }
  },
  methods: {
    goto() {
      window.open(this.url)
    }
  }
}
</script>
