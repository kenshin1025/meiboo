<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      {{ $store.state.meibo.meibo }}
      <card
        v-for="(member, i) in $store.state.meibo.meibo"
        :key="i"
        :user="member"
      />
    </v-col>
  </v-row>
</template>

<script>
import Card from '~/components/meibo/Card.vue'
export default {
  components: {
    Card,
  },
  created() {
    this.$store.dispatch('meibo/getMeibo')
  },
}
</script>
