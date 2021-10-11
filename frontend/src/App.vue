<template>
    <Header/>
    <router-view/>
    <Footer/>
</template>

<script>
import Header from "@/components/Header";
import Footer from "@/components/Footer";
export default {
  name: "App",
  components: {
    Header,
    Footer
  }
}
</script>

<style lang="scss">
  #app {
    background: #3B3A3C;
    padding: 0;
    margin: 0;
  }
</style>
