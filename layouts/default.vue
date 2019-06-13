<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" :mini-variant="miniVariant" :clipped="clipped" fixed app floating="false" class="border-right card-border">
      <v-toolbar flat class="transparent">
        <v-list >
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <img src="https://demos.creative-tim.com/vuetify-material-dashboard/img/vuetifylogo.png">
            </v-list-tile-avatar>

            <v-list-tile-content>
              <v-list-tile-title>NKN WEB</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-toolbar>
        <v-container
          fluid
          grid-list-lg
        ><v-divider></v-divider>
          <v-layout row wrap>
            <v-flex xs12>
              <v-card color="blue-grey darken-2" class="white--text">
                <v-card-title primary-title>
                  <div>
                    <div class="headline">Unlimited music now</div>
                    <span>Listen to your favorite artists and albums whenever and wherever, online and offline.</span>
                  </div>
                </v-card-title>
                <v-card-actions>
                  <v-btn flat dark>Listen now</v-btn>
                </v-card-actions>
              </v-card>
            </v-flex>
            <v-flex xs12>
              <v-card color="cyan darken-2" class="white--text">
                <v-layout>
                  <v-flex xs5>
                    <v-img
                      src="https://cdn.vuetifyjs.com/images/cards/foster.jpg"
                      height="125px"
                      contain
                    ></v-img>
                  </v-flex>
                  <v-flex xs7>
                    <v-card-title primary-title>
                      <div>
                        <div class="headline">Supermodel</div>
                        <div>Foster the People</div>
                        <div>(2014)</div>
                      </div>
                    </v-card-title>
                  </v-flex>
                </v-layout>
                <v-divider light></v-divider>
                <v-card-actions class="pa-3">
                  Rate this album
                  <v-spacer></v-spacer>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                </v-card-actions>
              </v-card>
            </v-flex>
            <v-flex xs12>
              <v-card color="purple" class="white--text">
                <v-layout row>
                  <v-flex xs7>
                    <v-card-title primary-title>
                      <div>
                        <div class="headline">Halycon Days</div>
                        <div>Ellie Goulding</div>
                        <div>(2013)</div>
                      </div>
                    </v-card-title>
                  </v-flex>
                  <v-flex xs5>
                    <v-img
                      src="https://cdn.vuetifyjs.com/images/cards/halcyon.png"
                      height="125px"
                      contain
                    ></v-img>
                  </v-flex>
                </v-layout>
                <v-divider light></v-divider>
                <v-card-actions class="pa-3">
                  Rate this album
                  <v-spacer></v-spacer>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                  <v-icon>star_border</v-icon>
                </v-card-actions>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>

    </v-navigation-drawer>
    <v-toolbar :clipped-left="clipped" dark fixed app tabs color="indigo" dense>
      <v-toolbar-side-icon @click="drawer = !drawer"></v-toolbar-side-icon>
      <v-btn fab small color="pink" bottom left absolute @click="drawer = !drawer" style="left: 280px" app>
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
      <Footer/>
    </v-content>


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
