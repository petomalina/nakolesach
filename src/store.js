import Vue from 'vue'
import Vuex from 'vuex'

import Vapi from "vuex-rest-api"

Vue.use(Vuex)

const facts = new Vapi({
  baseURL: "https://europe-west1-nakolesach-sk.cloudfunctions.net",
    state: {
      rides: [],
    }
  })
  // Step 3
  // .get({
  //   action: "getFact",
  //   property: "fact",
  //   path: "/jokes/random"
  // })
  // .get({
  //   action: "updateFact",
  //   property: "fact",
  //   path: "/jokes/random"
  // })
  .get({
    action: "ListRides",
    property: "rides",
    path: "/ListRides"
  })
  // Step 4
  .getStore()

  export default new Vuex.Store(facts)
