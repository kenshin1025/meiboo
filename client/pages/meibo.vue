<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      {{ $store.state.meibo.meibo }}
      <Card
        v-for="(member, i) in $store.state.meibo.meibo"
        :key="i"
        :user="member"
      />
      <AddCard></AddCard>
    </v-col>
  </v-row>
</template>

<script>
import Card from '~/components/meibo/Card.vue'
import AddCard from '~/components/meibo/AddCard.vue'
export default {
  components: {
    Card,
    AddCard,
  },
  created() {
    this.$store.dispatch('meibo/getMeibo')
  },
}
</script>
