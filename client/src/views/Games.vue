<template>
  <div class="games">
    <h1>Jogos</h1>

    <h2>Os meus Jogos</h2>
    Ainda não tens jogos....
    
    <div>
      <div><img src="https://via.placeholder.com/150" /></div>
      <div>Maria Regina</div>
      <div>Ganhar ao Pai</div>
      <div>estado</div>
    </div>
    
    <button>Novo Jogo</button>

    <h2>Jogos Recentes</h2>
    <div>
      <div><img src="https://via.placeholder.com/150" /></div>
      <div>Maria Regina</div>
      <div>O meu jogo</div>
      <div>estado</div>
    </div>
    <div>
      <div><img src="https://via.placeholder.com/150" /></div>
      <div>Lobo Mau</div>
      <div>O meu jogo2</div>
      <div>estado</div>
    </div>
    <router-link to="/games">Ver todos...</router-link>
  </div>
</template>
<script>
import axios from '@/services/api'

export default {
  name: 'games',
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
