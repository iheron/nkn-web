<template>
  <v-footer height="auto">
    <v-card flat tile class="indigo lighten-1 white--text text-xs-center">
      <v-card-text>
        <v-tooltip top v-for="item in icons" :key="item.icon">
          <template v-slot:activator="{ on }">
            <v-btn href="https://github.com/nknorg" target="_blank" class="mx-3 white--text" icon v-on="on">
              <v-icon size="24px">{{item.icon}}</v-icon>
            </v-btn>
          </template>
          <span>{{item.tooltip}}</span>
        </v-tooltip>


      </v-card-text>

      <v-card-text class="white--text pt-0">
        <h4>NKN:
          Network Infra for Decentralized Internet</h4>
        NKN is the new kind of P2P network connectivity protocol & ecosystem powered by a novel public blockchain. We use economic incentives to motivate Internet users to share network connection and utilize unused bandwidth. NKN's open, efficient, and robust networking infrastructure enables application developers to build the decentralized Internet so everyone can enjoy secure, low cost, and universally accessible connectivity.
      </v-card-text>

      <v-divider></v-divider>

      <v-card-text class="white--text">
        Copyright &copy; 2019 — <strong>nkn-web</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>
<script>
  export default {
    data:() =>({
      icons: [
        {icon: 'fab fa-github', tooltip: 'GitHub'},
        {icon: 'fab fa-discord', tooltip: 'Discord'},
        {icon: 'fab fa-twitter', tooltip: 'Twitter'},
        {icon: 'fab fa-medium', tooltip: 'Medium'},
        {icon: 'fab fa-youtube', tooltip: 'Youtube'},
        {icon: 'fab fa-reddit', tooltip: 'Reddit'},
        {icon: 'fab fa-telegram', tooltip: 'Telegram'},
        {icon: 'fas fa-comments', tooltip: 'Forum'},
      ]
    })
  }
</script>
