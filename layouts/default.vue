<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" :mini-variant="miniVariant" :clipped="clipped" fixed app>
      <v-list>
        <v-list-tile v-for="(item, i) in items" :key="i" :to="item.to" router exact>
          <v-list-tile-action>
            <font-awesome-icon :icon="['fas', 'user']" class="fa-2x"/>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title"/>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar :clipped-left="clipped" dark fixed app tabs color="indigo" dense>
      <v-toolbar-side-icon></v-toolbar-side-icon>
      <v-btn fab small color="pink" bottom left absolute @click="dialog = !dialog" style="left: 280px">
        <v-icon>add</v-icon>
      </v-btn>
      <v-toolbar-title>NKN WEB</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-tabs align-with-title color="transparent" right>
        <v-tab href="#one">Item One</v-tab>
        <v-tab href="#two">Item Two</v-tab>
        <v-tab href="#three">Item Three</v-tab>
        <v-tabs-slider color="pink"></v-tabs-slider>
      </v-tabs>

    </v-toolbar>
    <v-content>
      <v-container>
        <nuxt/>
      </v-container>
    </v-content>
    <Footer/>

  </v-app>
</template>

<script>
  import Footer from '~/components/Footer.vue'

  export default {
    components: {
      Footer
    },
    data() {
      return {
        clipped: true,
        drawer: true,
        fixed: false,
        items: [
          {
            icon: 'apps',
            title: 'Welcome',
            to: '/'
          },
          {
            icon: 'bubble_chart',
            title: 'Inspire',
            to: '/inspire'
          }
        ],
        miniVariant: false,
        right: true,
        rightDrawer: false,
        title: 'Vuetify.js'


      }
    }
  }
</script>
