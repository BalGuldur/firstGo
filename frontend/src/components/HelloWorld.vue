<template>
  <div>
    <button @click="loadDevices">Load Devices</button>
    <devices-table/>
    <device-form @submit="saveDevice" submitTitle="Добавить"/>
  </div>
</template>

<script>
import DeviceForm from '@/components/devices/DeviceForm'
import DevicesTable from '@/components/devices/DevicesTable'

export default {
  components: { DeviceForm, DevicesTable },
  methods: {
    loadDevices () {
      this.$store.dispatch('devices/fetchAll')
    },
    saveDevice (device) {
      this.$store.dispatch('devices/saveDevice', device)
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
