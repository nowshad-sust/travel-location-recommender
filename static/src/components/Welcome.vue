<template>
	<div id="welcome">
		<div class="contaiter text-center">
			<button class="btn btn-primary" @click="onClickButton">
				Let's Go Somewhere!
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