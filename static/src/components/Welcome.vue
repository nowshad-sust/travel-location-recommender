<template>
	<div id="welcome">
		<div class="contaiter text-center">
			Hello World
			<button class="btn btn-primary" @click="onClickButton">
				Let's Go!
			</button>
		</div>
	</div>
</template>

<script>
	export default {
		methods: {
			onClickButton (event) {
				this.$emit('childStart', 'MainContent')
			}
		}
  	};
</script>

<style>
	
</style>