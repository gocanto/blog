import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '@pages/HomePage.vue';
import PostPage from '@pages/PostPage.vue';
import AboutPage from '@pages/AboutPage.vue';
import ProjectsPage from '@pages/ProjectsPage.vue';
import ResumePage from '@pages/ResumePage.vue';
import SubscribePage from '@pages/SubscribePage.vue';

const routerHistory = createWebHistory();

const router = createRouter({
	scrollBehavior(to) {
		if (to.hash) {
			window.scroll({ top: 0 });
		} else {
			document.querySelector('html').style.scrollBehavior = 'auto';
			window.scroll({ top: 0 });
			document.querySelector('html').style.scrollBehavior = '';
		}
	},
	history: routerHistory,
	routes: [
		{
			path: '/',
			component: HomePage,
		},
		{
			path: '/post',
			component: PostPage,
		},
		{
			path: '/about',
			component: AboutPage,
		},
		{
			path: '/projects',
			component: ProjectsPage,
		},
		{
			path: '/resume',
			component: ResumePage,
		},
		{
			path: '/subscribe',
			component: SubscribePage,
		},
	],
});

export default router;
