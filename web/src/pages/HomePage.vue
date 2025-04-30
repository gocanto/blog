<template>
	<div class="max-w-7xl mx-auto">
		<div class="min-h-screen flex">
			<SideNavPartial />

			<!-- Main content -->
			<main class="grow overflow-hidden px-6">
				<div class="w-full h-full max-w-[1072px] mx-auto flex flex-col">
					<HeaderPartial />
					<HeroPartial />

					<!-- Content -->
					<div class="grow md:flex space-y-8 md:space-y-0 md:space-x-8 pb-16 md:pb-20">
						<!-- Middle area -->
						<div class="grow">
							<div class="max-w-[700px]">
								<div class="space-y-10">
									<ArticlesListPartial />
									<TalksPartial />
									<FeaturedProjectsPartial />
								</div>
							</div>
						</div>

						<!-- Right sidebar -->
						<aside class="md:w-[240px] lg:w-[300px] shrink-0">
							<div class="space-y-6">
								<WidgetNewsletterPartial />
								<WidgetSponsorPartial />
								<WidgetBookPartial />
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
import HeroPartial from '@partials/HeroPartial.vue';
import ArticlesListPartial from '@partials/ArticlesListPartial.vue';
import TalksPartial from '@partials/TalksPartial.vue';
import FeaturedProjectsPartial from '@partials/FeaturedProjectsPartial.vue';
import WidgetNewsletterPartial from '@partials/WidgetNewsletterPartial.vue';
import WidgetSponsorPartial from '@partials/WidgetSponsorPartial.vue';
import WidgetBookPartial from '@partials/WidgetBookPartial.vue';
import FooterPartial from '@partials/FooterPartial.vue';

export default {
	name: 'HomePage',
	components: {
		SideNavPartial,
		HeaderPartial,
		HeroPartial,
		ArticlesListPartial,
		TalksPartial,
		FeaturedProjectsPartial,
		WidgetNewsletterPartial,
		WidgetSponsorPartial,
		WidgetBookPartial,
		FooterPartial,
	},
};
</script>
