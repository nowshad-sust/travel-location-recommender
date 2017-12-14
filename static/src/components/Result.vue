<template>
	<div id="result">

      <!-- <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLongTitle">Modal title</h5>
      </div> -->
      <div class="modal-body">

        <h4>You should go to - </h4>
        <div v-if="response" class="text-center">
          <h3>{{response.name}}</h3>
          <carousel :autoplay="true" 
                    :autoplayTimeout="3000"
                    :navigationEnabled="true" 
                    :perPage="1" 
                    :loop="true">
            <slide v-for="(image,index) in response.Images" :key="index">
              <img :src="image" class="img-responsive slider-image">
              <h6 class="custom-caption">
                source: {{ image }}
              </h6>
            </slide>
          </carousel>

          <youtube :video-id="$youtube.getIdFromURL(response.Video)"
                  player-width="920"
                  player-height="520"
          ></youtube>
        </div>
            
        <div v-else class="text-center">
          <img src="../assets/Internet.gif" alt="loading" class="img-responsive">
          <h4>Taking you to the place you should be...</h4>               
        </div>
      </div>
      <div class="modal-footer">
        Give your feedback
      </div>

	</div>
</template>

<script>
  import vm from '../App.vue';
  import { Carousel, Slide } from 'vue-carousel';
  import axios from 'axios';

	export default {
     components: {
      Carousel,
      Slide
    },
		data () {
		return {
			response: null
			}
    },
    created(){
      let localChoices = [];
      
      vm.choices.forEach(function(choice) {
        let choiceArray = choice.split(',');
        choiceArray.forEach(function(tag){
          localChoices.push(tag);
        });
      }, this);
      
      var ref = this;

      axios.post('/v2/post', {data: localChoices}, {
          emulateJSON:true,
          headers: {'Access-Control-Allow-Origin': '*'}
        })
        .then(response => {
          console.log(response.data);
          // get body data
          ref.response = response.data;
      }, response => {
        // error callback
        console.log(response);
      });
       
      }
  	};
</script>

<style>

.VueCarousel {
  width: 98%;
}
.VueCarousel-slide {
  /* height: 80vh; */
  /* width: 80%; */
  position: relative;
  font-family: Arial;
  font-size: 24px;
  text-align: center;
  min-height: 100px;
}

.VueCarousel-slide img{
  width: 90%;
}

.label {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.custom-caption {
  font-size: 10px;
}
	
</style>