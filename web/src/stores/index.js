import { defineStore } from 'pinia'

// 用户状态管理
export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: null,
    token: null,
  }),
  actions: {
    setUserInfo(userInfo) {
      this.userInfo = userInfo
    },
    setToken(token) {
      this.token = token
    },
    clearUser() {
      this.userInfo = null
      this.token = null
    },
  },
})

// 任务状态管理
export const useTaskStore = defineStore('task', {
  state: () => ({
    tasks: [],
    currentDate: new Date(),
  }),
  actions: {
    setTasks(tasks) {
      this.tasks = tasks
    },
    setCurrentDate(date) {
      this.currentDate = date
    },
  },
})

