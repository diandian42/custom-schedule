import { createStore } from 'vuex'

const store = createStore({
  state: {
    userInfo: null,
    token: null,
    tasks: [],
  },
  mutations: {
    SET_USER_INFO(state, userInfo) {
      state.userInfo = userInfo
    },
    SET_TOKEN(state, token) {
      state.token = token
    },
    SET_TASKS(state, tasks) {
      state.tasks = tasks
    },
  },
  actions: {
    setUserInfo({ commit }, userInfo) {
      commit('SET_USER_INFO', userInfo)
    },
    setToken({ commit }, token) {
      commit('SET_TOKEN', token)
    },
    setTasks({ commit }, tasks) {
      commit('SET_TASKS', tasks)
    },
  },
})

export default store

