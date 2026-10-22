/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Composables
import 'unfonts.css';
import App from './App.vue';
import {
    createApp
} from 'vue';
// Plugins
import {
    registerPlugins
} from '@/plugins';

const app = createApp( App );

registerPlugins( app );

app.mount( '#app' );
