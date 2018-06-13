<template>
  <button @click="loadDevices">Load Devices</button>
</template>

<script>
export default {
  methods: {
    loadDevices () {
      this.$store.dispatch('devices/fetchAll')
    }
  },
  beforeCreate () {
    // TODO: Refactor to mapActions
    this.$store.dispatch('api/initial')
  },
  created () {
    // TODO: Refactior to mapActions
    this.$store.dispatch('devices/fetchAll')
  }
}
</script>
