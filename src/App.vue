<template>
  <div id="app" class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary" md-elevation="0">
        <md-button class="md-icon-button" @click="toggleMenu" v-if="!menuVisible">
          <md-icon>menu</md-icon>
        </md-button>
        <span class="md-title">My Title</span>
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="menuVisible" md-persistent="mini">
        <md-toolbar class="md-transparent" md-elevation="0">
          <span>Navigation</span>

          <div class="md-toolbar-section-end">
            <md-button class="md-icon-button md-dense" @click="toggleMenu">
              <md-icon>keyboard_arrow_left</md-icon>
            </md-button>
          </div>
        </md-toolbar>

        <md-list>
          <router-link to="/rides">
            <md-list-item>
              <md-icon>directions_car</md-icon>
              <span class="md-list-item-text">Rides</span>
            </md-list-item>
          </router-link>

          <router-link to="/requests">
            <md-list-item>
              <md-icon>transfer_within_a_station</md-icon>
              <span class="md-list-item-text">Requests</span>
            </md-list-item>
          </router-link>
  
          <router-link to="/settings">
            <md-list-item>
              <md-icon>settings</md-icon>
              <span class="md-list-item-text">Settings</span>
            </md-list-item>
          </router-link>

        </md-list>
      </md-app-drawer>

      <md-app-content>
        <router-view/>
        <!--<md-button class="md-primary" @click="reload">New fact</md-button>
        Fact: {{ fact.value }}
        <p v-if="pending.fact">loading posts...</p>
        <p v-if="error.fact">loading failed</p>-->
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
  import { mapState, mapActions } from 'vuex'
  export default {
    name: 'PersistentMini',
    created() {
      this.getFact();
    },
    computed: mapState({
      fact: state => state.fact,
      pending: state => state.pending,
      error: state => state.error
    }),
    data: () => ({
      menuVisible: false
    }),
    methods: {
      toggleMenu () {
        this.menuVisible = !this.menuVisible
      },
      reload () {
        this.updateFact();
      },
      ...mapActions([
        "getFact",
        "updateFact",
      ])
    }
  }
</script>

<style lang="scss" scoped>
  @import "~vue-material/dist/theme/engine"; // Import the theme engine

  @include md-register-theme("default", (
    primary: md-get-palette-color(blue, A200), // The primary color of your application
    accent: md-get-palette-color(red, A200) // The accent or secondary color
  ));

  @import "~vue-material/dist/theme/all"; // Apply the theme

  .md-app {
    min-height: 350px;
    border: 1px solid rgba(#000, .12);
  }

   // Demo purposes only
  .md-drawer {
    width: 230px;
    max-width: calc(100vw - 125px);
  }
</style>
