<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      {{ $store.state.meibo.meibo }}
      <Card
        v-for="(member, i) in $store.state.meibo.meibo"
        :key="i"
        :member="member"
      />
      <AddMember></AddMember>
    </v-col>
  </v-row>
</template>

<script>
import Card from '~/components/meibo/Card.vue'
import AddMember from '~/components/meibo/AddMember.vue'
export default {
  components: {
    Card,
    AddMember,
  },
  created() {
    this.$store.dispatch('meibo/getMeibo')
  },
}
</script>
