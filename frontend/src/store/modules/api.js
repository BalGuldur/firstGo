// TODO: Refactor: remove logs, think about console.log
const state = {
  items: [],
  conn: null
}

const getters = {
  conn: state => state.conn
}

const actions = {
  initial ({state, commit, dispatch}) {
    console.log('initial')
    if (state.conn === null || state.conn.readyState !== 1) {
      let conn = new WebSocket('ws://localhost:8085/ws')
      conn.onopen = () => {
        commit('setConn', conn)
      }
      conn.onmessage = (event) => {
        // TODO: Remove log
        console.log('Получены данные ' + event.data)
        dispatch('proceedMessage', event.data)
      }
      conn.onerror = (error) => {
        commit('closeConn')
        console.log('Ошибка ' + error.message)
      }
      conn.onclose = (event) => {
        commit('closeConn')
        if (event.wasClean) {
          console.log('Соединенеие закрыто частично')
        } else {
          console.log('Обрыв соединения')
        }
        console.log(`Код: ${event.code} причина: ${event.reason}`)
        setTimeout(() => {
          console.log('попытка восстановить')
          dispatch('initial')
        }, 1000)
      }
    }
  },
  proceedMessage ({dispatch}, message) {
    message = JSON.parse(message)
    // TODO: Remove log
    console.log('message in json', message)
    let {success} = message
    if (success === true) {
      let model = message.action.split('.')[0]
      let dispAction = message.action.replace('.', '/')
      let modelObj = message[model]
      dispatch(dispAction, modelObj, {root: true})
    } else {
      dispatch('showError', message)
    }
  },
  sendMessage ({state, dispatch}, message) {
    if (state.conn && state.conn.readyState === 1) {
      state.conn.send(JSON.stringify(message))
    } else {
      // TODO: Refactor wait before initial WS on home page
      setTimeout(() => {
        state.conn.send(JSON.stringify(message))
      }, 2000)
    }
  },
  showError (_, message) {
    alert('receive error: ' + message)
  }
}

const mutations = {
  setConn (state, conn) {
    state.conn = conn
  },
  closeConn (state) {
    state.conn = null
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
