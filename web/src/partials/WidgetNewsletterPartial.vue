<template>
	<div class="rounded-lg border border-slate-200 dark:border-slate-800 dark:bg-linear-to-t dark:from-slate-800 dark:to-slate-800/30 p-5">
		<div class="text-center mb-1">
			<div class="inline-flex -space-x-3 -ml-0.5">
				<img class="rounded-full border-2 border-white dark:border-slate-800 box-content" src="../images/avatar-01.jpg" width="24" height="24" alt="Avatar 01" />
				<img class="rounded-full border-2 border-white dark:border-slate-800 box-content" src="../images/avatar-02.jpg" width="24" height="24" alt="Avatar 02" />
				<img class="rounded-full border-2 border-white dark:border-slate-800 box-content" src="../images/avatar-03.jpg" width="24" height="24" alt="Avatar 03" />
				<img class="rounded-full border-2 border-white dark:border-slate-800 box-content" src="../images/avatar-04.jpg" width="24" height="24" alt="Avatar 04" />
				<img class="rounded-full border-2 border-white dark:border-slate-800 box-content" src="../images/avatar-05.jpg" width="24" height="24" alt="Avatar 05" />
			</div>
		</div>
		<div class="text-center mb-8">
			<div class="font-aspekta font-[650] mb-1">Never miss an update!</div>
			<p class="text-sm text-slate-500 dark:text-slate-400">Subscribe and join 100K+ developers.</p>
		</div>
		<form>
			<div class="mb-2">
				<label class="sr-only" for="newsletter">Your email…</label>
				<input id="newsletter" type="email" class="form-input py-1 w-full" placeholder="Your email…" />
			</div>
			<button class="btn-sm w-full text-slate-100 bg-sky-500 hover:bg-sky-600" type="submit">Subscribe</button>
		</form>
	</div>
</template>

<script>
export default {
	name: 'WidgetNewsletterPartial',
};
</script>
