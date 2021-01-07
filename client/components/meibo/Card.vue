<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-list-item three-line>
      <v-list-item-avatar size="80" color="grey">
        <v-img :src="'https://joeschmoe.io/api/v1/' + user.img"></v-img>
      </v-list-item-avatar>
      <v-list-item-content>
        <div class="overline mb-4">OVERLINE</div>
        <v-list-item-title class="headline mb-1">
          {{ user.name }}
        </v-list-item-title>
        <v-list-item-subtitle>{{ user.comment }}</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
  </v-card>
</template>

<script>
export default {
  props: { user: Object },
}
</script>
