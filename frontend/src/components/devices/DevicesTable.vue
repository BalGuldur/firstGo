<template>
  <div>
    <device-form
      v-for="device in devices"
      :key="device.id"
      :initObj="device"
      submitTitle="Изменить"
      @submit="saveDevice"
    >
      <button @click="removeDevice(device)">Удалить</button>
    </device-form>
  </div>
</template>

<script>
import DeviceForm from '@/components/devices/DeviceForm'

import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'DevicesTable',
  components: { DeviceForm },
  computed: {
    ...mapGetters({devices: 'devices/items'})
  },
  methods: {
    ...mapActions('devices', ['saveDevice', 'removeDevice'])
  }
}
</script>
