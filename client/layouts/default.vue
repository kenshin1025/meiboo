<template>
  <v-app dark>
    <v-app-bar fixed app>
      <v-toolbar-title
        v-text="title"
        class="font-weight-black text-h3 red--text text--darken-1"
      />
      <v-spacer />
      <v-btn icon>
        <v-icon>mdi-menu</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer absolute app>
      <a>https://github.com/kenshin1025/meiboo</a>
      <v-spacer />
      <span>this is {{ this.$route.params.token }}</span>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      title: 'Meiboo!',
    }
  },
}
</script>
