<template>
  <v-layout>
    <v-flex text-xs-center>
      <img
        src="/v.png"
        alt="Vuetify.js"
        class="mb-5"
      >
      <blockquote class="blockquote">
        {{$t('language')}} {{$t('welcome')}}
        &#8220;First, solve the problem. Then, write the code.&#8221;
        <footer>
          <small>
            <em>&mdash;John Johnson</em>
          </small>
        </footer>
      </blockquote>
    </v-flex>
  </v-layout>
</template>
<script>
  export default {
    nuxtI18n: {
      paths: {
        en: '/inspire', // -> accessible at /about-us (no prefix since it's the default locale)
        zh: '/index', // -> accessible at /fr/a-propos

      }
    }
  }
</script>
