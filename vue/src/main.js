import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import 'primeflex/primeflex.css' 
import Card from 'primevue/card';
import Message from 'primevue/message';
import Toolbar from 'primevue/toolbar';
import ScrollPanel from 'primevue/scrollpanel';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Listbox from 'primevue/listbox';
import Splitter from 'primevue/splitter';
import SplitterPanel from 'primevue/splitterpanel';
import ToastService from 'primevue/toastservice';
import Toast from 'primevue/toast';
import 'primeicons/primeicons.css'
import 'primevue/resources/themes/lara-dark-indigo/theme.css'
import 'primevue/resources/primevue.min.css'

/* eslint-disable */
const app = createApp(App)

app.use(PrimeVue)
app.use(ToastService)
app.component('Card', Card)
app.component("Message", Message)
app.component("Toolbar", Toolbar)
app.component('ListBox',Listbox)
app.component("InputText", InputText)
app.component("Button", Button)
app.component("Dialog", Dialog)
app.component("Toast", Toast)
app.component("Splitter", Splitter)
app.component("SplitterPanel ", SplitterPanel)
app.component("ScrollPanel  ", ScrollPanel)
app.mount('#app')