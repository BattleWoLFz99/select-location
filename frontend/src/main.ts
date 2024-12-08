// import './assets/main.css'
import App from './App.vue'
import { createApp, provide, h } from 'vue'
import { ApolloClients,DefaultApolloClient } from '@vue/apollo-composable'
import apolloClient from './tools/apollo'
import { provideApolloClient } from "@vue/apollo-composable";

const app = createApp({
    setup () {
        provide(ApolloClients, {
            default: apolloClient,
        })
    },
    render: () => h(App),
});
provideApolloClient(apolloClient);
app.mount('#app');
