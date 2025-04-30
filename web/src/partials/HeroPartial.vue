<template>
	<section>
		<div class="max-w-[700px]">
			<div class="pt-8 pb-10">
				<AvatarPartial width="w-14" />
				<h1 class="h1 blog-h1">
					I write about coding, engineering, and
					<span class="blog-fun-title-word-highlight"
						>leadership</span
					>
					as a service.
				</h1>
				<p class="text-lg text-slate-400 dark:text-slate-300 mb-5">Writer, Speaker, Developer, Founder, and Leadership.</p>
				<p class="mb-5 font-aspekta text-slate-500 dark:text-slate-500">
					I'm a full-stack Software Engineer leader with over two decades of building complex web systems and
                    products, specialising in areas like e-commerce, banking, cross-payment solutions, cyber
					security, and customer success.
				</p>
			</div>
		</div>
	</section>
</template>

<script>
import AvatarPartial from '@partials/AvatarPartial.vue';

export default {
	name: 'HeroPartial',
	components: {
		AvatarPartial,
	},
};
</script>
