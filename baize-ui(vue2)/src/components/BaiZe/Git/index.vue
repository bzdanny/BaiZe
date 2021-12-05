<template>
  <div>
    <svg-icon icon-class="github" @click="goto"/>
  </div>
</template>

<script>
export default {
  name: 'BaiZeit',
  data() {
    return {
      url: 'https://gitee.com/smell2/BaiZe'
    }
  },
  methods: {
    goto() {
      window.open(this.url)
    }
  }
}
</script>
