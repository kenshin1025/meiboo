<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-list-item three-line>
      <v-list-item-avatar size="80" color="grey"> </v-list-item-avatar>
      <v-list-item-content>
        <div class="overline mb-4"></div>
        <v-list-item-title class="headline mb-1">
          <v-text-field placeholder="名前" solo @change="add()"></v-text-field>
        </v-list-item-title>
        <v-list-item-subtitle>
          <v-textarea
            placeholder="一言コメント"
            no-resize
            rows="2"
            solo
          ></v-textarea>
        </v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
  </v-card>
</template>

<script>
export default {
  data() {
    return {}
  },
  methods: {
    add() {
      console.log('ok')
    },
  },
}
</script>
