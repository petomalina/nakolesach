import Vue from 'vue'
import Vuex from 'vuex'
import firebase from 'firebase'
import { firebaseAction, firebaseMutations } from 'vuexfire'

Vue.use(Vuex)
firebase.initializeApp({
  apiKey: "AIzaSyAW5mWlWwBfZOGAFz_D7rWrDLKxgjFxONY",
  authDomain: "nakolesach-sk.firebaseapp.com",
  databaseURL: "https://nakolesach-sk.firebaseio.com",
  projectId: "nakolesach-sk",
  storageBucket: "",
  messagingSenderId: "968808681042"
})

const db = firebase.firestore()
db.settings({ timestampsInSnapshots: true })

export default new Vuex.Store({
  strict: true,
  state: {
    rides: [], // Will be bound as an array
  },
  mutations: {
    ...firebaseMutations,
  },
  actions: {
    add: (action, ride) => {
      db.collection('rides').add(ride);
      console.log(action);
      console.log(ride);
    },
    bindRef: firebaseAction(({ bindFirebaseRef }, { name, ref }) => {
      bindFirebaseRef(name, ref)
    }),
    init: firebaseAction(({ bindFirebaseRef }) => {
      bindFirebaseRef('rides', db.collection('rides'))
    }),
  }
})
