<template>
  <v-app>
    <v-navigation-drawer class="border-right card-border" v-model="drawer" :mini-variant="miniVariant" :clipped="clipped" fixed floating app>

      <v-toolbar flat class="transparent">
        <v-list>
          <v-list-tile>
            <v-list-tile-title class="title">
              My wallet
            </v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-toolbar>
      <div class="divider"></div>
      <v-container grid-list-lg>
        <v-layout row wrap>
          <v-flex xs12>
            <Wallet></Wallet>
          </v-flex>
          <v-flex xs12>
            <Wallet2></Wallet2>
          </v-flex>
          <v-flex xs12>
            <AddWallet></AddWallet>
          </v-flex>

        </v-layout>
      </v-container>

    </v-navigation-drawer>
    <v-toolbar :clipped-left="clipped" dark fixed app tabs color="indigo" dense>
      <v-toolbar-side-icon @click="drawer = !drawer">
        <v-icon class="fas fa-wallet"></v-icon>
      </v-toolbar-side-icon>
      <div class="divider-vertical mx-3"></div>
      <v-toolbar-title>
        <v-avatar tile size="30">
          <img src="/img/logo.png" alt="avatar">
        </v-avatar>
        NKN WEB
      </v-toolbar-title>
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
      <Footer/>
    </v-content>


  </v-app>
</template>

<script>
  import Footer from '~/components/Footer.vue'
  import Wallet from '~/components/wallet/Wallet.vue'
  import Wallet2 from '~/components/wallet/Wallet2.vue'
  import AddWallet from '~/components/wallet/AddWallet.vue'
  export default {
    components: {
      Footer,
      Wallet,
      Wallet2,
      AddWallet
    },
    data() {
      return {
        clipped: false,
        drawer: true,
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
