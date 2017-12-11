<template>
	<div id="landing">
        <div class="thumbnail text-center">
        <img class="img-fluid banner" src="/dist/inner.png">
        <div class="caption">
                <h2>
                <a class="unstyled" href="/">
                    {{ title }}
                </a>
                </h2>
            </div>
        </div>
	</div>
</template>

<script>
	export default {
        data () {
            return {
                title: 'Travel The World'
            }
        },
  	};
</script>

<style>
	.caption a {
        color: black;
        text-decoration: none;
    }
    .caption a:hover, .caption a:focus {
        text-decoration: blue;
    }
</style>