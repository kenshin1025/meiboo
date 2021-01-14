<template>
  <v-container>
    <v-row>
      <v-chip-group active-class="primary--text" column>
        <v-chip v-for="tag in $store.getters['meibo/filterTags']" :key="tag.id">
          {{ tag.name }}
        </v-chip>
      </v-chip-group>
    </v-row>
    <v-row>
      <v-checkbox
        v-for="tag in $store.state.meibo.tags"
        :key="tag.id"
        v-model="tags"
        :label="tag.name"
        :value="tag"
        class="mr-6"
        hide-details
        @change="setFilterTag"
      ></v-checkbox>
    </v-row>
    <v-row>
      <v-spacer></v-spacer>
      <v-btn text color="primary" @click.stop="dialog = true">
        +タグの追加
      </v-btn>
      <v-dialog v-model="dialog" width="500">
        <v-card>
          <v-card-title class="headline grey lighten-2 mb-4">
            タグ追加
          </v-card-title>

          <v-card-text>
            <v-text-field
              v-model="tagName"
              label="タグの名前"
              outlined
            ></v-text-field>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" :disabled="!tagName" @click="addTag">
              タグを追加する
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      word: '',
      tags: [],
      dialog: false,
      tagName: '',
    }
  },
  methods: {
    setSearchWord() {
      this.$store.commit('meibo/setSearchWord', this.word)
    },
    setFilterTag() {
      this.$store.commit('meibo/setFilterTag', this.tags)
    },
    addTag() {
      this.dialog = false
      this.$store.dispatch('meibo/addTag', {
        tagName: this.tagName,
      })
      this.tagName = ''
    },
  },
}
</script>
