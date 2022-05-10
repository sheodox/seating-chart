/// <reference types="vite/client" />
/// <reference types="svelte" />
//import 'vite/modulepreload-polyfill';
import { styles } from 'sheodox-ui';
//import './scss/style.scss';
//import '../../node_modules/fontawesome-free/css/all.min.css';

import './stores/routing';
import App from './App.svelte';
import '../node_modules/@fortawesome/fontawesome-free/css/all.css';

// import Trancemaker from './components/Trancemaker';
// new Trancemaker();

new App({
	target: document.getElementById('app-root'),
});
