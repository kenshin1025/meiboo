<template>
  <v-row justify="center" align="center">
    <v-col cols="8">
      <SearchBox></SearchBox>
      <TagFilter></TagFilter>
      <v-row>
        <AddMember></AddMember>
        <v-col
          v-for="(member, i) in $store.getters['meibo/searchMeibo']"
          :key="i"
          cols="12"
          sm="8"
          md="6"
        >
          <Card :member="member" />
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import Card from '~/components/meibo/Card.vue'
import AddMember from '~/components/meibo/AddMember.vue'
import SearchBox from '~/components/meibo/SearchBox.vue'
import TagFilter from '~/components/meibo/TagFilter.vue'
export default {
  components: {
    Card,
    AddMember,
    SearchBox,
    TagFilter,
  },
  created() {
    this.$store.dispatch('meibo/getMeibo')
    this.$store.dispatch('meibo/getTags')
  },
}
</script>
