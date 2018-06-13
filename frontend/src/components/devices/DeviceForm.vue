<template>
  <div>
    <input v-model="device.name" placeholder="наименование"/>
    <button v-if="showCancel" @click="cancel">Отмена</button>
    <button @click="submit">{{submitTitle}}</button>
    <slot/>
  </div>
</template>

<script>
import isEqual from 'lodash/isequal'

export default {
  name: 'DeviceForm',
  props: {
    showCancel: {
      type: Boolean,
      default: false
    },
    submitTitle: {
      type: String,
      default: 'Сохранить'
    },
    initObj: {
      type: Object,
      default () {
        return { name: '' }
      }
    }
  },
  data () {
    return {
      device: { ...this.initObj }
    }
  },
  methods: {
    submit () {
      this.$emit('submit', this.device)
      if (!('initObj' in this.$options.propsData)) {
        this.device = { ...this.initObj }
        this.changed = false
      }
    }
  },
  computed: {
    changed () {
      return !isEqual(this.initObj, this.device)
    }
  },
  watch: {
    initObj () {
      !this.changed && (this.device = { ...this.initObj })
    }
  }
}
</script>
