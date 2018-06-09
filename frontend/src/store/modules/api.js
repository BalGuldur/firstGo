// import WebSocket from 'faye-websocket'

const state = {
  items: [],
  conn: null
}

const getters = {
  conn: state => state.conn
}

const actions = {
  initial ({commit}) {
    let conn = new WebSocket('ws://localhost:8085/ws')
    conn.onmessage = function(event) {
      alert("Получены данные " + event.data);
    };
    commit('setConn', conn)
  }
}

const mutations = {
  setConn (state, conn) {
    state.conn = conn
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
