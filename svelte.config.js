import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess({
		postcss: true
	}),
	vite: {
		css: {
		  preprocessorOptions: {
			scss: {
			  additionalData: '@use "src/variables.scss" as *;',
			},
		  },
		},
		ssr: {
		  noExternal: ['@fortawesome/*'],
		}
	  },
	kit: {
		adapter: adapter({
			pages: 'pb_public',
			assets: 'pb_public',
			fallback: "index.html",
			precompress: false
		})
	}
};

export default config;
