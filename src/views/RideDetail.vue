<template>
  <div class="about">
    <h1 v-if="!$route.params.id">New ride</h1>
    <h1>{{ ride }}</h1>
  </div>
</template>

<script>

import { mapState, mapActions } from 'vuex'
var moment = require('moment');

export default {
    name: 'ride-detail',
    created() {
        if (this.$route.params.id) {
            this.GetRide({ params: { id: this.$route.params.id } });
        } else {
            // this.GetRidePreset();
        }
    },
    computed: mapState({
        ride: state => state.ride,
        pending: state => state.pending,
        error: state => state.error
    }),
    data: () => ({
        menuVisible: false,
        showDialog: false
    }),
    filters: {
        date: function(value) {
            return moment(value).format('MMMM Do YYYY, hh:mm');
        }
    },
    methods: {
        ...mapActions([
            "GetRide",
        ])
    }
}
</script>