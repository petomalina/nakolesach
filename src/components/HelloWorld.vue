<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <ul>
        <li v-for="ride in rides" :key="ride.id">{{ ride.text }}</li>
    </ul>
    <form @submit.prevent="addTodo">
      <input v-model="newRide" /> <button>Add</button>
    </form>
  </div>
</template>

<script lang="ts">
import firebase from 'firebase';
import { Component, Prop, Vue } from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';
import { mapState } from 'vuex';

@Component
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;

  private newRide: string = '';
  @State rides: any;
  @Getter('rides') ridesGetter: any;
  @Action init: any;
  @Action add: any;

  created () {
    this.init();
  }
  mounted () {
    console.log(this.rides);
  }

  addTodo() {
    if (this.newRide) {
      this.add({
        text: this.newRide,
        created: firebase.firestore.FieldValue.serverTimestamp(),
      })
      this.newRide = ''
    }
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
