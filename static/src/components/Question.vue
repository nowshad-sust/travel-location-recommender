<template>
	<!-- <transition name="slide-fade" mode="out-in"> -->
	
  <div v-if="questions == null" class="text-center">
    <h4> </h4>
    <img src="../../dist/loading.gif" alt="loading">
  </div>

  <div v-else id="question" :key="this.index">
	
		<h4 class="title">{{q.text}}</h4>

		<div class="row">
			<div v-for="(option,i) in q.options" 
					:key="i" 
					:class="colLength(q.options.length)">
				<label>
				<input type="radio" 
						:name="this.index"
						class="form-check-input radio"
						:value="option.tags"
						@change="selected">

				<img v-lazy="option.img" 
					alt="Loading..."
					class="loaded-image img-fluid rounded">
				</label>
				{{option.name}}
			</div>
		</div>
	</div>
	<!-- </transition> -->
</template>

<script>
  import vm from '../App.vue';
  import axios from 'axios';
	export default {
        name: 'question',
		data () {
			return {
				q: null,
        index: 0,
        choices: [],
				questions: null
			}
    },
    http: {
      emulateJSON: true,
      emulateHTTP: true
    },
		created(){
      axios.get("/json").then(response => {
        // success callback
        this.questions = response.data;
        this.q = this.questions[0];
      }, response => {
        // error callback
        console.log(response);
      });
		},
		methods: {
			colLength: function(length){
				return ("col-md-"+(12/length));
			},
			selected: function(event){
        if(this.index < (this.questions.length-1)){
          this.q = this.questions[++this.index];
          this.choices.push(event.target.value);
				}else{
          if(this.choices.length < this.questions.length){
            this.choices.push(event.target.value);
          }
          this.$emit('childStart', "Result");
          vm.choices = this.choices;
				}
				event.target.checked = false;
			}
		}
  	};
</script>

<style>
label > input{ /* HIDE RADIO */
  visibility: hidden; /* Makes input not-clickable */
  position: absolute; /* Remove input from document flow */
}
label > input + img{ /* IMAGE STYLES */
  cursor:pointer;
  border:2px solid transparent;
}
label > input:checked + img{ /* (RADIO CHECKED) IMAGE STYLES */
  /* border: 2px solid lightblue; */
  -webkit-transition: all 0.10s;
  transition: all 0.10s;
  filter: brightness(0.8) drop-shadow(8px 8px 10px black);
  -webkit-filter: brightness(0.8) drop-shadow(8px 8px 10px black);
}

/* Enter and leave animations can use different */
/* durations and timing functions.              */
/* .slide-fade-enter-active {
  transition: all .1s ease;
}
.slide-fade-leave-active {
  transition: all .3s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to {
  transform: translateX(10px);
  opacity: 0;
} */
</style>