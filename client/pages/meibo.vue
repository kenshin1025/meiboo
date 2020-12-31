<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      {{ $store.state.meibo.meibo }}
    </v-col>
  </v-row>
</template>

<script>
export default {
  // components: {
  // },
  created() {
    this.$store.dispatch('meibo/getMeibo')
  },
}
</script>
