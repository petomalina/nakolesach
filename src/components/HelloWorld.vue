<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <ul>
        <li v-for="ride in rides" :key="ride.id">{{ ride.text }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { State, Action } from 'vuex-class';
import { mapState } from 'vuex';

@Component
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;
  @State rides: any;
  @Action init: any;

  created () {
    this.init();
  }
  mounted () {
    console.log(this.rides);
  }
  updated () { }
  destroyed () { }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
