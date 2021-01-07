import axios from 'axios'

export const state = () => ({
  meibo: [],
})

export const mutations = {
  setMeibo(state, meibo) {
    state.meibo = meibo
  },
}

export const actions = {
  async getMeibo({ commit }) {
    await axios.get('http://localhost:8080/meibo').then((res) => {
      if (res.status === 200) {
        commit('setMeibo', res.data.meibo)
      }
    })
  },
}
