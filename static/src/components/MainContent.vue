<template>
	<div id="maincontent">
		<question @childStart="onClickChild"></question>
	</div>
</template>

<script>
	import Question from './Question.vue';
	export default {
        name: 'main-content',
		data () {
			return {
				
			}
		},
		components: {
			Question
		},
		methods: {
			onClickChild (value) {
				this.$emit('childStart', value);
			}
		}
  	};
</script>

<style>
	
</style>