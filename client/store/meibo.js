import axios from 'axios'

export const state = () => ({
  meibo: [],
})

export const mutations = {
  setMeibo(state, meibo) {
    state.meibo = meibo
  },
  updateMemberName(state, { index, name }) {
    const member = state.meibo[index]
    member.name = name
    state.meibo.splice(index, 1, member)
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
  editMemberName({ commit, state }, { token, name }) {
    const index = state.meibo.findIndex((member) => member.token === token)
    commit('updateMemberName', { index, name })
  },
}
