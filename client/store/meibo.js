import axios from 'axios'

export const state = () => ({
  meibo: [],
  tags: [],
  searchWord: '',
  filterTags: [],
})

export const mutations = {
  setMeibo(state, meibo) {
    state.meibo = meibo
  },
  addMember(state, member) {
    state.meibo.push(member)
  },
  setTags(state, tags) {
    state.tags = tags
  },
  setSearchWord(state, searchWord) {
    state.searchWord = searchWord
  },
  setFilterTag(state, tags) {
    state.filterTags = tags
  },
  addFilterTag(state, tag) {
    state.filterTags.push(tag)
  },
  removeFilterTag(state, tagID) {
    state.tags.splice(
      state.tags.findIndex((tag) => tag.id === tagID),
      1
    )
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
  async getTags({ commit }) {
    await axios.get('http://localhost:8080/meibo/tags').then((res) => {
      if (res.status === 200) {
        commit('setTags', res.data.tags)
      }
    })
  },
  async addMember({ commit }, postMeiboRequest) {
    await axios
      .post('http://localhost:8080/meibo', postMeiboRequest)
      .then((res) => {
        if (res.status === 200) {
          const member = { ...postMeiboRequest, token: res.data.token }
          commit('addMember', member)
        }
      })
      .catch((err) => {
        console.log(err)
      })
  },
  editMemberName({ commit, state }, { token, name }) {
    const index = state.meibo.findIndex((member) => member.token === token)
    commit('updateMemberName', { index, name })
  },
}

export const getters = {
  searchWord: (state) => {
    return state.searchWord
  },
  filterTags: (state) => {
    return state.filterTags
  },
  searchMeibo: (state, getters) => {
    if (getters.filterTags.length && getters.searchWord) {
      return state.meibo.filter(
        (member) =>
          (member.name.includes(getters.searchWord) ||
            member.comment.includes(getters.searchWord) ||
            member.tags.some((tag) => tag.name.includes(getters.searchWord))) &&
          getters.filterTags.every((filterTag) =>
            member.tags.some((tag) => tag.name === filterTag.name)
          )
      )
    } else if (getters.searchWord) {
      return state.meibo.filter(
        (member) =>
          member.name.includes(getters.searchWord) ||
          member.comment.includes(getters.searchWord) ||
          member.tags.some((tag) => tag.name.includes(getters.searchWord))
      )
    } else if (getters.filterTags.length) {
      return state.meibo.filter((member) =>
        getters.filterTags.every((filterTag) =>
          member.tags.some((tag) => tag.name === filterTag.name)
        )
      )
    } else {
      return state.meibo
    }
  },
}
