<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-list-item three-line>
      <v-list-item-avatar size="80" color="grey">
        <v-img :src="'https://joeschmoe.io/api/v1/' + member.image"></v-img>
      </v-list-item-avatar>
      <v-list-item-content
        ><v-chip-group show-arrows center-active
          ><v-chip v-for="i in 5" :key="i">tag</v-chip></v-chip-group
        >
        <v-list-item-title
          v-if="!name_edit"
          class="headline mb-1"
          @dblclick="toggleNameEdit"
        >
          {{ member.name }}
        </v-list-item-title>
        <v-list-item-title v-else class="headline mb-1">
          <v-text-field
            :value="member.name"
            :placeholder="member.name"
            solo
            @change="changeName($event)"
          ></v-text-field>
        </v-list-item-title>
        <v-list-item-subtitle>{{ member.comment }}</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
  </v-card>
</template>

<script>
export default {
  props: { member: Object },
  data() {
    return {
      name_edit: false,
      name: '',
      comment_edit: false,
    }
  },
  methods: {
    toggleNameEdit() {
      this.name_edit = !this.name_edit
    },
    changeName(event) {
      this.$store.dispatch('meibo/editMemberName', {
        token: this.member.token,
        name: event,
      })
      this.toggleNameEdit()
    },
  },
}
</script>
