<template>
  <v-row>
    <v-text-field
      v-model="word"
      filled
      label="検索"
      prepend-inner-icon="mdi-magnify"
      @input="setSearchWord"
    >
    </v-text-field>
  </v-row>
</template>

<script>
export default {
  data() {
    return {
      word: '',
    }
  },
  methods: {
    setSearchWord() {
      this.$store.commit('meibo/setSearchWord', this.word)
    },
  },
}
</script>
