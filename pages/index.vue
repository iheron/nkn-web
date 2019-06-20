<template>
  <v-layout flex row wrap>
    <v-flex>
      <v-checkbox v-model="selected" label="John" value="John"></v-checkbox>
      <v-checkbox v-model="selected" label="Jacob" value="Jacob"></v-checkbox>
      <LastMiningRewards></LastMiningRewards>
    </v-flex>
    <v-flex>
      <Neighbors></Neighbors>
    </v-flex>
  </v-layout>

</template>

<script>
  import LastMiningRewards from '~/components/widget/LastMiningRewards.vue'
  import Neighbors from '~/components/widget/Neighbors.vue'

  export default {
    components: {
      LastMiningRewards,
      Neighbors
    },

    async mounted() {

    }
  }
</script>
