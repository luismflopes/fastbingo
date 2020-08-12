<template>
  <div class="play">
    <h1>This is the Play page</h1>

    <input v-model="name" placeholder="Name"/>

    <button @click="clickMe">aaaa</button>
  </div>
</template>
<script>
import axios from '@/services/api'

export default {
  name: 'play',
  components: {
  },
  data () {
    return {
      name: ''
    }
  },
  mounted () {
    // this.loadCmsPage(this.$route.params.cmsSlug)
  },
  watch: {
    // '$route.params.cmsSlug': function (slug) {
    //   this.loadCmsPage(slug)
    // }
  },
  methods: {
    clickMe (e) {
      console.log(e, this.name);

      axios.post('/play', {name: this.name}).then(response => {
        console.log(response.data)
        this.cmsPage = response.data;
      }).catch(err => {
        console.log(err);
      });
    }
    // loadCmsPage (slug) {
    //   axios.get('/cms-pages/' + slug).then(response => {
    //     console.log(response.data)
    //     this.cmsPage = response.data
    //   }).catch(err => {
    //     console.log(err)
    //   })
    // }
  }
};
</script>
