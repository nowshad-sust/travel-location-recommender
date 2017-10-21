<template>
  <div id="app">
    <top-panel></top-panel>
    <hr>
    <div class="container-fluid">
      <keep-alive>
        <transition appear name="component-fade" mode="out-in">
          <component v-bind:is="currentView" @childStart="onClickChild">
            <!-- component changes when vm.currentView changes! -->
          </component>
        </transition>
      </keep-alive>
    </div>
  </div>
</template>

<script>
import TopPanel from './components/TopPanel.vue';
import Welcome from './components/Welcome.vue';
import MainContent from './components/MainContent.vue';
import Question from './components/Question.vue';
import Result from './components/Result.vue';

export default {
  name: 'app',
  data () {
    return {
      title: 'Travel The World',
      currentView: 'Welcome'
    }
  },
  components: {
    TopPanel,
    Welcome,
    MainContent,
    Question,
    Result
  },
  methods: {
    onClickChild (value) {
      this.currentView = value;
    }
  }
}
</script>

<style>
.thumbnail {
    position: relative;
}
.caption {
    color: white;
    position: absolute;
    top: 45%;
    left: 0;
    width: 100%;
}
.banner {
  width: 100%;
}

.component-fade-enter-active, .component-fade-leave-active {
  transition: all .3s ease;
}
.component-fade-enter, .component-fade-leave-to
/* .component-fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

</style>

