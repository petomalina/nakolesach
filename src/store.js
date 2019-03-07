import Vue from 'vue'
import Vuex from 'vuex'

import Vapi from "vuex-rest-api"

Vue.use(Vuex)

const facts = new Vapi({
  baseURL: "https://api.chucknorris.io",
    state: {
      fact: "",
    }
  })
  // Step 3
  .get({
    action: "getFact",
    property: "fact",
    path: "/jokes/random"
  })
  .get({
    action: "updateFact",
    property: "fact",
    path: "/jokes/random"
  })
  // Step 4
  .getStore()

  export default new Vuex.Store(facts)
