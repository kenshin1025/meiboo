<template>
  <v-col cols="12" sm="8" md="6">
    <v-card
      class="pa-auto"
      width="400"
      height="140"
      outlined
      @click.stop="dialog = true"
    >
      <v-card-text>
        <h1>メンバーを追加</h1>
      </v-card-text>
    </v-card>
    <v-dialog v-model="dialog" width="500">
      <v-card>
        <v-card-title class="headline grey lighten-2 mb-4">
          メンバー追加
        </v-card-title>

        <v-card-text>
          <v-form v-model="valid">
            <v-file-input
              solo
              accept="image/png, image/jpeg, image/bmp"
              prepend-icon="mdi-camera"
              label="アイコン"
            ></v-file-input>
            <v-text-field
              v-model="name"
              :rules="nameRules"
              outlined
              label="名前"
              required
            ></v-text-field>
            <v-textarea
              v-model="comment"
              outlined
              label="コメント"
              required
            ></v-textarea>
            <v-combobox
              v-model="tags"
              :items="$store.state.meibo.tags"
              item-text="name"
              return-object
              multiple
              small-chips
              label="タグ"
            ></v-combobox>
          </v-form>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" :disabled="!valid" @click="submit">
            メンバーを追加する
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      valid: false,
      name: '',
      comment: '',
      nameRules: [(v) => !!v || '名前を入力してください'],
      tags: [],
    }
  },
  methods: {
    submit() {
      this.dialog = false
      this.$store.dispatch('meibo/addMember', {
        images: '',
        name: this.name,
        comment: this.comment,
        tags: this.tags,
      })
      this.name = ''
      this.comment = ''
      this.tags.splice(0, this.tags.length)
    },
  },
}
</script>
