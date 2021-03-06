import Vue from 'vue'

const state = {
  items: {}
}

const getters = {
  items: state => Object.values(state.items)
}

const actions = {
  fetchAll ({dispatch}) {
    let msg = {action: 'devices.all'}
    dispatch('api/sendMessage', msg, {root: true})
  },
  saveDevice ({dispatch}, device) {
    let msg = {action: 'devices.save', device}
    dispatch('api/sendMessage', msg, {root: true})
  },
  removeDevice ({dispatch}, device) {
    let msg = {action: 'devices.delete', device}
    dispatch('api/sendMessage', msg, {root: true})
  },
  // actions from messages
  save ({commit}, devices) {
    devices.forEach(device => commit('set', device))
  },
  delete ({commit}, devices) {
    devices.forEach(device => commit('delete', device))
  },
  all ({commit}, devices) {
    commit('setAll', devices)
  },
  get ({dispatch}, devices) {
    dispatch('save', devices)
  }
}

const mutations = {
  set (state, device) {
    device && Vue.set(state.items, device.id, device)
  },
  delete (state, device) {
    device && Vue.delete(state.items, device.id)
  },
  setAll (state, devices) {
    let items = devices && devices.reduce((prev, device) => {
      return {...prev, [device.id]: device}
    }, {})
    state.items = items || {}
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
