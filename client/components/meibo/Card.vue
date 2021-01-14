<template>
  <div>
    <v-card
      class="mx-auto"
      width="400"
      height="140"
      outlined
      @click.stop="dialog = true"
    >
      <v-list-item three-line>
        <v-list-item-avatar size="80" color="grey">
          <v-img
            :src="'https://picsum.photos/seed/' + member.token + '/100'"
          ></v-img>
        </v-list-item-avatar>
        <v-list-item-content
          ><v-chip-group center-active column
            ><v-chip v-for="(tag, i) in member.tags" :key="i" x-small>
              {{ tag.name }}
            </v-chip></v-chip-group
          >
          <v-list-item-title class="headline mb-1" @dblclick="toggleNameEdit">
            {{ member.name }}
          </v-list-item-title>
          <v-list-item-subtitle>{{ member.comment }}</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </v-card>
    <v-dialog v-model="dialog" width="500">
      <v-card>
        <v-list-item three-line>
          <v-list-item-avatar size="160" color="grey">
            <v-img
              :src="'https://picsum.photos/seed/' + member.token + '/200'"
            ></v-img>
          </v-list-item-avatar>
          <v-list-item-content
            ><v-chip-group center-active column
              ><v-chip v-for="(tag, i) in member.tags" :key="i" x-small>
                {{ tag.name }}
              </v-chip></v-chip-group
            >
            <v-list-item-title
              v-if="!name_edit"
              class="text-h4 mb-1"
              @dblclick="toggleNameEdit"
            >
              {{ member.name }}
            </v-list-item-title>
            <v-list-item-title v-else class="mb-1">
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
    </v-dialog>
  </div>
</template>

<script>
export default {
  props: { member: Object },
  data() {
    return {
      name_edit: false,
      name: '',
      comment_edit: false,
      dialog: false,
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
