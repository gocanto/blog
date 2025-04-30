<template>
	<div class="max-w-7xl mx-auto">
		<div class="min-h-screen flex">
			<SideNavPartial />

			<!-- Main content -->
			<main class="grow overflow-hidden px-6">
				<div class="rounded-lg w-full h-full max-w-[1072px] mx-auto flex flex-col">
					<HeaderPartial />

					<!-- Content -->
					<div class="grow md:flex space-y-8 md:space-y-0 md:space-x-8 pt-12 md:pt-16 pb-16 md:pb-20">

                        <!-- Middle area -->
						<div class="grow">
							<div class="max-w-[700px]">
								<section>
									<!-- Page title -->
									<h1 class="h1 font-aspekta mb-5 mt-5 text-slate-700 dark:text-slate-300">
										Hi. I'm {{ user.profile.nickname }}
										<span class="blog-fun-title-word-highlight">
											<a :href="user.social.x.url" :title="user.social.x.title" target="_blank">{{ user.social.x.handle }}</a>
										</span>
									</h1>

									<img class="rounded-lg w-full mb-5" :src="aboutPicture" alt="About" />

									<!-- Page content -->
									<div class="space-y-8">
										<div class="space-y-4">
											<h2 class="h2 font-aspekta text-slate-700 dark:text-slate-300">About</h2>
											<p class="block mb-5 text-slate-500">
												<span class="block">
													I am a dedicated engineering leader passionate about building seamless, high-quality experiences for organizations and
													<a class="blog-link" target="_blank" :href="user.social.github.url">open source</a>. With over twenty years of&nbsp;
													<router-link v-slot="{ href, navigate }" to="/resume">
														<a class="blog-link" :href="href" @click="navigate">experience</a>
													</router-link>
													in software development, workplace technology, and infrastructure management.
												</span>
												<span class="block mt-5 text-slate-500">
													I specialize in Golang, Node.js, TypeScript, and PHP. I am also proficient in Laravel, Symfony, and modern web frameworks like Next.Js. Furthermore, I have led teams designing and
													implementing scalable, high-performance systems, ensuring reliability and efficiency across complex technical environments.
												</span>
											</p>
											<p class="block mb-3 text-slate-500">
												Beyond technical expertise, I have a strong <a class="blog-link" :href="user.social.linkedin.url" target="_blank">leadership background</a> in managing cross-functional teams, optimizing
												workflows, and implementing best practices that drive productivity and innovation. I thrive in fast-paced environments that demand strategic thinking, problem-solving, and a commitment to
												delivering high-quality results.
											</p>
										</div>

										<ExperiencePartial :experience="user.experience" />

										<div class="mt-5 space-y-5">
											<h2 class="h2 font-aspekta text-slate-700 dark:text-slate-300">Let's Connect</h2>
											<p>
												I’m excited to connect by <a class="blog-link" title="follow me on x" :href="`mailto:${user.profile.email}`">email</a> or
												<a class="blog-link" target="_blank" :href="user.social.x">X</a> to chat about projects and ideas. I’m always open to freelance or long-term projects, so please feel free to reach out.
											</p>
											<p>Tell me about your vision, and if it seems like a good fit, we can explore collaborating down the road.</p>
										</div>
									</div>
								</section>
								<!-- content --->
							</div>
						</div>

						<!-- Right sidebar -->
						<aside class="md:w-[240px] lg:w-[300px] shrink-0">
							<div class="space-y-6">
								<WidgetNewsletterPartial />

								<WidgetSponsorPartial />
							</div>
						</aside>
					</div>

					<FooterPartial />
				</div>
			</main>
		</div>
	</div>
</template>

<script>
import SideNavPartial from '@partials/SideNavPartial.vue';
import HeaderPartial from '@partials/HeaderPartial.vue';
import ExperiencePartial from '@partials/ExperiencePartial.vue';
import WidgetNewsletterPartial from '@partials/WidgetNewsletterPartial.vue';
import WidgetSponsorPartial from '@partials/WidgetSponsorPartial.vue';
import FooterPartial from '@partials/FooterPartial.vue';

import AboutPicture from '@images/profile/about.png';
import UserResponse from '@response/user-response.json';

import { useRouter } from 'vue-router';

export default {
	name: 'AboutPage',
	components: {
		SideNavPartial,
		HeaderPartial,
		ExperiencePartial,
		WidgetNewsletterPartial,
		WidgetSponsorPartial,
		FooterPartial,
	},
	setup() {
		const currentRoute = useRouter().currentRoute.value;
		const user = UserResponse;

		return {
			currentRoute,
			user,
		};
	},

	computed: {
		aboutPicture() {
			return AboutPicture;
		},
	},
};
</script>
