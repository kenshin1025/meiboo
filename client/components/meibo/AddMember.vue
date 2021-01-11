<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-card-title>メンバーを追加する</v-card-title>
    <v-card-text>
      <v-text-field
        hint="追加したいメンバーの名前を入力してください"
        placeholder="名前"
        solo
        @change="add()"
      ></v-text-field>
    </v-card-text>
    <v-card-actions><v-spacer></v-spacer><v-btn>追加</v-btn></v-card-actions>
    <!-- <v-list-item three-line>
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
    </v-list-item> -->
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
