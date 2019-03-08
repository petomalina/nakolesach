<template>
  <div class="rides">
    <md-card v-for="ride in rides" :key="ride.id">
      <md-card-header>
        <div class="md-title">From {{ride.from}} to {{ride.to}}</div>
      </md-card-header>

      <md-card-content>
        Time of ride: {{ride.boardingTime | date}}
      </md-card-content>

      <md-card-actions>
        <md-button>Edit</md-button>
        <md-button>Delete</md-button>
      </md-card-actions>
    </md-card>
  </div>
</template>

<script>

import { mapState, mapActions } from 'vuex'
var moment = require('moment');

export default {
    name: 'rides',
    created() {
        this.ListRides();
    },
    computed: mapState({
        rides: state => state.rides,
        pending: state => state.pending,
        error: state => state.error
    }),
    data: () => ({
        menuVisible: false
    }),
    filters: {
        date: function(value) {
            return moment(value).format('MMMM Do YYYY, hh:mm');
        }
    },
    methods: {
        toggleMenu () {
            this.menuVisible = !this.menuVisible
        },
        reload () {
            this.updateFact();
        },
        ...mapActions([
            "ListRides",
        ])
    }
}
</script>

<style lang="scss" scoped>
  .md-card {
      margin-left: 0px;
      margin-right: 0px;
  }
</style>